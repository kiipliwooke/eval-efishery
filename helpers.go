package main

import (
	"io"
	"net/http"
	"os"
)

func SaveFile(r *http.Request) (bool, string) {
	//get file
	file, header, errFile := r.FormFile("file")
	if errFile != nil {
		return false, ""
	}
	//gatau ngapain
	defer file.Close()

	//making directory
	filename := header.Filename
	var f *os.File
	f, errFile = os.OpenFile("./static/sign/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if errFile != nil {
		return false, ""
	}
	defer f.Close()
	//copying file
	io.Copy(f, file)
	return true, filename
}

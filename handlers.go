package main

import (
	"encoding/json"
	ent "eval-efishery/entity"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if LoginUser.Username != "" {
		fmt.Println("masuk sini")
		json.NewEncoder(w).Encode("Already login")
	} else if r.Method == http.MethodPost {
		fmt.Println("masuk sono")
		//decode the request body
		var newUserLogin ent.Account
		json.NewDecoder(r.Body).Decode(&newUserLogin)

		//find the account
		result := db.Model(&ent.Account{}).Where("username = ? AND password = ?", newUserLogin.Username, newUserLogin.Password).Find(&newUserLogin)
		if result.Error != nil || result.RowsAffected <= 0 {
			json.NewEncoder(w).Encode("Username or password invalid")
		} else {
			LoginUser.Username = newUserLogin.Username
			LoginUser.Admin = newUserLogin.Admin
			json.NewEncoder(w).Encode(newUserLogin)
		}
	} else {
		json.NewEncoder(w).Encode("Nothing happened")
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	LoginUser = ent.UserLogin{
		Username: "",
		Admin:    false,
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("logout successful")
}

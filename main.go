package main

import (
	ent "eval-efishery/entity"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var LoginUser ent.UserLogin
var err error
var db *gorm.DB

func main() {
	//get env file
	err = godotenv.Load()
	DB_URI := os.Getenv("DB_URI")
	db, err = gorm.Open(postgres.Open(DB_URI), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//migrate all struct
	db.AutoMigrate(&ent.Account{}, &ent.Document{}, &ent.Loan{}, &ent.LoanApplication{}, &ent.User{})

	//creating dummy if haven't
	var first ent.Account
	result := db.First(&first)
	if result.RowsAffected <= 0 {
		DummyGenerator()
	}

	mux := http.NewServeMux()

	//list of all handlers
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/logout", Logout)
	mux.HandleFunc("/allLoan", GetAllLoans)
	mux.HandleFunc("/loan", GetLoan)
	mux.HandleFunc("/uploadDocument", UploadDocument)

	http.ListenAndServe(":3000", mux)
}

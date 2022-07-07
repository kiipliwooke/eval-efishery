package main

import (
	ent "eval-efishery/entity"
	"fmt"
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

	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/logout", Logout)

	http.ListenAndServe(":3000", mux)
	fmt.Println("Hello, World!")
}

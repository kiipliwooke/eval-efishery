package main

import (
	ent "eval-efishery/entity"
	"fmt"
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

	//creating dummy
	DummyGenerator()

	fmt.Println("Hello, World!")
}

package main

import (
	"encoding/json"
	ent "eval-efishery/entity"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if LoginUser.Username != "" {
		json.NewEncoder(w).Encode("Already login")
	} else if r.Method == http.MethodPost {
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

func GetAllLoans(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if LoginUser.Username == "" {
		json.NewEncoder(w).Encode("Please login first")
	} else if r.Method == http.MethodGet {
		//if admin
		var loans []ent.LoanApplication
		if LoginUser.Admin {
			db.Model(&ent.LoanApplication{}).Find(&loans)
		} else { //not admin
			//find user_id asumes there IS ONE with current username
			var user_id int
			db.Model(&ent.User{}).Select("user_id").Where("username = ?", LoginUser.Username).Find(&user_id)
			//find loan with current user_id
			db.Model(&ent.LoanApplication{}).Where("user_id = ?", user_id).Find(&loans)
		}
		json.NewEncoder(w).Encode(loans)
	} else {
		json.NewEncoder(w).Encode("Nothing happened")
	}
}

func GetLoan(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if LoginUser.Username == "" {
		json.NewEncoder(w).Encode("Please login first")
	} else if r.Method == http.MethodGet {
		//parsing query string
		loanID := r.URL.Query().Get("lan")

		var accessable = true
		//checking not admin
		if !LoginUser.Admin {
			//find userID
			var userID int
			//asume there is one userID
			db.Model(&ent.User{}).Select("user_id").Where("username = ?", LoginUser.Username).Find(&userID)

			//find combination userID and loanID
			var laTmp ent.LoanApplication
			result := db.Model(&ent.LoanApplication{}).Where("user_id = ? AND loan_id = ?", userID, loanID).Find(&laTmp)
			//handling
			if result.RowsAffected <= 0 {
				accessable = false
			}
		}
		if accessable {
			var detailedLoan ent.DetailedLoan
			//finding loan
			result := db.Model(&ent.Loan{}).Where("lan = ?", loanID).Find(&detailedLoan.Loan)
			if result.RowsAffected <= 0 {
				json.NewEncoder(w).Encode("cannot find loan id")
			} else {
				//get attachments
				db.Model(&ent.Document{}).Where("attachment_id = ?", loanID).Find(&detailedLoan.Documents)
				json.NewEncoder(w).Encode(detailedLoan)
			}

		} else {
			json.NewEncoder(w).Encode("Access disabled")
		}
	} else {
		json.NewEncoder(w).Encode("Nothing happened")
	}
}

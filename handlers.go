package main

import (
	"encoding/json"
	ent "eval-efishery/entity"
	"io"
	"net/http"
	"os"
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
		var detailedLoan ent.DetailedLoan
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
			detailedLoan.LenderSign = laTmp.LenderSign
			detailedLoan.LoanerSign = laTmp.LoanerSign
		}
		if accessable {
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

func UploadDocument(w http.ResponseWriter, r *http.Request) {
	if LoginUser.Username == "" {
		json.NewEncoder(w).Encode("Please login first")
	} else if r.Method == http.MethodPost {
		var newDocument ent.Document
		//parsing query string
		newDocument.Category = r.URL.Query().Get("category")
		newDocument.AttachmentID = r.URL.Query().Get("attachment")

		//find userID
		var userID int
		//asume there is one userID
		db.Model(&ent.User{}).Select("user_id").Where("username = ?", LoginUser.Username).Find(&userID)

		//find combination userID and loanID
		var laTmp ent.LoanApplication
		result := db.Model(&ent.LoanApplication{}).Where("user_id = ? AND loan_id = ?", userID, newDocument.AttachmentID).Find(&laTmp)
		//handling
		if result.RowsAffected <= 0 {
			json.NewEncoder(w).Encode("access denied")
			return
		}

		//get file
		file, header, errFile := r.FormFile("file")
		if errFile != nil {
			json.NewEncoder(w).Encode("Error on processing file")
			return
		}
		//gatau ngapain
		defer file.Close()

		//making directory
		filename := header.Filename
		newDocument.FileName = filename
		var f *os.File
		f, errFile = os.OpenFile("./static/document/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
		if errFile != nil {
			json.NewEncoder(w).Encode("Error on making file")
			return
		}
		defer f.Close()
		//copying file
		io.Copy(f, file)

		//create new record
		result = db.Create(&newDocument)
		if result.Error != nil {
			json.NewEncoder(w).Encode("error on inserting data")
			return
		}

		//change status of the loan to review
		result = db.Model(&ent.Loan{}).Where("lan = ?", newDocument.AttachmentID).Update("status", "Review")
		if result.Error != nil {
			json.NewEncoder(w).Encode("error on inserting data")
			return
		}
		json.NewEncoder(w).Encode(newDocument)
	} else {
		json.NewEncoder(w).Encode("Nothing happened")
	}
}

func DeleteLoan(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if LoginUser.Username == "" {
		json.NewEncoder(w).Encode("Please login first")
	} else if !LoginUser.Admin {
		json.NewEncoder(w).Encode("access denied")
	} else if r.Method == http.MethodDelete {
		var deletionStatus ent.DeletionLoan
		//parsing query string
		attachmentID := r.URL.Query().Get("id")

		//delete all documents that attached
		result := db.Where("attachment_id = ?", attachmentID).Delete(&ent.Document{})
		if result.Error == nil && result.RowsAffected > 0 {
			deletionStatus.Documents = true
		}

		//delete loan application
		result = db.Where("loan_id = ?", attachmentID).Delete(&ent.LoanApplication{})
		if result.Error == nil && result.RowsAffected > 0 {
			deletionStatus.LoanApplication = true
		}

		//delete loan
		result = db.Where("lan = ?", attachmentID).Delete(&ent.Loan{})
		if result.Error == nil && result.RowsAffected > 0 {
			deletionStatus.Loan = true
		}

		json.NewEncoder(w).Encode(deletionStatus)
	} else {
		json.NewEncoder(w).Encode("Nothing happened")
	}
}

func Approve(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if LoginUser.Username == "" {
		json.NewEncoder(w).Encode("Please login first")
	} else if r.Method == http.MethodPut {
		var newLoanApplication ent.LoanApplication
		//parsing query string
		newLoanApplication.LoanID = r.URL.Query().Get("id")

		//get the loan application
		db.Model(&ent.LoanApplication{}).Where("loan_id = ?", newLoanApplication.LoanID).Find(&newLoanApplication)
		if newLoanApplication.UserID <= 0 {
			json.NewEncoder(w).Encode("error on query")
			return
		}

		//admin
		var provable bool
		if LoginUser.Admin {
			//condition
			if newLoanApplication.LoanerSign != "" {
				provable = true
			}

			if provable {
				//save file
				result, filename := SaveFile(r)

				if !result {
					json.NewEncoder(w).Encode("error on processing file")
				} else {
					//update record
					newLoanApplication.LenderSign = filename
					err := db.Model(&ent.LoanApplication{}).Where("loan_id = ?", newLoanApplication.LoanID).Update("lender_sign", filename)
					if err.Error != nil {
						json.NewEncoder(w).Encode("error on updating")
						return
					} else {
						json.NewEncoder(w).Encode(newLoanApplication)
					}
				}
			} else {
				json.NewEncoder(w).Encode("requirements not met")
			}

		} else { //common user
			//get status loan
			var status string
			db.Model(&ent.Loan{}).Select("status").Where("lan = ?", newLoanApplication.LoanID).Find(&status)

			//get userID
			var userID int
			db.Model(&ent.User{}).Select("user_id").Where("username = ?", LoginUser.Username).Find(&userID)

			//condition
			if status == "Approved" && userID == newLoanApplication.UserID {
				provable = true
			}

			if provable {
				//save file
				result, filename := SaveFile(r)

				if !result {
					json.NewEncoder(w).Encode("error on processing file")
				} else {
					//update record
					newLoanApplication.LoanerSign = filename
					err := db.Model(&ent.LoanApplication{}).Where("loan_id = ?", newLoanApplication.LoanID).Update("loaner_sign", filename)
					if err.Error != nil {
						json.NewEncoder(w).Encode("error on updating")
						return
					} else {
						json.NewEncoder(w).Encode(newLoanApplication)
					}
				}
			} else {
				json.NewEncoder(w).Encode("requirements not met")
			}
		}
	} else {
		json.NewEncoder(w).Encode("Nothing happened")
	}
}

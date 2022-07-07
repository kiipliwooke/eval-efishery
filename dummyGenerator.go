package main

import (
	ent "eval-efishery/entity"
	"fmt"
)

func DummyGenerator() {
	var accounts = []ent.Account{
		{
			Username: "kiipliwooke",
			Password: "123",
			Admin:    false,
		},
		{
			Username: "phrye",
			Password: "iyvlinimut",
			Admin:    false,
		},
		{
			Username: "iyvlin",
			Password: "456",
			Admin:    true,
		},
	}
	var documents = []ent.Document{
		{
			Category:           "NPWP",
			FileName:           "doc1.pdf",
			VerificationStatus: false,
			AttachmentID:       "e3da5756-fdb6-11ec-b939-0242ac120002",
		},
		{
			Category:           "BPKB",
			FileName:           "doc2.pdf",
			VerificationStatus: true,
			AttachmentID:       "e3da5634-fdb6-11ec-b939-0242ac120002",
		},
		{
			Category:           "SKT",
			FileName:           "doc3.pdf",
			VerificationStatus: true,
			AttachmentID:       "e3da5756-fdb6-11ec-b939-0242ac120002",
		},
		{
			Category:           "KTP",
			FileName:           "doc4.pdf",
			VerificationStatus: true,
			AttachmentID:       "e3da5436-fdb6-11ec-b939-0242ac120002",
		},
	}
	var loans = []ent.Loan{
		{
			Lan:      "e3da5436-fdb6-11ec-b939-0242ac120002",
			Value:    120,
			Category: "Unspecified",
			Status:   "Unsufficient Document",
			Interest: 2.5,
			DueDate:  "2022-12-12",
		},
		{
			Lan:      "e3da5634-fdb6-11ec-b939-0242ac120002",
			Value:    140,
			Category: "Personal",
			Status:   "Approved",
			Interest: 2,
			DueDate:  "2023-07-07",
		},
		{
			Lan:      "e3da5756-fdb6-11ec-b939-0242ac120002",
			Value:    160,
			Category: "Unspecified",
			Status:   "Review",
			Interest: 5,
			DueDate:  "2024-01-01",
		},
	}
	var loanApplications = []ent.LoanApplication{
		{
			UserID: 1,
			LoanID: "e3da5436-fdb6-11ec-b939-0242ac120002",
			Date:   "2022-05-12",
		},
		{
			UserID: 1,
			LoanID: "e3da5634-fdb6-11ec-b939-0242ac120002",
			Date:   "2022-07-03",
		},
	}
	var users = []ent.User{
		{
			FullName:  "kipli z",
			BirthDate: "1995-01-01",
			Email:     "kipli@gmail.com",
			Phone:     "08112345678910",
			Username:  "kiipliwooke",
		},
		{
			FullName:  "phrye punya selera",
			BirthDate: "1990-12-13",
			Email:     "phrye@gmail.com",
			Phone:     "0899876543210",
			Username:  "phrye",
		},
	}
	result := db.Create(&accounts)
	if result.Error != nil {
		fmt.Println("Terjadi error saat generating dummy account")
		return
	}
	result = db.Create(&documents)
	if result.Error != nil {
		fmt.Println("Terjadi error saat generating dummy document")
		return
	}
	result = db.Create(&loans)
	if result.Error != nil {
		fmt.Println("Terjadi error saat generating dummy loan")
		return
	}
	result = db.Create(&loanApplications)
	if result.Error != nil {
		fmt.Println("Terjadi error saat generating dummy loan appplication")
		return
	}
	result = db.Create(&users)
	if result.Error != nil {
		fmt.Println("Terjadi error saat generating dummy user")
		return
	}
}

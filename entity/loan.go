package entity

type Loan struct {
	Lan      string  `gorm:"primaryKey"`
	Value    int     ``
	Category string  ``
	Status   string  ``
	Interest float64 ``
	DueDate  string  ``
}

type DetailedLoan struct {
	LoanerSign string
	LenderSign string
	Loan       Loan
	Documents  []Document
}

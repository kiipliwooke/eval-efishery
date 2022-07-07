package entity

type Loan struct {
	lan      string  `gorm:"primaryKey"`
	value    int     ``
	category string  ``
	interest float64 ``
	dueDate  string  ``
}

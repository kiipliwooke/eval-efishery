package entity

type Loan struct {
	Lan      string  `gorm:"primaryKey"`
	Value    int     ``
	Category string  ``
	Status   string  ``
	Interest float64 ``
	DueDate  string  ``
}

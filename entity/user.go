package entity

type User struct {
	userID    string `gorm:"primaryKey;autoIncrement"`
	fullName  string ``
	birthDate string ``
	email     string ``
	phone     string ``
	username  string ``
}

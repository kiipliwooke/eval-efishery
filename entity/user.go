package entity

type User struct {
	UserID    int    `gorm:"primaryKey;autoIncrement"`
	FullName  string ``
	BirthDate string ``
	Email     string ``
	Phone     string ``
	Username  string ``
}

package entity

type Account struct {
	username string `gorm:"primaryKey"`
	password string ``
	admin    bool   ``
}

type UserLogin struct {
	username string ``
	admin    bool   ``
}

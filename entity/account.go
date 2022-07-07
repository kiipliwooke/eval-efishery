package entity

type Account struct {
	Username string `json:"username" gorm:"primaryKey"`
	Password string ``
	Admin    bool   ``
}

type UserLogin struct {
	Username string ``
	Admin    bool   ``
}

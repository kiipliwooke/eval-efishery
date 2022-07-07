package entity

type Account struct {
	Username string `json:"username" gorm:"primaryKey"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

type UserLogin struct {
	Username string ``
	Admin    bool   ``
}

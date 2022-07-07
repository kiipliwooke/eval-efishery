package entity

type Document struct {
	documentID         int    `gorm:"autoIncrement;primaryKey"`
	category           string ``
	fileName           string ``
	verificationStatus bool   ``
	attachmentID       string ``
}

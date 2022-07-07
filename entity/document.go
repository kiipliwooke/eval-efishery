package entity

type Document struct {
	DocumentID         int    `gorm:"autoIncrement;primaryKey"`
	Category           string ``
	FileName           string ``
	VerificationStatus bool   ``
	AttachmentID       string ``
}

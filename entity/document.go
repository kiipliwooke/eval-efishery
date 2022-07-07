package entity

type Document struct {
	DocumentID         int    `gorm:"autoIncrement;primaryKey"`
	Category           string `json:"category"`
	FileName           string ``
	VerificationStatus bool   ``
	AttachmentID       string `json:"attachment"`
}

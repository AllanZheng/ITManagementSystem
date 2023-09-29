package model

type FeeRecord struct {
	ID         uint    `gorm:"primary_key" json:"id"`
	Company    string  `gorm:"company" json:"company"`
	FeeType    string  `gorm:"feetype" json:"type"`
	Content    string  `gorm:"content" json:"content"`
	Date       int     `gorm:"date" json:"Date"`
	Amount     float32 `gorm:"amount" json:"amount"`
	Operator   string  `gorm:"operator" json:"operator"`
	Comment    string  `gorm:"comment" json:"comment"`
	CreatedOn  int     `json:"created_on"`
	ModifiedOn int     `json:"modified_on"`
	DeletedOn  int     `json:"deleted_on"`
}

type Query struct {
	Id uint `"json:Id"`
}

type BadRecord struct {
	ID         uint    `gorm:"primary_key" json:"id"`
	Company    string  `gorm:"company" json:"company"`
	FeeType    string  `gorm:"feetype" json:"type"`
	Content    string  `gorm:"content" json:"content"`
	Date       int     `gorm:"date" json:"Date"`
	Amount     float32 `gorm:"amount" json:"amount"`
	Operator   string  `gorm:"operator" json:"operator"`
	Comment    string  `gorm:"comment" json:"comment"`
	CreatedOn  int     `json:"created_on"`
	ModifiedOn int     `json:"modified_on"`
	DeletedOn  int     `json:"deleted_on"`
}

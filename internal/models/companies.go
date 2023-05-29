package models

type Company struct {
	Shared

	BusinessName string `gorm:"unique"`
	TradeName    string
	CNPJ         string `gorm:"unique"`
	Email        string `gorm:"unique"`
	Phone        string
}

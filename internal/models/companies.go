package models

import "github.com/google/uuid"

type Company struct {
	Shared

	LegalName string `gorm:"unique"`
	TradeName string
	CNPJ      string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Phone     string
}

type CompanyContentRecommendation struct {
	Shared

	MBTIType string `gorm:"index:idx_mbti_type;not null"`

	CompanyID uuid.UUID
	Company   Company `gorm:"foreignKey:CompanyID"`

	CategoryID uuid.UUID
	Category   Category `gorm:"foreignKey:CategoryID"`
}

type DefaultContentRecommendation struct {
	Shared

	MBTIType string `gorm:"index:idx_mbti_type;not null"`

	CategoryID uuid.UUID
	Category   Category `gorm:"foreignKey:CategoryID"`
}
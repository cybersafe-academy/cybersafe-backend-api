package db

import (
	"cybersafe-backend-api/internal/models"

	"gorm.io/gorm"
)

type DBStorer struct {
	Conn *gorm.DB
}

func (db *DBStorer) GetUserByCPF(cpf string) (models.User, error) {
	user := models.User{}

	result := db.Conn.Where("CPF = ?", cpf).First(&user)

	return user, result.Error

}

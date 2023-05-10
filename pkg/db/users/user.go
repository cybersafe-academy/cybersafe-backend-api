package users

import (
	"cybersafe-backend-api/internal/models"

	"gorm.io/gorm"
)

type UserDB struct {
	Conn *gorm.DB
}

func (ub *UserDB) GetByCPF(cpf string) (models.User, error) {
	user := models.User{}

	result := ub.Conn.Where("CPF = ?", cpf).First(&user)

	return user, result.Error

}

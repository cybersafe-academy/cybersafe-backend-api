package users

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UsersManagerDB struct {
	DBConnection *gorm.DB
}

func (ub *UsersManagerDB) GetByCPF(cpf string) (models.User, error) {
	user := models.User{}
	result := ub.DBConnection.Where("CPF = ?", cpf).First(&user)
	return user, result.Error

}

func (ub *UsersManagerDB) GetByID(id uuid.UUID) (models.User, error) {
	var user models.User
	result := ub.DBConnection.First(&user, id)
	return user, result.Error
}

func (ub *UsersManagerDB) List(offset, limit int) ([]models.User, int64) {
	var users []models.User
	ub.DBConnection.Offset(offset).Limit(limit).Find(&users)

	var count int64
	ub.DBConnection.Model(&models.User{}).Count(&count)
	return users, count
}

func (ub *UsersManagerDB) Create(user *models.User) error {
	result := ub.DBConnection.Create(user)
	return result.Error
}

func (ub *UsersManagerDB) Delete(id uuid.UUID) error {
	result := ub.DBConnection.Delete(&models.User{}, id)
	return result.Error
}

func (ub *UsersManagerDB) Update(user *models.User) (int, error) {
	result := ub.DBConnection.Model(user).Clauses(clause.Returning{}).Updates(user)
	return int(result.RowsAffected), result.Error
}

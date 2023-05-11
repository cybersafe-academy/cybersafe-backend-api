package users

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserDB struct {
	DBConnection *gorm.DB
}

func (ub *UserDB) GetByCPF(cpf string) (models.User, error) {
	user := models.User{}
	result := ub.DBConnection.Where("CPF = ?", cpf).First(&user)
	return user, result.Error

}

func (ub *UserDB) GetByID(id uuid.UUID) (models.User, error) {
	var user models.User
	result := ub.DBConnection.First(&user, id)
	return user, result.Error
}

func (ub *UserDB) List(offset, limit int) ([]models.User, int64) {
	var users []models.User
	ub.DBConnection.Offset(offset).Limit(limit).Find(&users)

	var count int64
	ub.DBConnection.Model(&models.User{}).Count(&count)
	return users, count
}

func (ub *UserDB) Create(user *models.User) error {
	result := ub.DBConnection.Create(user)
	return result.Error
}

func (ub *UserDB) Delete(id uuid.UUID) error {
	result := ub.DBConnection.Delete(&models.User{}, id)
	return result.Error
}

func (ub *UserDB) Update(user *models.User) (int, error) {
	result := ub.DBConnection.Model(user).Clauses(clause.Returning{}).Updates(user)
	return int(result.RowsAffected), result.Error
}

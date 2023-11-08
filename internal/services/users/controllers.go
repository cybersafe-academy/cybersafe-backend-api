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

func (um *UsersManagerDB) GetByCPF(cpf string) (models.User, error) {
	user := models.User{}
	result := um.DBConnection.Where("CPF = ?", cpf).Preload(clause.Associations).First(&user)
	return user, result.Error
}

func (um *UsersManagerDB) GetByID(id uuid.UUID) (models.User, error) {
	var user models.User
	result := um.DBConnection.Preload(clause.Associations).First(&user, id)
	return user, result.Error
}

func (um *UsersManagerDB) ListWithPagination(offset, limit int) ([]models.User, int64) {
	var users []models.User
	um.DBConnection.Offset(offset).Limit(limit).Preload(clause.Associations).Find(&users)

	var count int64
	um.DBConnection.Model(&models.User{}).Count(&count)
	return users, count
}

func (um *UsersManagerDB) Create(user *models.User) error {
	result := um.DBConnection.Create(user)
	return result.Error
}

func (um *UsersManagerDB) Delete(id uuid.UUID) error {
	result := um.DBConnection.Delete(&models.User{}, id)
	return result.Error
}

func (um *UsersManagerDB) Update(user *models.User) (int, error) {
	result := um.DBConnection.Model(user).
		Where("email = ?", user.Email).
		Clauses(clause.Returning{}).
		Updates(user)

	return int(result.RowsAffected), result.Error
}

func (um *UsersManagerDB) ExistsByEmail(email string) bool {
	result := um.DBConnection.Where("email = ?", email).First(&models.User{})
	return result.Error == nil
}

func (um *UsersManagerDB) ExistsDisabledByEmail(cpf string) bool {
	result := um.DBConnection.Where("email = ?", cpf).Where("enabled = false").First(&models.User{})
	return result.Error == nil
}

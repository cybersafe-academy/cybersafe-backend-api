package users

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
)

type UsersDBMock struct {
	GetByCPFMock func(string) (models.User, error)
	ListMock     func(int, int) ([]models.User, int64)
	GetByIDMock  func(uuid.UUID) (models.User, error)
	CreateMock   func(*models.User) error
	DeleteMock   func(uuid.UUID) error
	UpdateMock   func(*models.User) (int, error)
}

func (ubm *UsersDBMock) GetByCPF(cpf string) (models.User, error) {
	return ubm.GetByCPFMock(cpf)
}
func (ubm *UsersDBMock) List(offset, limit int) ([]models.User, int64) {
	return ubm.ListMock(offset, limit)
}
func (ubm *UsersDBMock) GetByID(id uuid.UUID) (models.User, error) {
	return ubm.GetByIDMock(id)
}
func (ubm *UsersDBMock) Create(user *models.User) error {
	return ubm.CreateMock(user)
}
func (ubm *UsersDBMock) Delete(id uuid.UUID) error {
	return ubm.DeleteMock(id)
}
func (ubm *UsersDBMock) Update(user *models.User) (int, error) {
	return ubm.UpdateMock(user)
}

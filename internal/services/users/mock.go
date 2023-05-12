package users

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
)

type UsersManagerMock struct {
	GetByCPFMock           func(string) (models.User, error)
	ListWithPaginationMock func(int, int) ([]models.User, int64)
	GetByIDMock            func(uuid.UUID) (models.User, error)
	CreateMock             func(*models.User) error
	DeleteMock             func(uuid.UUID) error
	UpdateMock             func(*models.User) (int, error)
	ExistsByEmailMock      func(string) bool
}

func (umm *UsersManagerMock) GetByCPF(cpf string) (models.User, error) {
	return umm.GetByCPFMock(cpf)
}

func (umm *UsersManagerMock) ListWithPagination(offset, limit int) ([]models.User, int64) {
	return umm.ListWithPaginationMock(offset, limit)
}

func (umm *UsersManagerMock) GetByID(id uuid.UUID) (models.User, error) {
	return umm.GetByIDMock(id)
}

func (umm *UsersManagerMock) Create(user *models.User) error {
	return umm.CreateMock(user)
}

func (umm *UsersManagerMock) Delete(id uuid.UUID) error {
	return umm.DeleteMock(id)
}

func (umm *UsersManagerMock) Update(user *models.User) (int, error) {
	return umm.UpdateMock(user)
}

func (umm *UsersManagerMock) ExistsByEmail(email string) bool {
	return umm.ExistsByEmailMock(email)
}

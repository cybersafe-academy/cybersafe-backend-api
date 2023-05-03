package db

import (
	"cybersafe-backend-api/internal/models"
)

type MockStorer struct {
	GetUserByCPFMock func(string) (models.User, error)
}

func (ms *MockStorer) GetUserByCPF(cpf string) (models.User, error) {
	return ms.GetUserByCPFMock(cpf)
}

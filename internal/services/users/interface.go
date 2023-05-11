package users

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
)

type Users interface {
	GetByCPF(string) (models.User, error)
	List(int, int) ([]models.User, int64)
	GetByID(uuid.UUID) (models.User, error)
	Create(*models.User) error
	Delete(uuid.UUID) error
	Update(*models.User) (int, error)
}

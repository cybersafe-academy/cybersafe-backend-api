package users

import "cybersafe-backend-api/internal/models"

type Users interface {
	GetByCPF(string) (models.User, error)
}

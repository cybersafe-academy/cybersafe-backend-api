package db

import "cybersafe-backend-api/internal/models"

type Storer interface {
	GetUserByCPF(string) (models.User, error)
}

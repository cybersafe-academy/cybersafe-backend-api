package users

import (
	"cybersafe-backend-api/internal/models"
	"time"
)

func ToListResponse(users []models.User) []ResponseContent {

	var usersResponse []ResponseContent

	for _, user := range users {
		usersResponse = append(usersResponse, ToResponse(user))
	}

	return usersResponse
}

func ToResponse(user models.User) ResponseContent {
	return ResponseContent{
		UserFields: UserFields{
			Name:      user.Name,
			BirthDate: user.BirthDate.Truncate(24 * time.Hour).String(),
			CPF:       user.CPF,
			Role:      user.Role,
			Email:     user.Email,
		},
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

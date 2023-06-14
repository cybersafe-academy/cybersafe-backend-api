package users

import (
	"cybersafe-backend-api/internal/models"
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
			BirthDate: user.BirthDate.Format("2006-01-02"),
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

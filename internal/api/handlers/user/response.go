package user

import "cybersafe-backend-api/internal/models"

func ToListResponse(users []models.User) []ResponseContent {

	var usersResponse []ResponseContent

	for _, user := range users {
		usersResponse = append(usersResponse, ResponseContent{
			ID: user.ID,
			UserFields: UserFields{
				Name: user.Name,
				Age:  user.Age,
				CPF:  user.CPF,
			},
		})
	}

	return usersResponse
}

func ToResponse(user models.User) ResponseContent {
	return ResponseContent{
		UserFields: UserFields{
			Name: user.Name,
			Age:  user.Age,
			CPF:  user.CPF,
			// Role:  user.Role,
			Email: user.Email,
		},
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

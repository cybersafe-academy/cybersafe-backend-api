package user

import "cybersafe-backend-api/pkg/models"

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

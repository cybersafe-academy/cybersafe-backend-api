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
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		UserFields: UserFields{
			Name:              user.Name,
			Role:              user.Role,
			Email:             user.Email,
			BirthDate:         user.BirthDate.Format("2006-01-02"),
			CPF:               user.CPF,
			ProfilePictureURL: user.ProfilePictureURL,
		},
		CompanyResponse: CompanyResponse{
			ID:        user.CompanyID,
			LegalName: user.Company.LegalName,
			TradeName: user.Company.TradeName,
			CNPJ:      user.Company.CNPJ,
			Email:     user.Company.Email,
			Phone:     user.Company.Phone,
		},
	}
}

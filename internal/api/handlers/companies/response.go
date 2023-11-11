package companies

import (
	"cybersafe-backend-api/internal/models"
)

func ToResponse(company models.Company) ResponseContent {
	return ResponseContent{
		CompanyFields: CompanyFields{
			LegalName: company.LegalName,
			TradeName: company.TradeName,
			CNPJ:      company.CNPJ,
			Email:     company.Email,
			Phone:     company.Phone,
		},
		ID:        company.ID,
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
		DeletedAt: company.DeletedAt,
	}
}

func ToListResponse(companies []models.Company) []ResponseContent {
	var companiesResponse []ResponseContent

	for _, company := range companies {
		companyResponse := ResponseContent{
			CompanyFields: CompanyFields{
				LegalName: company.LegalName,
				TradeName: company.TradeName,
				CNPJ:      company.CNPJ,
				Email:     company.Email,
				Phone:     company.Phone,
			},
			ID:        company.ID,
			CreatedAt: company.CreatedAt,
			UpdatedAt: company.UpdatedAt,
			DeletedAt: company.DeletedAt,
		}

		companiesResponse = append(companiesResponse, companyResponse)
	}

	return companiesResponse
}

func ToCompanyContentRecommendationResponse(recommendations []models.CompanyContentRecommendation) CompanyContentRecommendationResponseContent {
	response := CompanyContentRecommendationResponseContent{
		ID:        recommendations[0].ID,
		CompanyID: recommendations[0].CompanyID,
		CompanyContentRecommendationFields: CompanyContentRecommendationFields{
			MBTIType: recommendations[0].MBTIType,
		},
		CreatedAt: recommendations[0].CreatedAt,
		UpdatedAt: recommendations[0].UpdatedAt,
		DeletedAt: recommendations[0].DeletedAt,
	}

	var categories []string
	for _, recommendation := range recommendations {
		categories = append(categories, recommendation.CategoryID.String())
	}

	response.Categories = categories

	return response
}

func ToCompanyContentRecommendationByMBTIResponse(recommendations []models.CompanyContentRecommendation) CompanyContentRecommendationByMBTIResponseContent {
	response := CompanyContentRecommendationByMBTIResponseContent{
		ID:        recommendations[0].ID,
		CompanyID: recommendations[0].CompanyID,
		MbtiType:  recommendations[0].MBTIType,
		CreatedAt: recommendations[0].CreatedAt,
		UpdatedAt: recommendations[0].UpdatedAt,
		DeletedAt: recommendations[0].DeletedAt,
	}

	var categories []CategoryResponse
	for _, recommendation := range recommendations {
		category := CategoryResponse{
			ID:   recommendation.CategoryID,
			Name: recommendation.Category.Name,
		}

		categories = append(categories, category)
	}

	response.Categories = categories

	return response
}

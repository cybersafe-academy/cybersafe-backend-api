package companies

import "cybersafe-backend-api/internal/models"

func ToResponse(company models.Company) ResponseContent {
	return ResponseContent{
		CompanyFields: CompanyFields{
			LegalName:  company.LegalName,
			TradeName:  company.TradeName,
			CNPJ:       company.CNPJ,
			Email:      company.Email,
			WebsiteURL: company.WebsiteURL,
			Phone:      company.Phone,
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
				LegalName:  company.LegalName,
				TradeName:  company.TradeName,
				CNPJ:       company.CNPJ,
				Email:      company.Email,
				WebsiteURL: company.WebsiteURL,
				Phone:      company.Phone,
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

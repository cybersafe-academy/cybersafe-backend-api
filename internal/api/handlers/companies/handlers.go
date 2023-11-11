package companies

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/server/middlewares"
	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/pagination"
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

// ListCompaniesHandler
//
//	@Summary	List companies with paginated response
//
//	@Tags		Company
//	@success	200		{array}	pagination.PaginationData{data=ResponseContent}	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		page	query		int					false	"Page number"
//	@Param		limit	query		int					false	"Limit of elements per page"
//	@Router		/companies [get]
//	@Security	Bearer
//	@Security	Language
func ListCompaniesHandler(c *components.HTTPComponents) {
	paginationData, err := pagination.GetPaginationData(c.HttpRequest.URL.Query())

	if errors.Is(err, errutil.ErrInvalidPageParam) {
		components.HttpErrorResponse(c, http.StatusNotFound, err)
		return
	} else if errors.Is(err, errutil.ErrInvalidLimitParam) {
		components.HttpErrorResponse(c, http.StatusNotFound, err)
		return
	}

	currentUser, ok := c.HttpRequest.Context().Value(middlewares.UserKey).(*models.User)
	if !ok {
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrUnexpectedError",
		}))
		return
	}

	var companies []models.Company
	var count int

	if currentUser.Role == models.MasterUserRole {
		companies, count = c.Components.Resources.Companies.ListWithPagination(paginationData.Offset, paginationData.Limit)
	} else {
		company, err := c.Components.Resources.Companies.GetByID(currentUser.CompanyID)
		if err != nil {
			components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrUnexpectedError",
			}))
			return
		}

		companies = append(companies, company)
		count = 1
	}

	response := paginationData.ToResponse(
		ToListResponse(companies), int(count),
	)

	components.HttpResponseWithPayload(c, response, http.StatusOK)
}

// GetCompanyByIDHandler retrieves a company by ID
//
//	@Summary	Get company by ID
//	@Tags		Company
//	@Param		id		path		string			true	"ID of the company to be retrieved"
//	@Success	200		{object}	ResponseContent	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Router		/companies/{id} [get]
//	@Security	Bearer
//	@Security	Language
func GetCompanyByIdHandler(c *components.HTTPComponents) {
	id := chi.URLParam(c.HttpRequest, "id")

	if !govalidator.IsUUID(id) {
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest, c.Components.Localizer.MustLocalize(
			&i18n.LocalizeConfig{
				MessageID: "ErrInvalidUUID",
			},
		))
		return
	}

	company, err := c.Components.Resources.Companies.GetByID(uuid.MustParse(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			components.HttpErrorLocalizedResponse(c, http.StatusNotFound, c.Components.Localizer.MustLocalize(
				&i18n.LocalizeConfig{
					MessageID: "ErrCompanyResourceNotFound",
				},
			))
			return
		} else {
			components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(
				&i18n.LocalizeConfig{
					MessageID: "ErrUnexpectedError",
				},
			))
			return
		}
	}

	components.HttpResponseWithPayload(c, ToResponse(company), http.StatusOK)
}

// CreateCompanyHandler
//
//	@Summary	Create a company
//
//	@Tags		Company
//	@Success	200		{object}	ResponseContent	"OK"
//	@Failure	409		"Conflict"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		request	body		RequestContent		true	"Request payload for creating a new company"
//	@Router		/companies [post]
//	@Security	Bearer
//	@Security	Language
func CreateCompanyHandler(c *components.HTTPComponents) {
	var requestContent RequestContent
	err := components.ValidateRequest(c, &requestContent)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	company := requestContent.ToEntity()
	err = c.Components.Resources.Companies.Create(company)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			components.HttpErrorLocalizedResponse(c, http.StatusConflict, c.Components.Localizer.MustLocalize(
				&i18n.LocalizeConfig{
					MessageID: "ErrCNPJorEmailorBusinessNameAlreadyInUse",
				},
			))
			return
		} else {
			components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(
				&i18n.LocalizeConfig{
					MessageID: "ErrUnexpectedError",
				},
			))
			return
		}
	}

	components.HttpResponseWithPayload(c, ToResponse(*company), http.StatusOK)
}

// UpdateCompanyHandler
//
//	@Summary	Update company by ID
//
//	@Tags		Company
//	@success	200		{object}	ResponseContent	"OK"
//	@Failure	400		"Bad Request"
//	@Failure	404		"Company not found"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		request	body		RequestContent		true	"Request payload for updating company information"
//	@Param		id		path		string				true	"ID of company to be updated"
//	@Router		/companies/{id} [put]
//	@Security	Bearer
//	@Security	Language
func UpdateCompanyHandler(c *components.HTTPComponents) {
	companyRequest := RequestContent{}
	err := components.ValidateRequest(c, &companyRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	id := chi.URLParam(c.HttpRequest, "id")

	if !govalidator.IsUUID(id) {
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest, c.Components.Localizer.MustLocalize(
			&i18n.LocalizeConfig{
				MessageID: "ErrInvalidUUID",
			},
		))
		return
	}

	company := companyRequest.ToEntity()

	company.ID = uuid.MustParse(id)
	rowsAffected, err := c.Components.Resources.Companies.Update(company)

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			components.HttpErrorLocalizedResponse(c, http.StatusConflict, c.Components.Localizer.MustLocalize(
				&i18n.LocalizeConfig{
					MessageID: "ErrCNPJorEmailorBusinessNameAlreadyInUse",
				},
			))
			return
		} else {
			components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(
				&i18n.LocalizeConfig{
					MessageID: "ErrUnexpectedError",
				},
			))
			return
		}
	}
	if rowsAffected == 0 {
		components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(
			&i18n.LocalizeConfig{
				MessageID: "ErrCompanyResourceNotFound",
			},
		))
		return
	}

	components.HttpResponseWithPayload(c, ToResponse(*company), http.StatusOK)
}

// DeleteCompanyHandler
//
//	@Summary	Delete a company by ID
//
//	@Tags		Company
//	@Success	204		"No content"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		id		path		string				true	"ID of the company to be deleted"
//	@Router		/companies/{id} [delete]
//	@Security	Bearer
//	@Security	Language
func DeleteCompanyHandler(c *components.HTTPComponents) {
	id := chi.URLParam(c.HttpRequest, "id")

	if !govalidator.IsUUID(id) {
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest, c.Components.Localizer.MustLocalize(
			&i18n.LocalizeConfig{
				MessageID: "ErrInvalidUUID",
			},
		))
		return
	}

	err := c.Components.Resources.Companies.Delete(uuid.MustParse(id))

	if err != nil {
		components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(
			&i18n.LocalizeConfig{
				MessageID: "ErrUnexpectedError",
			},
		))
		return
	}

	components.HttpResponse(c, http.StatusNoContent)
}

// UpdateCompanyContentRecommendationsHandler
//
//	@Summary	Update company content recommendations for a given MBTI type
//
//	@Tags		Company
//	@Success	200		{object}	ResponseContent	"OK"
//	@Failure	400		"Bad Request"
//	@Failure	404		"Company not found"
//	@Response	default	{object}	components.Response							"Standard error example object"
//	@Param		request	body		CompanyContentRecommendationRequestContent	true	"Request payload for updating company content recommendations"
//	@Param		id		path		string										true	"ID of company to be updated"
//	@Router		/companies/{id}/content-recommendations [put]
//	@Security	Bearer
//	@Security	Language
func UpdateCompanyContentRecommendationsHandler(c *components.HTTPComponents) {
	var requestContent CompanyContentRecommendationRequestContent
	err := components.ValidateRequest(c, &requestContent)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	id := chi.URLParam(c.HttpRequest, "id")
	if !govalidator.IsUUID(id) {
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest, c.Components.Localizer.MustLocalize(
			&i18n.LocalizeConfig{
				MessageID: "ErrInvalidUUID",
			},
		))
		return
	}

	_, err = c.Components.Resources.Companies.GetByID(uuid.MustParse(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			components.HttpErrorLocalizedResponse(c, http.StatusNotFound, c.Components.Localizer.MustLocalize(
				&i18n.LocalizeConfig{
					MessageID: "ErrCompanyResourceNotFound",
				},
			))
			return
		} else {
			components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(
				&i18n.LocalizeConfig{
					MessageID: "ErrUnexpectedError",
				},
			))
			return
		}
	}

	user, ok := c.HttpRequest.Context().Value(middlewares.UserKey).(*models.User)
	if !ok {
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrUnexpectedError",
		}))
		return
	}

	if (user.Role != models.MasterUserRole) && (user.CompanyID.String() != id) {
		components.HttpErrorLocalizedResponse(c, http.StatusUnauthorized, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrInsufficientPermission",
		}))
		return
	}

	requestContent.CompanyID = uuid.MustParse(id)
	contentRecommendations := requestContent.ToEntity()

	err = c.Components.Resources.Companies.UpdateContentRecommendations(contentRecommendations)
	if err != nil {
		components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(
			&i18n.LocalizeConfig{
				MessageID: "ErrUnexpectedError",
			},
		))
		return
	}

	components.HttpResponseWithPayload(c, ToCompanyContentRecommendationResponse(contentRecommendations), http.StatusOK)
}

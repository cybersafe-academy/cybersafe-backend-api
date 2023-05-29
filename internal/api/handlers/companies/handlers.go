package companies

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/pkg/errutil"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

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
			components.HttpErrorResponse(c, http.StatusConflict, errutil.ErrCNPJorEmailorBusinessNameAlreadyInUse)
			return
		} else {
			components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
			return
		}
	}

	components.HttpResponseWithPayload(c, *company, http.StatusOK)
}

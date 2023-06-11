package companies

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/pagination"
	"errors"
	"net/http"

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

	companies, count := c.Components.Resources.Companies.ListWithPagination(paginationData.Offset, paginationData.Limit)

	response := paginationData.ToResponse(
		ToListResponse(companies), int(count),
	)

	components.HttpResponseWithPayload(c, response, http.StatusOK)
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
			components.HttpErrorResponse(c, http.StatusConflict, errutil.ErrCNPJorEmailorBusinessNameAlreadyInUse)
			return
		} else {
			components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
			return
		}
	}

	components.HttpResponseWithPayload(c, ToResponse(*company), http.StatusOK)
}

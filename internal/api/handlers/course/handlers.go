package course

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/pkg/db"

	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/pagination"
	"errors"
	"net/http"
)

// ListCoursesHandler
//
//	@Summary	List courses with paginated response
//
//	@Tags		Courses
//	@success	200		{array}	pagination.PaginationData{data=ResponseContent}	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		page	query		int					false	"Page number"
//	@Param		limit	query		int					false	"Limit of elements per page"
//	@Router		/courses [get]
//	@Security	Bearer
//	@Security	Language
func ListCoursesHandler(c *components.HTTPComponents) {

	dbConn := db.MustGetDbConn()

	paginationData, err := pagination.GetPaginationData(c.HttpRequest.URL.Query())

	if errors.Is(err, errutil.ErrInvalidPageParam) {
		components.HttpErrorResponse(c, http.StatusNotFound, err)
		return
	} else if errors.Is(err, errutil.ErrInvalidLimitParam) {
		components.HttpErrorResponse(c, http.StatusNotFound, err)
		return
	}

	var courses []models.Course
	dbConn.Offset(paginationData.Offset).Limit(paginationData.Limit).Find(&courses)

	var count int64
	dbConn.Model(&models.Course{}).Count(&count)

	response := paginationData.ToResponse(
		ToListResponse(courses), int(count),
	)

	components.HttpResponseWithPayload(c, response, http.StatusOK)
}

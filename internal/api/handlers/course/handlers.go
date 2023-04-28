package course

import (
	"cybersafe-backend-api/internal/api/components"

	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/pagination"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ListCoursesHandler
//
//	@Summary	List courses with paginated response
//
//	@Tags		Course
//	@success	200		{array}	pagination.PaginationData{data=ResponseContent}	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		page	query		int					false	"Page number"
//	@Param		limit	query		int					false	"Limit of elements per page"
//	@Router		/courses [get]
//	@Security	Bearer
//	@Security	Language
func ListCoursesHandler(c *components.HTTPComponents) {

	dbConn := c.Components.DB

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

// GetCourseByID retrieves a course by ID
//
//	@Summary	Get course by ID
//	@Tags		Course
//	@Param		id		path		string			true	"ID of the course to be retrieved"
//	@Success	200		{object}	ResponseContent	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Router		/courses/{id} [get]
//	@Security	Bearer
//	@Security	Language
func GetCourseByID(c *components.HTTPComponents) {
	id := chi.URLParam(c.HttpRequest, "id")

	dbConn := c.Components.DB

	var course models.Course
	result := dbConn.First(&course, uuid.MustParse(id))

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			components.HttpErrorResponse(c, http.StatusNotFound, errutil.ErrUserResourceNotFound)
			return
		} else {
			components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
			return
		}
	}

	components.HttpResponseWithPayload(c, ToResponse(course), http.StatusOK)
}

// CreateCourseHandler
//
//	@Summary	Create a course
//
//	@Tags		Course
//	@Success	200		{object}	ResponseContent	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		request	body		RequestContent		true	"Request payload for creating a new course"
//	@Router		/courses [post]
//	@Security	Bearer
//	@Security	Language
func CreateCourseHandler(c *components.HTTPComponents) {
	courseRequest := RequestContent{}
	err := components.ValidateRequest(c, &courseRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	course := courseRequest.ToEntity()
	dbConn := c.Components.DB

	result := dbConn.Create(course)

	if result.Error != nil {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}

	components.HttpResponseWithPayload(c, ToResponse(*course), http.StatusOK)
}

// DeleteCourseHandler
//
//	@Summary	Delete a course by ID
//
//	@Tags		Course
//	@success	204		"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		id		path		string				true	"ID of the course to be deleted"
//	@Router		/courses/{id} [delete]
//	@Security	Bearer
//	@Security	Language
func DeleteCourseHandler(c *components.HTTPComponents) {
	id := chi.URLParam(c.HttpRequest, "id")

	dbConn := c.Components.DB

	result := dbConn.Delete(&models.Course{}, uuid.MustParse(id))

	if result.Error != nil {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}

	components.HttpResponse(c, http.StatusNoContent)
}

// UpdateCourseHandler
//
//	@Summary	Update course by ID
//
//	@Tags		Course
//	@success	200		{object}	ResponseContent	"OK"
//	@Failure	400		"Bad Request"
//	@Failure	404		"Course not found"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		request	body		RequestContent		true	"Request payload for updating course information"
//	@Param		id		path		string				true	"ID of course to be updated"
//	@Router		/courses/{id} [put]
//	@Security	Bearer
//	@Security	Language
func UpdateCourseHandler(c *components.HTTPComponents) {
	courseRequest := RequestContent{}
	err := components.ValidateRequest(c, &courseRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	dbConn := c.Components.DB
	id := chi.URLParam(c.HttpRequest, "id")

	course := &models.Course{}
	result := dbConn.First(course, id)

	if result.Error != nil {
		components.HttpErrorResponse(c, http.StatusNotFound, errutil.ErrUserResourceNotFound)
		return
	}

	updatedCourse := courseRequest.ToEntity()
	result = dbConn.Model(course).Updates(updatedCourse)

	if result.Error != nil {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}

	components.HttpResponseWithPayload(c, ToResponse(*updatedCourse), http.StatusOK)
}

package courses

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
	paginationData, err := pagination.GetPaginationData(c.HttpRequest.URL.Query())

	if errors.Is(err, errutil.ErrInvalidPageParam) {
		components.HttpErrorResponse(c, http.StatusNotFound, err)
		return
	} else if errors.Is(err, errutil.ErrInvalidLimitParam) {
		components.HttpErrorResponse(c, http.StatusNotFound, err)
		return
	}

	courses, count := c.Components.Resources.Courses.ListWithPagination(paginationData.Offset, paginationData.Limit)

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

	if !govalidator.IsUUID(id) {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrInvalidUUID)
		return
	}

	course, err := c.Components.Resources.Courses.GetByID(uuid.MustParse(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			components.HttpErrorResponse(c, http.StatusNotFound, errutil.ErrCourseResourceNotFound)
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

	err = c.Components.Resources.Courses.Create(course)

	if err != nil {
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

	if !govalidator.IsUUID(id) {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrInvalidUUID)
		return
	}

	err := c.Components.Resources.Courses.Delete(uuid.MustParse(id))

	if err != nil {
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

	id := chi.URLParam(c.HttpRequest, "id")

	if !govalidator.IsUUID(id) {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrInvalidUUID)
		return
	}

	course := courseRequest.ToEntity()
	course.ID = uuid.MustParse(id)

	rowsAffected, err := c.Components.Resources.Courses.Update(course)

	if err != nil {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}
	if rowsAffected == 0 {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrCourseResourceNotFound)
		return
	}

	components.HttpResponseWithPayload(c, ToResponse(*course), http.StatusOK)
}

// CreateCourseReview
//
//	@Summary	Create review
//
//	@Tags		Course
//	@Success	200		{object}	ReviewResponse	"OK"
//	@Failure	409		"Conflict"
//	@Response	default	{object}	components.Response		"Standard error example object"
//	@Param		request	body		ReviewRequestContent	true	"Request payload for creating a review"
//	@Router		/courses/{id}/review [post]
//	@Security	Bearer
//	@Security	Language
func CreateCourseReview(c *components.HTTPComponents) {

	user, ok := c.HttpRequest.Context().Value(middlewares.UserKey).(*models.User)

	if !ok {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}

	var requestContent ReviewRequestContent
	err := components.ValidateRequest(c, &requestContent)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	review := requestContent.ToEntityReview()
	review.UserID = user.ID

	err = c.Components.Resources.Reviews.Create(review)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			components.HttpErrorResponse(c, http.StatusConflict, errutil.ErrReviewAlreadyExists)
			return
		} else {
			components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
			return
		}
	}
	components.HttpResponseWithPayload(c, ToReviewResponse(*review), http.StatusOK)
}

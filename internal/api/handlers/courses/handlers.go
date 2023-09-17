package courses

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/handlers/courses/httpmodels"
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
//	@success	200		{array}	pagination.PaginationData{data=httpmodels.ResponseContent}	"OK"
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

// ListCoursesHandler
//
//	@Summary	List all courses grouped by category
//
//	@Tags		Course
//	@success	200		{object}	httpmodels.CourseByCategoryResponse	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Router		/courses [get]
//	@Security	Bearer
//	@Security	Language
func FetchCoursesHandler(c *components.HTTPComponents) {
	courses := c.Components.Resources.Courses.ListByCategory()

	components.HttpResponseWithPayload(c, courses, http.StatusOK)
}

// GetCourseByID retrieves a course by ID
//
//	@Summary	Get course by ID
//	@Tags		Course
//	@Param		id		path		string						true	"ID of the course to be retrieved"
//	@Success	200		{object}	httpmodels.ResponseContent	"OK"
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
//	@Success	200		{object}	httpmodels.ResponseContent	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response			"Standard error example object"
//	@Param		request	body		httpmodels.RequestContent	true	"Request payload for creating a new course"
//	@Router		/courses [post]
//	@Security	Bearer
//	@Security	Language
func CreateCourseHandler(c *components.HTTPComponents) {
	courseRequest := httpmodels.RequestContent{}
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
//	@success	204		"No Content"
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
//	@success	200		{object}	httpmodels.ResponseContent	"OK"
//	@Failure	400		"Bad Request"
//	@Failure	404		"Course not found"
//	@Response	default	{object}	components.Response			"Standard error example object"
//	@Param		request	body		httpmodels.RequestContent	true	"Request payload for updating course information"
//	@Param		id		path		string						true	"ID of course to be updated"
//	@Router		/courses/{id} [put]
//	@Security	Bearer
//	@Security	Language
func UpdateCourseHandler(c *components.HTTPComponents) {
	courseRequest := httpmodels.RequestContent{}
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
//	@Success	200		{object}	httpmodels.ReviewResponse	"OK"
//	@Failure	409		"Conflict"
//	@Response	default	{object}	components.Response				"Standard error example object"
//	@Param		id		path		string							true	"ID of course"
//	@Param		request	body		httpmodels.ReviewRequestContent	true	"Request payload for creating a review"
//	@Router		/courses/{id}/review [post]
//	@Security	Bearer
//	@Security	Language
func CreateCourseReview(c *components.HTTPComponents) {
	courseID := chi.URLParam(c.HttpRequest, "id")
	if !govalidator.IsUUID(courseID) {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrInvalidUUID)
		return
	}

	user, ok := c.HttpRequest.Context().Value(middlewares.UserKey).(*models.User)
	if !ok {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}

	var requestContent httpmodels.ReviewRequestContent
	err := components.ValidateRequest(c, &requestContent)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	review := requestContent.ToEntityReview()
	review.UserID = user.ID
	review.CourseID = uuid.MustParse(courseID)

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

// AddAnswer
//
//	@Summary	Correct Answer
//
//	@Tags		Course
//	@success	204		"No content"
//	@Failure	409		"Conflict"
//	@Response	default	{object}	components.Response			"Standard error example object"
//	@Param		request	body		httpmodels.AddAnswerRequest	true	"Request payload for creating a review"
//	@Router		/courses/{id}/questions [post]
//	@Security	Bearer
//	@Security	Language
func AddAnswer(c *components.HTTPComponents) {

	courseID := chi.URLParam(c.HttpRequest, "id")

	if !govalidator.IsUUID(courseID) {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrInvalidUUID)
		return
	}

	currentUser, ok := c.HttpRequest.Context().Value(middlewares.UserKey).(*models.User)
	if !ok {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}

	var addAnswerRequest httpmodels.AddAnswerRequest
	err := components.ValidateRequest(c, &addAnswerRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	err = c.Components.Resources.Answers.SaveUserAnswer(&models.UserAnswer{
		QuestionID: addAnswerRequest.QuestionID,
		AnswerID:   addAnswerRequest.AnswerID,
		UserID:     currentUser.ID,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			components.HttpErrorResponse(c, http.StatusConflict, errutil.ErrUserAlreadyAnswerdQuestion)
			return
		} else {
			components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
			return
		}
	}

	c.Components.Resources.Courses.UpdateEnrollmentProgress(uuid.MustParse(courseID), currentUser.ID)

	components.HttpResponse(c, http.StatusNoContent)
}

// GetEnrollmentInfo
//
//	@Summary	Get Enrollment info
//
//	@Tags		Course
//	@success	200		"No content"
//	@Failure	400		"Bad Request"
//	@Failure	404		"Course not found"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		id		path		string				true	"ID of course"
//	@Router		/courses/{id}/enrollment [get]
//	@Security	Bearer
//	@Security	Language
func GetEnrollmentInfo(c *components.HTTPComponents) {

	courseID := chi.URLParam(c.HttpRequest, "id")

	if !govalidator.IsUUID(courseID) {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrInvalidUUID)
		return
	}

	currentUser, ok := c.HttpRequest.Context().Value(middlewares.UserKey).(*models.User)

	if !ok {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}

	enrollment, err := c.Components.Resources.Courses.GetEnrollmentProgress(uuid.MustParse(courseID), currentUser.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			components.HttpErrorResponse(c, http.StatusConflict, errutil.ErrCourseResourceNotFound)
			return
		} else {
			components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
			return
		}
	}

	components.HttpResponseWithPayload(c, ToEnrollmentResponse(enrollment), http.StatusNoContent)
}

// GetQuestionsByCourseID
//
//	@Summary	Get the questions by the course ID
//
//	@Tags		Course
//	@success	200		"No content"
//	@Failure	400		"Bad Request"
//	@Failure	404		"Course not found"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		id		path		string				true	"ID of course"
//	@Router		/courses/{id}/questions [get]
//	@Security	Bearer
//	@Security	Language
func GetQuestionsByCourseID(c *components.HTTPComponents) {

	courseID := chi.URLParam(c.HttpRequest, "id")

	if !govalidator.IsUUID(courseID) {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrInvalidUUID)
		return
	}

	questions, err := c.Components.Resources.Courses.GetQuestionsByCourseID(uuid.MustParse(courseID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			components.HttpErrorResponse(c, http.StatusConflict, errutil.ErrCourseResourceNotFound)
			return
		} else {
			components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
			return
		}
	}

	components.HttpResponseWithPayload(c, ToQuestionsListResponse(questions), http.StatusOK)
}

// GetReviewsByCourseID
//
//	@Summary	Get the reviews by the course ID
//
//	@Tags		Course
//	@success	200		"No content"
//	@Failure	400		"Bad Request"
//	@Failure	404		"Course not found"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		id		path		string				true	"ID of course"
//	@Router		/courses/{id}/reviews [get]
//	@Security	Bearer
//	@Security	Language
func GetReviewsByCourseID(c *components.HTTPComponents) {

	courseID := chi.URLParam(c.HttpRequest, "id")

	if !govalidator.IsUUID(courseID) {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrInvalidUUID)
		return
	}

	reviews, err := c.Components.Resources.Courses.GetReviewsByCourseID(uuid.MustParse(courseID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			components.HttpErrorResponse(c, http.StatusConflict, errutil.ErrCourseResourceNotFound)
			return
		} else {
			components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
			return
		}
	}

	components.HttpResponseWithPayload(c, ToReviewsListResponse(reviews), http.StatusOK)
}

// ListCategoriesHandler
//
//	@Summary	List categories with paginated response
//
//	@Tags		Course
//	@success	200		{array}	pagination.PaginationData{data=httpmodels.CategoryResponse}	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		page	query		int					false	"Page number"
//	@Param		limit	query		int					false	"Limit of elements per page"
//	@Router		/courses/categories [get]
//	@Security	Bearer
//	@Security	Language
func ListCategoriesHandler(c *components.HTTPComponents) {
	paginationData, err := pagination.GetPaginationData(c.HttpRequest.URL.Query())

	if errors.Is(err, errutil.ErrInvalidPageParam) {
		components.HttpErrorResponse(c, http.StatusNotFound, err)
		return
	} else if errors.Is(err, errutil.ErrInvalidLimitParam) {
		components.HttpErrorResponse(c, http.StatusNotFound, err)
		return
	}

	categories, count := c.Components.Resources.Categories.ListWithPagination(paginationData.Offset, paginationData.Limit)

	response := paginationData.ToResponse(
		ToCategoryListResponse(categories), int(count),
	)

	components.HttpResponseWithPayload(c, response, http.StatusOK)
}

// CreateCourseCategory
//
//	@Summary	Create course category
//
//	@Tags		Course
//	@Success	201		{object}	httpmodels.CategoryResponse	"OK"
//	@Failure	409		"Conflict"
//	@Response	default	{object}	components.Response			"Standard error example object"
//	@Param		request	body		httpmodels.CategoryRequest	true	"Request payload for creating a course category"
//	@Router		/courses/categories [post]
//	@Security	Bearer
//	@Security	Language
func CreateCourseCategory(c *components.HTTPComponents) {
	categoryRequest := httpmodels.CategoryRequest{}

	err := components.ValidateRequest(c, &categoryRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	category := categoryRequest.ToEntity()

	err = c.Components.Resources.Categories.Create(category)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			components.HttpErrorResponse(c, http.StatusConflict, errutil.ErrCategoryAlreadyExists)
			return
		} else {
			components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
			return
		}
	}

	components.HttpResponseWithPayload(c, ToCategoryResponse(*category), http.StatusCreated)
}

// UpdateCategoryHandler
//
//	@Summary	Update category by ID
//
//	@Tags		Course
//	@success	200		{object}	httpmodels.CategoryResponse	"OK"
//	@Failure	400		"Bad Request"
//	@Failure	404		"Category not found"
//	@Response	default	{object}	components.Response			"Standard error example object"
//	@Param		request	body		httpmodels.CategoryRequest	true	"Request payload for updating category information"
//	@Param		id		path		string						true	"ID of category to be updated"
//	@Router		/courses/categories/{id} [put]
//	@Security	Bearer
//	@Security	Language
func UpdateCategoryHandler(c *components.HTTPComponents) {
	categoryRequest := httpmodels.CategoryRequest{}
	err := components.ValidateRequest(c, &categoryRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	id := chi.URLParam(c.HttpRequest, "id")
	if !govalidator.IsUUID(id) {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrInvalidUUID)
		return
	}

	category := categoryRequest.ToEntity()
	category.ID = uuid.MustParse(id)

	rowsAffected, err := c.Components.Resources.Categories.Update(category)

	if err != nil {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}
	if rowsAffected == 0 {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrCourseResourceNotFound)
		return
	}

	components.HttpResponseWithPayload(c, ToCategoryResponse(*category), http.StatusOK)
}

// DeleteCategoryHandler
//
//	@Summary	Delete a category by ID
//
//	@Tags		Course
//	@success	204		"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		id		path		string				true	"ID of the category to be deleted"
//	@Router		/courses/categories/{id} [delete]
//	@Security	Bearer
//	@Security	Language
func DeleteCategoryHandler(c *components.HTTPComponents) {
	id := chi.URLParam(c.HttpRequest, "id")
	if !govalidator.IsUUID(id) {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrInvalidUUID)
		return
	}

	err := c.Components.Resources.Categories.Delete(uuid.MustParse(id))
	if err != nil {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}

	components.HttpResponse(c, http.StatusNoContent)
}

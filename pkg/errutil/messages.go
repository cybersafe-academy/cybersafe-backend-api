package errutil

import (
	"errors"
)

var (
	//General
	ErrUnexpectedError      = errors.New("ErrUnexpectedError")
	ErrInvalidUUID          = errors.New("ErrInvalidUUID")
	ErrFutureDateNotAllowed = errors.New("ErrFutureDateNotAllowed")

	//Pagination
	ErrInvalidPageParam  = errors.New("ErrInvalidPageParam")
	ErrInvalidLimitParam = errors.New("ErrInvalidLimitParam")

	//Authentication
	ErrLoginOrPasswordIncorrect = errors.New("ErrLoginOrPasswordIncorrect")

	//JWT
	ErrInvalidJWT             = errors.New("ErrInvalidJWT")
	ErrCredentialsMissing     = errors.New("ErrCredentialsMissing")
	ErrInvalidSignature       = errors.New("ErrInvalidSignature")
	ErrTokenExpiredOrPending  = errors.New("ErrTokenExpiredOrPending")
	ErrInvalidClaims          = errors.New("ErrInvalidClaims")
	ErrInsufficientPermission = errors.New("ErrInsufficientPermission")

	//User
	ErrUserResourceNotFound   = errors.New("ErrUserResourceNotFound")
	ErrInvalidUserRole        = errors.New("ErrInvalidUserRole")
	ErrCPFAlreadyInUse        = errors.New("ErrCPFAlreadyInUse")
	ErrEmailAlreadyInUse      = errors.New("ErrEmailAlreadyInUse")
	ErrCPFOrEmailAlreadyInUse = errors.New("ErrCPFOrEmailAlreadyInUse")

	//Course
	ErrInvalidCourseLevel     = errors.New("ErrInvalidCourseLevel")
	ErrCourseResourceNotFound = errors.New("ErrCourseResourceNotFound")

	//Companies
	ErrCNPJorEmailorBusinessNameAlreadyInUse = errors.New("ErrCNPJorEmailorBusinessNameAlreadyInUse")
	ErrCompanyResourceNotFound               = errors.New("ErrCompanyResourceNotFound")

	//Review
	ErrReviewAlreadyExists = errors.New("ErrReviewAlreadyExists")

	//Course
	ErrInvalidContentType = errors.New("ErrInvalidContentType")

	//Answer
	ErrUserAlreadyAnswerdQuestion    = errors.New("ErrUserAlreadyAnswerdQuestion")
	ErrCourseHasNoQuestionsAvailable = errors.New("ErrCourseHasNoQuestionsAvailable")

	//Personality Test
	ErrInvalidMBTIType = errors.New("ErrInvalidMBTIType")

	//Categories
	ErrCategoryAlreadyExists = errors.New("ErrCategoryAlreadyExists")
)

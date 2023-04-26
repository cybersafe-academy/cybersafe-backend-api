package errutil

import "errors"

var (
	//General
	ErrUnexpectedError = errors.New("unexpected error")
	ErrInvalidUUID     = errors.New("invalid uuid")

	//Pagination
	ErrInvalidPageParam  = errors.New("invalid page param")
	ErrInvalidLimitParam = errors.New("invalid limit param")

	//JWT
	ErrInvalidJWT             = errors.New("invalid JWT token")
	ErrCredentialsMissing     = errors.New("credentials missing")
	ErrInvalidSignature       = errors.New("invalid signature")
	ErrTokenExpiredOrPending  = errors.New("token expired or pending")
	ErrInvalidClaims          = errors.New("invalid claims")
	ErrInsufficientPermission = errors.New("insufficient permission for given resource")

	//User
	ErrUserResourceNotFound = errors.New("user not found with given identifier")
	ErrInvalidUserRole      = errors.New("invalid user role")

	//Course
	ErrInvalidCourseLevel = errors.New("invalid course level")

	//Course
	ErrInvalidContentType = errors.New("invalid content type")
)

package errutil

import "errors"

var (
	//General
	ErrUnexpectedError      = errors.New("unexpected error")
	ErrInvalidUUID          = errors.New("invalid uuid")
	ErrFutureDateNotAllowed = errors.New("future date not allowed")

	//Pagination
	ErrInvalidPageParam  = errors.New("invalid page param")
	ErrInvalidLimitParam = errors.New("invalid limit param")

	//Authentication
	ErrLoginOrPasswordIncorrect = errors.New("login or password incorrect")

	//JWT
	ErrInvalidJWT             = errors.New("invalid JWT token")
	ErrCredentialsMissing     = errors.New("credentials missing")
	ErrInvalidSignature       = errors.New("invalid signature")
	ErrTokenExpiredOrPending  = errors.New("token expired or pending")
	ErrInvalidClaims          = errors.New("invalid claims")
	ErrInsufficientPermission = errors.New("insufficient permission for given resource")

	//User
	ErrUserResourceNotFound   = errors.New("user not found with given identifier")
	ErrInvalidUserRole        = errors.New("invalid user role")
	ErrCPFAlreadyInUse        = errors.New("cpf already in use")
	ErrEmailAlreadyInUse      = errors.New("email already in use")
	ErrCPFOrEmailAlreadyInUse = errors.New("cpf or email already in use")

	//Course
	ErrInvalidCourseLevel     = errors.New("invalid course level")
	ErrCourseResourceNotFound = errors.New("course not found with given identifier")

	//Companies
	ErrCNPJorEmailorBusinessNameAlreadyInUse = errors.New("cnpj or email or business name already in use")
	ErrCompanyResourceNotFound               = errors.New("company not found with given identifier")

	//Course
	ErrInvalidContentType = errors.New("invalid content type")
)

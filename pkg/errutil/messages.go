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
	ErrInvalidJWT            = errors.New("invalid JWT token")
	ErrCredentialsMissing    = errors.New("credentials missing")
	ErrInvalidSignature      = errors.New("invalid signature")
	ErrTokenExpiredOrPending = errors.New("token expired or pending")
	ErrInvalidClaims         = errors.New("invalid claims")

	//User
	ErrUserResourceNotFound = errors.New("user not found with given identifier")
)

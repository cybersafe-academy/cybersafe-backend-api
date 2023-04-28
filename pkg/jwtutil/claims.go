package jwtutil

import (
	"cybersafe-backend-api/pkg/errutil"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	UserID string `json:"userID"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func (c CustomClaims) Validate() error {

	_, err := uuid.Parse(c.ID)

	if err != nil {
		return errutil.ErrInvalidUUID
	}

	// if !govalidator.IsIn(c.Role, models.ValidUserRoles...) {
	// 	return errutil.ErrInvalidUserRole
	// }

	return nil
}

package services

import (
	"cybersafe-backend-api/internal/services/courses"
	"cybersafe-backend-api/internal/services/users"
)

type Resources struct {
	Users   users.UsersManager
	Courses courses.CoursesManager
}

package courses

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
)

type CoursesManagerMock struct {
	ListWithPaginationMock       func(int, int) ([]models.Course, int)
	GetByIDMock                  func(uuid.UUID) (models.Course, error)
	CreateMock                   func(*models.Course) error
	DeleteMock                   func(uuid.UUID) error
	UpdateMock                   func(*models.Course) (int, error)
	IsRightAnswerMock            func(*models.Answer) bool
	UpdateEnrollmentProgressMock func(uuid.UUID, uuid.UUID)
	GetEnrollmentProgressMock    func(uuid.UUID, uuid.UUID) (models.Enrollment, error)
}

func (cm *CoursesManagerMock) ListWithPagination(limit, offset int) ([]models.Course, int) {
	return cm.ListWithPaginationMock(limit, offset)
}

func (cm *CoursesManagerMock) GetByID(id uuid.UUID) (models.Course, error) {
	return cm.GetByIDMock(id)
}

func (cm *CoursesManagerMock) Create(course *models.Course) error {
	return cm.CreateMock(course)
}

func (cm *CoursesManagerMock) Delete(id uuid.UUID) error {
	return cm.DeleteMock(id)
}

func (cm *CoursesManagerMock) Update(course *models.Course) (int, error) {
	return cm.UpdateMock(course)
}

func (cm *CoursesManagerMock) IsRightAnswer(answer *models.Answer) bool {
	return cm.IsRightAnswerMock(answer)
}

func (cm *CoursesManagerMock) UpdateEnrollmentProgress(courseID, userID uuid.UUID) {
	cm.UpdateEnrollmentProgressMock(courseID, userID)
}

func (cm *CoursesManagerMock) GetEnrollmentProgress(courseID, userID uuid.UUID) (models.Enrollment, error) {
	return cm.GetEnrollmentProgressMock(courseID, userID)
}

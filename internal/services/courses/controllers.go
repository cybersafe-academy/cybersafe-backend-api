package courses

import (
	"cybersafe-backend-api/internal/api/handlers/courses/httpmodels"
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CoursesManagerDB struct {
	DBConnection *gorm.DB
}

func (cm *CoursesManagerDB) ListWithPagination(offset, limit int) ([]models.CourseExtraFields, int) {
	var courses []models.CourseExtraFields
	var count int64

	cm.DBConnection.
		Table("courses").
		Preload("Contents").
		Preload("Reviews").
		Preload("Categories").
		Joins("LEFT JOIN reviews ON reviews.course_id = courses.id").
		Select("courses.*, avg(reviews.rating) as avg_rating").
		Where("courses.deleted_at IS NULL").
		Group("courses.id").
		Offset(offset).
		Limit(limit).
		Find(&courses).
		Count(&count)

	return courses, int(count)
}

func (cm *CoursesManagerDB) ListByCategory() *httpmodels.CourseByCategoryResponse {
	var results []httpmodels.RawCoursesByCategory

	cm.DBConnection.Raw(`
		SELECT ct.name AS category_name,
				AVG(r.rating) AS avg_rating,
				c.id AS course_id,
				c.title AS course_title,
				c.description AS course_description,
				c.content_in_hours AS course_content_in_hours,
				c.thumbnail_url AS course_thumbnail_url,
				c.level AS course_level
		FROM categories ct
		LEFT JOIN courses c ON c.category_id = ct.id
		LEFT JOIN reviews r ON r.course_id = c.id
	GROUP BY ct.name, c.id, c.title, c.description, c.content_in_hours, c.thumbnail_url, c.level;
	`).Scan(&results)

	response := GroupCoursesByCategory(results)

	return &response
}

func (cm *CoursesManagerDB) GetByID(id uuid.UUID) (models.Course, error) {
	var course models.Course

	result := cm.DBConnection.First(&course, id)

	return course, result.Error
}

func (cm *CoursesManagerDB) Create(course *models.Course) error {
	result := cm.DBConnection.Create(course)
	return result.Error
}

func (cm *CoursesManagerDB) Delete(id uuid.UUID) error {
	result := cm.DBConnection.Delete(&models.Course{}, id)
	return result.Error
}

func (cm *CoursesManagerDB) Update(course *models.Course) (int, error) {
	result := cm.DBConnection.Model(course).Clauses(clause.Returning{}).Updates(course)
	return int(result.RowsAffected), result.Error
}

func (cm *CoursesManagerDB) IsRightAnswer(answer *models.Answer) bool {
	result := cm.DBConnection.
		First(answer)
	return result.Error == nil
}

func (cm *CoursesManagerDB) UpdateEnrollmentProgress(courseID, userID uuid.UUID) {
	var questionsIDs []int64
	var userAnswersCount int64

	cm.DBConnection.Model(&models.Question{}).
		Where("course_id = ?", courseID).
		Pluck("id", &questionsIDs)

	if len(questionsIDs) <= 0 {
		return
	}

	cm.DBConnection.Model(&models.UserAnswer{}).
		Where("question_id IN(?)", questionsIDs).
		Count(&userAnswersCount)

	progress_percentage := float64((int(userAnswersCount) / len(questionsIDs)) * 100)

	cm.DBConnection.Model(&models.Enrollment{}).
		Where("course_id = ?", courseID).
		Where("user_id = ?", userID).
		Update("progress", progress_percentage)
}

func (cm *CoursesManagerDB) GetEnrollmentProgress(courseID, userID uuid.UUID) (models.Enrollment, error) {
	var enrollment models.Enrollment

	result := cm.DBConnection.
		Where("course_id = ?", courseID).
		Where("user_id = ?", userID).
		First(&enrollment)

	return enrollment, result.Error
}

func (cm *CoursesManagerDB) GetQuestionsByCourseID(courseID uuid.UUID) ([]models.Question, error) {
	var questions []models.Question

	result := cm.DBConnection.
		Preload(clause.Associations).
		Where("course_id = ?", courseID).
		Find(&questions)

	return questions, result.Error
}

func (cm *CoursesManagerDB) GetReviewsByCourseID(courseID uuid.UUID) ([]models.Review, error) {
	var reviews []models.Review

	result := cm.DBConnection.
		Preload(clause.Associations).
		Where("course_id = ?", courseID).
		Find(&reviews)

	return reviews, result.Error
}

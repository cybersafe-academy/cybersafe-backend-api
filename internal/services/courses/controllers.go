package courses

import (
	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/pkg/errutil"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CoursesManagerDB struct {
	DBConnection *gorm.DB
}

func (cm *CoursesManagerDB) ListWithPagination(offset, limit int) ([]models.CourseExtraFields, int) {
	var courses []models.CourseExtraFields
	cm.DBConnection.
		Table("courses").
		Preload("Reviews").
		Preload("Category").
		Preload("Questions").
		Preload("Questions.Answers").
		Joins("LEFT JOIN reviews ON reviews.course_id = courses.id").
		Select("courses.*, avg(reviews.rating) as avg_rating").
		Where("courses.deleted_at IS NULL").
		Group("courses.id").
		Offset(offset).
		Limit(limit).
		Find(&courses)

	var count int64
	cm.DBConnection.Model(&models.Course{}).Count(&count)

	return courses, int(count)
}

func (cm *CoursesManagerDB) ListCoursesWithRecommendation(userID, companyID uuid.UUID, userMBTIType string) ([]models.CourseExtraFields, error) {
	var response, subscribedCourses, recommendedCourses []models.CourseExtraFields
	var recommendedCategories []uuid.UUID

	// Base query for courses
	baseCourseQuery := func(db *gorm.DB) *gorm.DB {
		return db.Table("courses").
			Preload("Reviews").
			Preload("Category").
			Joins("LEFT JOIN reviews ON reviews.course_id = courses.id").
			Select("courses.*, avg(reviews.rating) as avg_rating").
			Where("courses.deleted_at IS NULL").
			Group("courses.id")
	}

	// Fetch default courses
	if err := baseCourseQuery(cm.DBConnection).Find(&response).Error; err != nil {
		return nil, err
	}

	// Fetch subscribed courses
	if err := baseCourseQuery(cm.DBConnection).
		Joins("JOIN enrollments ON enrollments.course_id = courses.id").
		Where("enrollments.user_id = ?", userID).
		Find(&subscribedCourses).Error; err != nil {
		log.Error().Err(err).Msg("Error while fetching subscribed courses")
	} else {
		for i := range subscribedCourses {
			subscribedCourses[i].Category.Name = "Subscribed"
		}
		response = append(response, subscribedCourses...)
	}

	// Fetching company-specific or default recommendations
	companyRecommendationsTX := cm.DBConnection.
		Model(&models.CompanyContentRecommendation{}).
		Where("company_id = ? AND mbti_type = ?", companyID, userMBTIType).
		Pluck("category_id", &recommendedCategories)

	if len(recommendedCategories) == 0 && companyRecommendationsTX.Error == nil {
		cm.DBConnection.
			Model(&models.DefaultContentRecommendation{}).
			Where("mbti_type = ?", userMBTIType).
			Pluck("category_id", &recommendedCategories)
	}

	if len(recommendedCategories) > 0 {
		if err := baseCourseQuery(cm.DBConnection).
			Where("courses.category_id IN (?)", recommendedCategories).
			Find(&recommendedCourses).Error; err != nil {
			log.Error().Err(err).Msg("Error while fetching recommended courses")
		} else {
			for i := range recommendedCourses {
				recommendedCourses[i].Category.Name = "Made for you"
			}
			response = append(response, recommendedCourses...)
		}
	}

	return response, nil
}

func (cm *CoursesManagerDB) GetEnrolledCourses(userID uuid.UUID) []models.Course {
	var courses []models.Course

	cm.DBConnection.
		Preload(clause.Associations).
		Joins("JOIN enrollments ON enrollments.course_id = courses.id").
		Where("enrollments.user_id = ?", userID).
		Find(&courses)

	return courses
}

func (cm *CoursesManagerDB) GetByID(id uuid.UUID) (models.Course, error) {
	var course models.Course

	result := cm.DBConnection.
		Table("courses").
		Preload("Reviews").
		Preload("Category").
		Preload("Questions").
		Preload("Questions.Answers").
		Joins("LEFT JOIN reviews ON reviews.course_id = courses.id").
		Select("courses.*, avg(reviews.rating) as avg_rating").
		Where("courses.deleted_at IS NULL").
		Group("courses.id").
		First(&course, id)

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
	questions := course.Questions

	err := cm.DBConnection.Model(course).Association("Questions").Clear()
	if err != nil {
		return 0, err
	}

	course.Questions = questions

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
		Update("quiz_progress", progress_percentage)
}

func (cm *CoursesManagerDB) UpdateEnrollmentStatus(courseID, userID uuid.UUID) (float64, error) {
	var questionsIDs []string
	var userAnswersCount int64

	cm.DBConnection.
		Model(&models.Question{}).
		Where("course_id = ?", courseID).
		Pluck("id", &questionsIDs)

	cm.DBConnection.
		Model(&models.UserAnswer{}).
		Preload("Answers").
		Joins("LEFT JOIN answers ON answers.id = user_answers.answer_id").
		Where("answers.question_id IN (?)", questionsIDs).
		Where("answers.is_correct = ?", true).
		Count(&userAnswersCount)

	if len(questionsIDs) <= 0 {
		return 0, errutil.ErrCourseHasNoQuestionsAvailable
	}

	hitsPercentage := float64((int(userAnswersCount) / len(questionsIDs)) * 100)

	var courseStatus string

	if hitsPercentage >= 70 {
		courseStatus = models.ApprovedStatus
	} else {
		courseStatus = models.FailedStatus
	}

	cm.DBConnection.Model(&models.Enrollment{}).
		Where("course_id = ?", courseID).
		Where("user_id = ?", userID).
		Updates(models.Enrollment{Status: courseStatus, QuizProgress: hitsPercentage})

	return hitsPercentage, nil
}

func (cm *CoursesManagerDB) Enroll(enrollment *models.Enrollment) error {
	result := cm.DBConnection.Create(enrollment)
	return result.Error
}

func (cm *CoursesManagerDB) AddComment(comment *models.Comment) error {
	result := cm.DBConnection.Create(comment)
	return result.Error
}

func (cm *CoursesManagerDB) ListCommentsByCourse(courseID uuid.UUID) []models.Comment {
	var companies []models.Comment

	cm.DBConnection.Preload(clause.Associations).
		Find(&companies)

	return companies
}

func (cm *CoursesManagerDB) AddLikeToComment(commentID, userID uuid.UUID) error {

	stmnt := cm.DBConnection.
		Where("comment_id = ?", commentID).
		Where("user_id = ?", userID)

	result := stmnt.
		First(&models.Comment{})

	if result.Error == nil {
		// If the comment was found, remove it
		result := stmnt.
			Delete(&models.Like{})

		return result.Error
	}

	result = cm.DBConnection.Create(&models.Like{
		CommentID: commentID,
		UserID:    userID,
	})

	return result.Error
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

func (cm *CoursesManagerDB) ExistsEnrollmentByUserIDAndCourseID(userID, courseID uuid.UUID) bool {
	result := cm.DBConnection.
		Where("user_id = ?", userID).
		Where("course_id = ?", courseID).
		First(&models.Enrollment{})

	return result.Error == nil
}

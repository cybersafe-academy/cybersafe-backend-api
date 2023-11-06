package analytics

import (
	"cybersafe-backend-api/internal/models"

	"gorm.io/gorm"
)

type AnalyticsManagerDB struct {
	DBConnection *gorm.DB
}

type MBTICount struct {
	MBTIType string
	Count    int
}

type AnalyticsData struct {
	NumberOfUsers     int         `json:"numberOfUsers"`
	CourseCompletion  float64     `json:"courseCompletion"`
	AccuracyInQuizzes float64     `json:"accuracyInQuizzes"`
	MBTICount         []MBTICount `json:"mbti"`
}

func (am *AnalyticsManagerDB) GetData() (*AnalyticsData, error) {

	var mbtiCounts []MBTICount

	userModel := &models.User{}

	// Number of users
	var userCount int64
	result := am.DBConnection.Model(userModel).Count(&userCount)
	if result.Error != nil {
		return nil, result.Error
	}

	// Course completion
	var completedAndFailedCount int64
	if err := am.DBConnection.Model(&models.Enrollment{}).
		Where("status IN (?, ?)", models.ApprovedStatus, models.FailedStatus).
		Count(&completedAndFailedCount).
		Error; err != nil {
		return nil, err
	}

	var inProgressCount int64
	if err := am.DBConnection.Model(&models.Enrollment{}).
		Where("status = ?", models.InProgressStatus).
		Count(&inProgressCount).
		Error; err != nil {
		return nil, err
	}

	totalEnrollments := completedAndFailedCount + inProgressCount
	completionPercentage := float64(completedAndFailedCount) / float64(totalEnrollments) * 100

	// Accuracy in quizzes
	var userAnswersCount int64
	if err := am.DBConnection.Model(&models.UserAnswer{}).
		Count(&userAnswersCount).
		Error; err != nil {
		return nil, err
	}

	var userCorrectAnswersCount int64
	if err := am.DBConnection.
		Model(&models.UserAnswer{}).
		Preload("Answers").
		Joins("LEFT JOIN answers ON answers.id = user_answers.answer_id").
		Where("answers.is_correct = ?", true).
		Count(&userCorrectAnswersCount).
		Error; err != nil {
		return nil, err
	}

	userAnswersPercentage := float64(userAnswersCount) / float64(userCorrectAnswersCount) * 100

	// Count by MBTI
	result = am.DBConnection.Model(userModel).
		Select("mbti_type, COUNT(*) as count").
		Group("mbti_type").
		Find(&mbtiCounts)
	if result.Error != nil {
		return nil, result.Error
	}

	data := AnalyticsData{
		NumberOfUsers:     int(userCount),
		CourseCompletion:  completionPercentage,
		AccuracyInQuizzes: userAnswersPercentage,
		MBTICount:         mbtiCounts,
	}

	return &data, nil
}

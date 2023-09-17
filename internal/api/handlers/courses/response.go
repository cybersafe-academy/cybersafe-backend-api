package courses

import (
	"cybersafe-backend-api/internal/api/handlers/courses/httpmodels"
	"cybersafe-backend-api/internal/models"
)

func ToListResponse(courses []models.CourseExtraFields) []httpmodels.CourseResponse {

	var coursesResponse []httpmodels.CourseResponse

	for _, course := range courses {

		courseResponse := httpmodels.CourseResponse{
			ID:        course.ID,
			CreatedAt: course.CreatedAt,
			UpdatedAt: course.UpdatedAt,
			DeletedAt: course.DeletedAt,
			AvgRating: course.AvgRating,
			CourseFields: httpmodels.CourseFields{
				Title:          course.Title,
				Description:    course.Description,
				ContentInHours: course.ContentInHours,
				ThumbnailURL:   course.Description,
				Level:          course.Level,
			},
			Category: httpmodels.CourseCategoryResponse{
				ID: course.Category.ID,
				CategoryFields: httpmodels.CategoryFields{
					Name: course.Category.Name,
				},
			},
		}

		var contentsResponse []httpmodels.ContentResponse

		for _, content := range course.Contents {
			contentsResponse = append(contentsResponse, httpmodels.ContentResponse{
				ContentFields: httpmodels.ContentFields{
					Title: content.Title,
					URL:   content.URL,
				},
				ID: content.ID,
			})
		}

		for _, question := range course.Questions {
			questionModel := models.Question{
				Wording: question.Wording,
			}

			for _, answer := range question.Answers {
				questionModel.Answers = append(questionModel.Answers, models.Answer{
					Text:      answer.Text,
					IsCorrect: answer.IsCorrect,
				})
			}

			course.Questions = append(course.Questions, questionModel)
		}

		courseResponse.Contents = contentsResponse
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

func ToResponse(course models.Course) httpmodels.CourseResponse {

	courseResponse := httpmodels.CourseResponse{
		ID:        course.ID,
		CreatedAt: course.CreatedAt,
		UpdatedAt: course.UpdatedAt,
		DeletedAt: course.DeletedAt,
		CourseFields: httpmodels.CourseFields{
			Title:          course.Title,
			Description:    course.Description,
			ContentInHours: course.ContentInHours,
			ThumbnailURL:   course.Description,
			Level:          course.Level,
		},
		Category: httpmodels.CourseCategoryResponse{
			ID: course.Category.ID,
			CategoryFields: httpmodels.CategoryFields{
				Name: course.Category.Name,
			},
		},
	}

	var contentsResponse []httpmodels.ContentResponse

	for _, content := range course.Contents {
		contentsResponse = append(contentsResponse, httpmodels.ContentResponse{
			ID: content.ID,
			ContentFields: httpmodels.ContentFields{
				Title: content.Title,
				URL:   content.URL,
			},
		})
	}

	for _, question := range course.Questions {
		questionResponse := httpmodels.QuestionResponse{
			ID: question.ID,
			QuestionFields: httpmodels.QuestionFields{
				Wording: question.Wording,
			},
		}

		for _, answer := range question.Answers {
			questionResponse.Answers = append(questionResponse.Answers,
				httpmodels.AnswerResponse{
					ID: answer.ID,
					AnswerFields: httpmodels.AnswerFields{
						Text:      answer.Text,
						IsCorrect: answer.IsCorrect,
					},
				})
		}

		courseResponse.Questions = append(courseResponse.Questions, questionResponse)
	}

	courseResponse.Contents = contentsResponse

	return courseResponse
}

func ToReviewResponse(review models.Review) httpmodels.ReviewResponse {

	reviewResponse := httpmodels.ReviewResponse{
		ID:        review.ID,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
		DeletedAt: review.DeletedAt,
		ReviewFields: httpmodels.ReviewFields{
			Comment: review.Comment,
			Rating:  review.Rating,
		},
		CourseID: review.CourseID,
		UserID:   review.UserID,
	}

	return reviewResponse
}

func ToCategoryResponse(category models.Category) httpmodels.CategoryResponse {
	categoryResponse := httpmodels.CategoryResponse{
		CategoryFields: httpmodels.CategoryFields{
			Name: category.Name,
		},
		ID:        category.ID,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		DeletedAt: category.DeletedAt,
	}

	return categoryResponse
}

func ToEnrollmentRespose(enrollment models.Enrollment) httpmodels.EnrollmentResponse {

	enrollmentResponse := httpmodels.EnrollmentResponse{
		EnrollmentFields: httpmodels.EnrollmentFields{
			Status:       enrollment.Status,
			QuizProgress: enrollment.QuizProgress,
		},
	}

	return enrollmentResponse
}

func ToQuestionsListResponse(questions []models.Question) []httpmodels.QuestionResponse {
	var questionsResponse []httpmodels.QuestionResponse

	for _, question := range questions {

		questionResponse := httpmodels.QuestionResponse{
			QuestionFields: httpmodels.QuestionFields{
				Wording: question.Wording,
			},
			ID: question.ID,
		}

		var answerResponse []httpmodels.AnswerResponse

		for _, answer := range question.Answers {
			answerResponse = append(answerResponse, httpmodels.AnswerResponse{
				AnswerFields: httpmodels.AnswerFields{
					Text: answer.Text,
				},
				ID: answer.ID,
			})
		}

		questionResponse.Answers = answerResponse
		questionsResponse = append(questionsResponse, questionResponse)

	}

	return questionsResponse

}

func ToReviewsListResponse(reviews []models.Review) []httpmodels.ReviewResponse {
	var reviewsResponse []httpmodels.ReviewResponse

	for _, review := range reviews {

		reviewResponse := httpmodels.ReviewResponse{
			ReviewFields: httpmodels.ReviewFields{
				Comment: review.Comment,
				Rating:  review.Rating,
			},
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
			DeletedAt: review.DeletedAt,
			CourseID:  review.CourseID,
			UserID:    review.UserID,
			ID:        review.ID,
		}

		reviewsResponse = append(reviewsResponse, reviewResponse)

	}

	return reviewsResponse
}

func ToCategoryListResponse(categories []models.Category) []httpmodels.CategoryResponse {
	var categoriesResponse []httpmodels.CategoryResponse

	for _, category := range categories {
		categoryResponse := httpmodels.CategoryResponse{
			CategoryFields: httpmodels.CategoryFields{
				Name: category.Name,
			},
			ID:        category.ID,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
			DeletedAt: category.DeletedAt,
		}

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return categoriesResponse
}

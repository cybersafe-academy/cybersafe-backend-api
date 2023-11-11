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
				Title:           course.Title,
				Description:     course.Description,
				TitlePtBr:       course.TitlePtBr,
				DescriptionPtBr: course.DescriptionPtBr,
				ContentInHours:  course.ContentInHours,
				ThumbnailURL:    course.ThumbnailURL,
				Level:           course.Level,
				ContentURL:      course.ContentURL,
			},
			Category: httpmodels.CourseCategoryResponse{
				ID: course.Category.ID,
				CategoryFields: httpmodels.CategoryFields{
					Name: course.Category.Name,
				},
			},
		}

		for _, question := range course.Questions {
			questionModel := httpmodels.QuestionResponse{
				QuestionFields: httpmodels.QuestionFields{
					Wording:     question.Wording,
					WordingPtBr: question.WordingPtBr,
				},
				ID: question.CourseID,
			}

			for _, answer := range question.Answers {
				questionModel.Answers = append(questionModel.Answers, httpmodels.AnswerResponse{
					AnswerFields: httpmodels.AnswerFields{
						Text:     answer.Text,
						TextPtBr: answer.TextPtBr,
					},
					ID: answer.ID,
				})
			}

			courseResponse.Questions = append(courseResponse.Questions, questionModel)
		}

		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

func ToCourseListResponse(courses []models.Course) []httpmodels.CourseResponse {
	var coursesResponse []httpmodels.CourseResponse

	for _, course := range courses {

		courseResponse := httpmodels.CourseResponse{
			ID:        course.ID,
			CreatedAt: course.CreatedAt,
			UpdatedAt: course.UpdatedAt,
			DeletedAt: course.DeletedAt,
			CourseFields: httpmodels.CourseFields{
				Title:           course.Title,
				Description:     course.Description,
				TitlePtBr:       course.TitlePtBr,
				DescriptionPtBr: course.DescriptionPtBr,
				ContentInHours:  course.ContentInHours,
				ThumbnailURL:    course.ThumbnailURL,
				Level:           course.Level,
				ContentURL:      course.ContentURL,
			},
			Category: httpmodels.CourseCategoryResponse{
				ID: course.Category.ID,
				CategoryFields: httpmodels.CategoryFields{
					Name: course.Category.Name,
				},
			},
		}

		for _, question := range course.Questions {
			questionModel := httpmodels.QuestionResponse{
				QuestionFields: httpmodels.QuestionFields{
					Wording:     question.Wording,
					WordingPtBr: question.WordingPtBr,
				},
				ID: question.CourseID,
			}

			for _, answer := range question.Answers {
				questionModel.Answers = append(questionModel.Answers, httpmodels.AnswerResponse{
					AnswerFields: httpmodels.AnswerFields{
						Text:     answer.Text,
						TextPtBr: answer.TextPtBr,
					},
					ID: answer.ID,
				})
			}

			courseResponse.Questions = append(courseResponse.Questions, questionModel)
		}

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
			Title:           course.Title,
			Description:     course.Description,
			TitlePtBr:       course.TitlePtBr,
			DescriptionPtBr: course.DescriptionPtBr,
			ContentInHours:  course.ContentInHours,
			ThumbnailURL:    course.ThumbnailURL,
			Level:           course.Level,
			ContentURL:      course.ContentURL,
		},
		Category: httpmodels.CourseCategoryResponse{
			ID: course.CategoryID,
			CategoryFields: httpmodels.CategoryFields{
				Name: course.Category.Name,
			},
		},
	}

	for _, question := range course.Questions {
		questionResponse := httpmodels.QuestionResponse{
			ID: question.ID,
			QuestionFields: httpmodels.QuestionFields{
				Wording:     question.Wording,
				WordingPtBr: question.WordingPtBr,
			},
		}

		for _, answer := range question.Answers {
			questionResponse.Answers = append(questionResponse.Answers,
				httpmodels.AnswerResponse{
					ID: answer.ID,
					AnswerFields: httpmodels.AnswerFields{
						Text:     answer.Text,
						TextPtBr: answer.TextPtBr,
					},
				})
		}

		courseResponse.Questions = append(courseResponse.Questions, questionResponse)
	}

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
		User: httpmodels.UserReviewFields{
			ID:   review.User.ID,
			Name: review.User.Name,
		},
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
			Status: enrollment.Status,
		},
		HitsPercentage: enrollment.QuizProgress,
	}

	return enrollmentResponse
}

func ToQuestionsListResponse(questions []models.Question) []httpmodels.QuestionResponse {
	var questionsResponse []httpmodels.QuestionResponse

	for _, question := range questions {

		questionResponse := httpmodels.QuestionResponse{
			QuestionFields: httpmodels.QuestionFields{
				Wording:     question.Wording,
				WordingPtBr: question.WordingPtBr,
			},
			ID: question.ID,
		}

		var answerResponse []httpmodels.AnswerResponse

		for _, answer := range question.Answers {
			answerResponse = append(answerResponse, httpmodels.AnswerResponse{
				AnswerFields: httpmodels.AnswerFields{
					Text:     answer.Text,
					TextPtBr: answer.TextPtBr,
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
			User: httpmodels.UserReviewFields{
				ID:   review.User.ID,
				Name: review.User.Name,
			},
			ID: review.ID,
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

func ToCommentsListResponse(comments []models.Comment) []httpmodels.CommentResponse {

	var commentsResponse []httpmodels.CommentResponse

	for _, comment := range comments {
		commentsResponse = append(commentsResponse, httpmodels.CommentResponse{
			ID: comment.ID,
			CommentFields: httpmodels.CommentFields{
				Text: comment.Text,
			},
			UserID:     comment.UserID,
			CourseID:   comment.CourseID,
			LikesCount: len(comment.Likes),
		})
	}

	return commentsResponse
}

func ToCourseByCategoryResponse(courses []models.CourseExtraFields) map[string][]httpmodels.CourseContentResponse {
	courseMap := make(map[string][]httpmodels.CourseContentResponse)

	for _, course := range courses {
		courseResponse := httpmodels.CourseContentResponse{
			ID:             course.ID,
			Title:          course.Title,
			TitlePtBr:      course.TitlePtBr,
			ThumbnailURL:   course.ThumbnailURL,
			ContentURL:     course.ContentURL,
			AvgRating:      course.AvgRating,
			Description:    course.Description,
			Level:          course.Level,
			ContentInHours: course.ContentInHours,
		}

		categoryName := course.Category.Name
		courseMap[categoryName] = append(courseMap[categoryName], courseResponse)
	}

	return courseMap
}

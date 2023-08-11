package courses

import "cybersafe-backend-api/internal/models"

func ToListResponse(courses []models.CourseExtraFields) []CourseResponse {

	var coursesResponse []CourseResponse

	for _, course := range courses {

		courseResponse := CourseResponse{
			ID:        course.ID,
			CreatedAt: course.CreatedAt,
			UpdatedAt: course.UpdatedAt,
			DeletedAt: course.DeletedAt,
			AvgRating: course.AvgRating,
			CourseFields: CourseFields{
				Title:          course.Title,
				Description:    course.Description,
				ContentInHours: course.ContentInHours,
				ThumbnailURL:   course.Description,
				Level:          course.Level,
			},
		}

		var contentsResponse []ContentResponse

		for _, content := range course.Contents {
			contentsResponse = append(contentsResponse, ContentResponse{
				ContentFields: ContentFields{
					Title:       content.Title,
					ContentType: content.ContentType,
					URL:         content.URL,
				},
				ID: content.ID,
			})
		}

		courseResponse.Contents = contentsResponse
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

func ToResponse(course models.Course) CourseResponse {

	courseResponse := CourseResponse{
		ID:        course.ID,
		CreatedAt: course.CreatedAt,
		UpdatedAt: course.UpdatedAt,
		DeletedAt: course.DeletedAt,
		CourseFields: CourseFields{
			Title:          course.Title,
			Description:    course.Description,
			ContentInHours: course.ContentInHours,
			ThumbnailURL:   course.Description,
			Level:          course.Level,
		},
	}

	var contentsResponse []ContentResponse

	for _, content := range course.Contents {
		contentsResponse = append(contentsResponse, ContentResponse{
			ID: content.ID,
			ContentFields: ContentFields{
				Title:       content.Title,
				ContentType: content.ContentType,
				URL:         content.URL,
			},
		})
	}

	courseResponse.Contents = contentsResponse

	return courseResponse
}

func ToReviewResponse(review models.Review) ReviewResponse {

	reviewResponse := ReviewResponse{
		ID:        review.ID,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
		DeletedAt: review.DeletedAt,
		ReviewFields: ReviewFields{
			Comment:  review.Comment,
			Rating:   review.Rating,
			CourseID: review.CourseID,
		},
		UserID: review.UserID,
	}

	return reviewResponse
}

package course

import "cybersafe-backend-api/internal/models"

func ToListResponse(courses []models.Course) []CourseResponse {

	var coursesResponse []CourseResponse

	for _, course := range courses {

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
				ContentFields: ContentFields{
					ContentType: content.ContentType,
					YoutubeURL:  content.YoutubeURL,
					PDFURL:      content.PDFURL,
					ImageURL:    content.ImageURL,
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
				YoutubeURL:  content.YoutubeURL,
				PDFURL:      content.PDFURL,
				ImageURL:    content.ImageURL,
			},
		})
	}

	courseResponse.Contents = contentsResponse

	return courseResponse
}

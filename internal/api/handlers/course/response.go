package course

import "cybersafe-backend-api/internal/models"

func ToListResponse(courses []models.Course) []ResponseContent {

	var coursesResponse []ResponseContent

	for _, course := range courses {

		courseResponse := ResponseContent{
			ID:        course.ID,
			CreatedAt: course.CreatedAt,
			UpdatedAt: course.UpdatedAt,
			DeletedAt: course.DeletedAt,
			CourseFields: CourseFields{
				Name:           course.Description,
				Description:    course.Description,
				ContentInHours: course.ContentInHours,
				ThumbnailURL:   course.Description,
				Level:          course.Level,
			},
		}

		coursesResponse = append(coursesResponse, courseResponse)

		var contentsResponse []ContentFields

		for _, content := range course.Contents {
			contentsResponse = append(contentsResponse, ContentFields{
				ID:          content.ID,
				ContentType: content.ContentType,
				YoutubeURL:  content.YoutubeURL,
				PDFURL:      content.PDFURL,
				ImageURL:    content.ImageURL,
			})
		}

		courseResponse.CourseFields.Contents = contentsResponse

	}

	return coursesResponse
}

func ToResponse(course models.Course) ResponseContent {

	courseResponse := ResponseContent{
		ID:        course.ID,
		CreatedAt: course.CreatedAt,
		UpdatedAt: course.UpdatedAt,
		DeletedAt: course.DeletedAt,
		CourseFields: CourseFields{
			Name:           course.Description,
			Description:    course.Description,
			ContentInHours: course.ContentInHours,
			ThumbnailURL:   course.Description,
			Level:          course.Level,
		},
	}

	var contentsResponse []ContentFields

	for _, content := range course.Contents {
		contentsResponse = append(contentsResponse, ContentFields{
			ID:          content.ID,
			ContentType: content.ContentType,
			YoutubeURL:  content.YoutubeURL,
			PDFURL:      content.PDFURL,
			ImageURL:    content.ImageURL,
		})
	}

	courseResponse.CourseFields.Contents = contentsResponse

	return courseResponse
}

package courses

import "cybersafe-backend-api/internal/api/handlers/courses/httpmodels"

func GroupCoursesByCategory(rawCourses []httpmodels.RawCoursesByCategory) httpmodels.CourseByCategoryResponse {
	responseMap := make(httpmodels.CourseByCategoryResponse)

	for _, course := range rawCourses {
		categoryName := course.CategoryName

		courseData := map[string]any{
			"id":             course.CourseID,
			"title":          course.CourseTitle,
			"thumbnailURL":   course.CourseThumbnailURL,
			"avgRating":      course.AvgRating,
			"contentURL":     course.CourseContentURL,
			"description":    course.CourseDescription,
			"level":          course.CourseLevel,
			"contentInHours": course.CourseContentInHours,
		}

		responseMap[categoryName] = append(responseMap[categoryName], courseData)
	}

	return responseMap
}
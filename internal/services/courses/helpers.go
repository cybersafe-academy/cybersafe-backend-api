package courses

// func GroupCoursesByCategory(rawCourses []httpmodels.RawCoursesByCategory) httpmodels.CourseByCategoryResponse {
// 	responseMap := make(httpmodels.CourseByCategoryResponse)

// 	for _, course := range rawCourses {
// 		if course.CourseID == uuid.Nil {
// 			continue
// 		}

// 		categoryName := course.CategoryName

// 		courseData := map[string]any{
// 			"id":             course.CourseID,
// 			"title":          course.CourseTitle,
// 			"titlePtBr":      course.CourseTitlePtBr,
// 			"title_pt_br":    course.CourseTitlePtBr,
// 			"thumbnailURL":   course.CourseThumbnailURL,
// 			"avgRating":      course.AvgRating,
// 			"contentURL":     course.CourseContentURL,
// 			"description":    course.CourseDescription,
// 			"level":          course.CourseLevel,
// 			"contentInHours": course.CourseContentInHours,
// 		}

// 		responseMap[categoryName] = append(responseMap[categoryName], courseData)
// 	}

// 	return responseMap
// }

package analytics

type MBTICount struct {
	MBTIType string `json:"mbtiType"`
	Count    int    `json:"count"`
}

type AnalyticsDataResponse struct {
	NumberOfUsers     int         `json:"numberOfUsers"`
	CourseCompletion  float64     `json:"courseCompletion"`
	AccuracyInQuizzes float64     `json:"accuracyInQuizzes"`
	MBTICount         []MBTICount `json:"mbtiCount"`
}

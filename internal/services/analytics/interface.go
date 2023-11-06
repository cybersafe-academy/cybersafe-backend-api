package analytics

type AnalyticsManager interface {
	GetData() (*AnalyticsData, error)
}

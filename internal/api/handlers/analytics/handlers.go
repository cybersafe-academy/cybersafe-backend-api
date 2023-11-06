package analytics

import (
	"cybersafe-backend-api/internal/api/components"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// GetAnalyticsData is the HTTP handler for getting data for dashboard
//
//	@Summary		Data analytics
//	@Description	Gets all analytics data
//	@Tags			Analytics
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	AnalyticsDataResponse
//	@Failure		400	"Bad Request"
//	@Router			/analytics/data [get]
//	@Security		Bearer
//	@Security		Language
func GetAnalyticsData(c *components.HTTPComponents) {

	result, err := c.Components.Resources.Analytics.GetData()
	if err != nil {
		components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrUnexpectedError",
		}))
	}

	response := AnalyticsDataResponse{
		NumberOfUsers:     result.NumberOfUsers,
		CourseCompletion:  result.CourseCompletion,
		AccuracyInQuizzes: result.AccuracyInQuizzes,
	}

	var mbtiCountResponse []MBTICount
	for _, mbti := range result.MBTICount {
		mbtiItem := MBTICount{
			MBTIType: mbti.MBTIType,
			Count:    mbti.Count,
		}
		mbtiCountResponse = append(mbtiCountResponse, mbtiItem)
	}

	response.MBTICount = mbtiCountResponse

	components.HttpResponseWithPayload(c, response, http.StatusOK)

}

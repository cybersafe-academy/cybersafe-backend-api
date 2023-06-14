package pagination

import "math"

type PaginationData struct {
	Limit  int
	Page   int
	Offset int
}

type PaginatedResponse struct {
	Data        any `json:"data"`
	Total       int `json:"total"`
	Limit       int `json:"limit"`
	CurrentPage int `json:"current"`
	TotalPages  int `json:"totalPages"`
}

func (pd *PaginationData) ToResponse(data any, total int) *PaginatedResponse {
	return &PaginatedResponse{
		Data:        data,
		Total:       total,
		Limit:       pd.Limit,
		CurrentPage: pd.Page,
		TotalPages:  int(math.Ceil(float64(total) / float64(pd.Limit))),
	}
}

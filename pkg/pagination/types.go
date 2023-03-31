package pagination

import "math"

type PaginationData struct {
	Limit  int
	Page   int
	Offset int
}

type PaginteResponse struct {
	Content     any `json:"content"`
	Total       int `json:"total"`
	Limit       int `json:"limit"`
	CurrentPage int `json:"current"`
	TotalPages  int `json:"totalPages"`
}

func (pd *PaginationData) ToResponse(content any, total int) *PaginteResponse {
	return &PaginteResponse{
		Content:     content,
		Total:       total,
		Limit:       pd.Limit,
		CurrentPage: pd.Page,
		TotalPages:  int(math.Ceil(float64(total) / float64(pd.Limit))),
	}
}

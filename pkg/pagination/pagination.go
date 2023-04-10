package pagination

import (
	"cybersafe-backend-api/pkg/errutil"
	"net/url"
	"strconv"
)

func GetPaginationData(queryParams url.Values) (*PaginationData, error) {

	paginationData := PaginationData{
		Limit: 10,
		Page:  1,
	}

	paginationData.Offset = (paginationData.Page - 1) * paginationData.Limit

	if pageParam := queryParams.Get("page"); pageParam != "" {
		page, err := strconv.Atoi(pageParam)

		paginationData.Page = page

		if err != nil {
			return nil, errutil.ErrInvalidPageParam
		}
	}

	if limitParam := queryParams.Get("limit"); limitParam != "" {
		limit, err := strconv.Atoi(limitParam)

		paginationData.Limit = limit

		if err != nil {
			return nil, errutil.ErrInvalidLimitParam
		}

	}

	return &paginationData, nil

}

package pagination

import (
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/shared/constants"
	"math"
)

func Paginate[T any](items []T, page, pageSize int) entity.PaginatedResponse[T] {
	if page <= 0 {
		page = constants.DefaultPageNum
	}
	if pageSize <= 0 {
		pageSize = constants.DefaultPageSize
	}

	totalRecords := len(items)
	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	if startIndex > totalRecords {
		startIndex = totalRecords
	}
	if endIndex > totalRecords {
		endIndex = totalRecords
	}

	paginatedItems := items[startIndex:endIndex]

	return entity.PaginatedResponse[T]{
		Items: paginatedItems,
		Pagination: entity.Pagination{
			Page:        page,
			TotalPage:   totalPages,
			TotalRecord: totalRecords,
		},
	}
}

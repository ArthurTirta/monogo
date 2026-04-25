package queryhelper

import (
	entitybase "github.com/ArthurTirta/monogo/internal/entity/base"
	dtobase "github.com/ArthurTirta/monogo/pkg/dto/base"
)

func SerializeFilterPaginationDtoToEntity(
	pagination dtobase.BaseReqQueryPagination,
) entitybase.BasePaginationFilter {
	return entitybase.BasePaginationFilter{
		MinCreated:  pagination.CreatedAtGTE,
		MaxCreated:  pagination.CreatedAtLTE,
		MinUpdated:  pagination.UpdatedAtGTE,
		MaxUpdated:  pagination.UpdatedAtLTE,
		WithDeleted: pagination.IncludeDeleted,
		ShowCount:   pagination.ShowCount,
		Offset:      pagination.Offset,
		Limit:       pagination.Limit,
		OrderBy:     pagination.OrderBy,
	}
}

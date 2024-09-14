// Code generated by MVRP Codegen Util. DO NOT EDIT.

package dto

import (
    "mvrp/data/model/query"
    "mvrp/data/model/entity"
)

type SearchEntityDTO struct {
    query.IServerSideGetRowsRequest `json:"server_side_get_rows_request"`
    Keyword      string                             `json:"keyword"`
    EntityType string          `json:"entity_type"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetEntityDTO struct {
    entity.Entity                 `json:"entity"`
}

type CreateEntityDTO struct {
    entity.Entity                 `json:"entity"`
}

type UpdateEntityDTO struct {
    entity.Entity                 `json:"entity"`
}
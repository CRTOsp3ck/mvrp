// Code generated by MVRP Codegen Util. DO NOT EDIT.

package dto

import (
    "mvrp/data/model/entity"
)

type SearchEntityDTO struct {
    Keyword      string                             `json:"keyword"`
    Type string          `json:"type"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type CreateEntityDTO struct {
    entity.Entity                 `json:"entity"`
}

type UpdateEntityDTO struct {
    entity.Entity                 `json:"entity"`
}
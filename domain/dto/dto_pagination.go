package dto

type PaginationDTO struct {
	TotalItems   int    `json:"total_items"`
	ItemsPerPage int    `json:"items_per_page"`
	Page         int    `json:"page"`
	SortBy       string `json:"sort_by"`
	OrderBy      string `json:"order_by"`
}

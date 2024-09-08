// Code generated by MVRP Codegen Util. DO NOT EDIT.

package dto

import (
    "mvrp/data/model/inventory"
)

type SearchInventoryDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetInventoryDTO struct {
    inventory.Inventory                 `json:"inventory"`
}

type CreateInventoryDTO struct {
    inventory.Inventory                 `json:"inventory"`
}

type UpdateInventoryDTO struct {
    inventory.Inventory                 `json:"inventory"`
}

type SearchInventoryTransactionDTO struct {
    Keyword      string                             `json:"keyword"`
    InventoryId string          `json:"inventory_id"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetInventoryTransactionDTO struct {
    inventory.InventoryTransaction                 `json:"inventory_transaction"`
}

type CreateInventoryTransactionDTO struct {
    inventory.InventoryTransaction                 `json:"inventory_transaction"`
}

type UpdateInventoryTransactionDTO struct {
    inventory.InventoryTransaction                 `json:"inventory_transaction"`
}

type SearchGoodsIssueNoteDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetGoodsIssueNoteDTO struct {
    inventory.GoodsIssueNote                 `json:"goods_issue_note"`
    Items []GetGoodsIssueNoteItemDTO                     `json:"items"`
}

type CreateGoodsIssueNoteDTO struct {
    inventory.GoodsIssueNote                 `json:"goods_issue_note"`
    Items []CreateGoodsIssueNoteItemDTO                     `json:"items"`
}

type UpdateGoodsIssueNoteDTO struct {
    inventory.GoodsIssueNote                 `json:"goods_issue_note"`
    Items []UpdateGoodsIssueNoteItemDTO                     `json:"items"`
}

type SearchGoodsIssueNoteItemDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetGoodsIssueNoteItemDTO struct {
    inventory.GoodsIssueNoteItem                 `json:"goods_issue_note_item"`
}

type CreateGoodsIssueNoteItemDTO struct {
    inventory.GoodsIssueNoteItem                 `json:"goods_issue_note_item"`
}

type UpdateGoodsIssueNoteItemDTO struct {
    inventory.GoodsIssueNoteItem                 `json:"goods_issue_note_item"`
}

type SearchStockCountSheetDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetStockCountSheetDTO struct {
    inventory.StockCountSheet                 `json:"stock_count_sheet"`
}

type CreateStockCountSheetDTO struct {
    inventory.StockCountSheet                 `json:"stock_count_sheet"`
}

type UpdateStockCountSheetDTO struct {
    inventory.StockCountSheet                 `json:"stock_count_sheet"`
}

type SearchReturnMerchandiseAuthorizationDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetReturnMerchandiseAuthorizationDTO struct {
    inventory.ReturnMerchandiseAuthorization                 `json:"return_merchandise_authorization"`
    Items []CreateReturnMerchandiseAuthorizationItemDTO                     `json:"items"`
}

type CreateReturnMerchandiseAuthorizationDTO struct {
    inventory.ReturnMerchandiseAuthorization                 `json:"return_merchandise_authorization"`
    Items []CreateReturnMerchandiseAuthorizationItemDTO                     `json:"items"`
}

type UpdateReturnMerchandiseAuthorizationDTO struct {
    inventory.ReturnMerchandiseAuthorization                 `json:"return_merchandise_authorization"`
    Items []UpdateReturnMerchandiseAuthorizationItemDTO                     `json:"items"`
}

type SearchReturnMerchandiseAuthorizationItemDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetReturnMerchandiseAuthorizationItemDTO struct {
    inventory.ReturnMerchandiseAuthorizationItem                 `json:"return_merchandise_authorization_item"`
}

type CreateReturnMerchandiseAuthorizationItemDTO struct {
    inventory.ReturnMerchandiseAuthorizationItem                 `json:"return_merchandise_authorization_item"`
}

type UpdateReturnMerchandiseAuthorizationItemDTO struct {
    inventory.ReturnMerchandiseAuthorizationItem                 `json:"return_merchandise_authorization_item"`
}
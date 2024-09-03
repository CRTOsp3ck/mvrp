// Code generated by MVRP Codegen Util. DO NOT EDIT.

package dto

import (
    "mvrp/data/model/base"
    "mvrp/data/model/sale"
)

type SearchSalesOrderDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type CreateSalesOrderDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.SalesOrder                 `json:"sales_order"`
    Items []CreateSalesOrderItemDTO                     `json:"items"`
}

type UpdateSalesOrderDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.SalesOrder                 `json:"sales_order"`
    Items []UpdateSalesOrderItemDTO                     `json:"items"`
}

type SearchSalesOrderItemDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type CreateSalesOrderItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.SalesOrderItem                 `json:"sales_order_item"`
}

type UpdateSalesOrderItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.SalesOrderItem                 `json:"sales_order_item"`
}

type SearchDeliveryNoteDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type CreateDeliveryNoteDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.DeliveryNote                 `json:"delivery_note"`
    Items []CreateDeliveryNoteItemDTO                     `json:"items"`
}

type UpdateDeliveryNoteDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.DeliveryNote                 `json:"delivery_note"`
    Items []UpdateDeliveryNoteItemDTO                     `json:"items"`
}

type SearchDeliveryNoteItemDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type CreateDeliveryNoteItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.DeliveryNoteItem                 `json:"delivery_note_item"`
}

type UpdateDeliveryNoteItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.DeliveryNoteItem                 `json:"delivery_note_item"`
}

type SearchGoodsReturnNoteDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type CreateGoodsReturnNoteDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.GoodsReturnNote                 `json:"goods_return_note"`
    Items []CreateGoodsReturnNoteItemDTO                     `json:"items"`
}

type UpdateGoodsReturnNoteDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.GoodsReturnNote                 `json:"goods_return_note"`
    Items []UpdateGoodsReturnNoteItemDTO                     `json:"items"`
}

type SearchGoodsReturnNoteItemDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type CreateGoodsReturnNoteItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.GoodsReturnNoteItem                 `json:"goods_return_note_item"`
}

type UpdateGoodsReturnNoteItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.GoodsReturnNoteItem                 `json:"goods_return_note_item"`
}

type SearchOrderConfirmationDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type CreateOrderConfirmationDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.OrderConfirmation                 `json:"order_confirmation"`
    Items []CreateOrderConfirmationItemDTO                     `json:"items"`
}

type UpdateOrderConfirmationDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.OrderConfirmation                 `json:"order_confirmation"`
    Items []UpdateOrderConfirmationItemDTO                     `json:"items"`
}

type SearchOrderConfirmationItemDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type CreateOrderConfirmationItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.OrderConfirmationItem                 `json:"order_confirmation_item"`
}

type UpdateOrderConfirmationItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.OrderConfirmationItem                 `json:"order_confirmation_item"`
}

type SearchSalesQuotationDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type CreateSalesQuotationDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.SalesQuotation                 `json:"sales_quotation"`
    Items []CreateSalesQuotationItemDTO                     `json:"items"`
}

type UpdateSalesQuotationDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.SalesQuotation                 `json:"sales_quotation"`
    Items []UpdateSalesQuotationItemDTO                     `json:"items"`
}

type SearchSalesQuotationItemDTO struct {
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type CreateSalesQuotationItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.SalesQuotationItem                 `json:"sales_quotation_item"`
}

type UpdateSalesQuotationItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.SalesQuotationItem                 `json:"sales_quotation_item"`
}
// Code generated by MVRP Codegen Util. DO NOT EDIT.

package dto

import (
    "mvrp/data/model/query"
    "mvrp/data/model/base"
    "mvrp/data/model/sale"
)

type SearchSalesOrderDTO struct {
    query.IServerSideGetRowsRequest `json:"server_side_get_rows_request"`
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetSalesOrderDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.SalesOrder                 `json:"sales_order"`
    Items []CreateSalesOrderItemDTO                     `json:"items"`
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
    query.IServerSideGetRowsRequest `json:"server_side_get_rows_request"`
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetSalesOrderItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.SalesOrderItem                 `json:"sales_order_item"`
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
    query.IServerSideGetRowsRequest `json:"server_side_get_rows_request"`
    Keyword      string                             `json:"keyword"`
    ShippingStatus string          `json:"shipping_status"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetDeliveryNoteDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.DeliveryNote                 `json:"delivery_note"`
    Items []CreateDeliveryNoteItemDTO                     `json:"items"`
}

type CreateDeliveryNoteDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.DeliveryNote                 `json:"delivery_note"`
    Items []CreateDeliveryNoteItemDTO                     `json:"items"`
    ToCreateFromSalesOrder bool     `json:"to_create_from_sales_order"`
}

type UpdateDeliveryNoteDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.DeliveryNote                 `json:"delivery_note"`
    Items []UpdateDeliveryNoteItemDTO                     `json:"items"`
}

type SearchDeliveryNoteItemDTO struct {
    query.IServerSideGetRowsRequest `json:"server_side_get_rows_request"`
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetDeliveryNoteItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.DeliveryNoteItem                 `json:"delivery_note_item"`
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
    query.IServerSideGetRowsRequest `json:"server_side_get_rows_request"`
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetGoodsReturnNoteDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.GoodsReturnNote                 `json:"goods_return_note"`
    Items []CreateGoodsReturnNoteItemDTO                     `json:"items"`
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
    query.IServerSideGetRowsRequest `json:"server_side_get_rows_request"`
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetGoodsReturnNoteItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.GoodsReturnNoteItem                 `json:"goods_return_note_item"`
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
    query.IServerSideGetRowsRequest `json:"server_side_get_rows_request"`
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetOrderConfirmationDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.OrderConfirmation                 `json:"order_confirmation"`
    Items []CreateOrderConfirmationItemDTO                     `json:"items"`
}

type CreateOrderConfirmationDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.OrderConfirmation                 `json:"order_confirmation"`
    Items []CreateOrderConfirmationItemDTO                     `json:"items"`
    ToCreateFromSalesOrder bool     `json:"to_create_from_sales_order"`
}

type UpdateOrderConfirmationDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.OrderConfirmation                 `json:"order_confirmation"`
    Items []UpdateOrderConfirmationItemDTO                     `json:"items"`
}

type SearchOrderConfirmationItemDTO struct {
    query.IServerSideGetRowsRequest `json:"server_side_get_rows_request"`
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetOrderConfirmationItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.OrderConfirmationItem                 `json:"order_confirmation_item"`
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
    query.IServerSideGetRowsRequest `json:"server_side_get_rows_request"`
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetSalesQuotationDTO struct {
    base.BaseDocument                 `json:"base_document"`
    sale.SalesQuotation                 `json:"sales_quotation"`
    Items []CreateSalesQuotationItemDTO                     `json:"items"`
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
    query.IServerSideGetRowsRequest `json:"server_side_get_rows_request"`
    Keyword      string                             `json:"keyword"`
    ItemsPerPage int                                `json:"items_per_page"`
    Page         int                                `json:"page"`
    SortBy       string                             `json:"sort_by"`
    OrderBy      string                             `json:"order_by"`
}

type GetSalesQuotationItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.SalesQuotationItem                 `json:"sales_quotation_item"`
}

type CreateSalesQuotationItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.SalesQuotationItem                 `json:"sales_quotation_item"`
}

type UpdateSalesQuotationItemDTO struct {
    base.BaseDocumentItem                 `json:"base_document_item"`
    sale.SalesQuotationItem                 `json:"sales_quotation_item"`
}
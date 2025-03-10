// Code generated by MVRP Codegen Util. DO NOT EDIT.

package dto

type EnumsDTO struct {
	BaseEnums      BaseEnumDTO    `json:"base_enums"`
	ItemEnums      ItemEnumDTO    `json:"item_enums"`
	EntityEnums      EntityEnumDTO    `json:"entity_enums"`
	InventoryEnums      InventoryEnumDTO    `json:"inventory_enums"`
	PurchaseEnums      PurchaseEnumDTO    `json:"purchase_enums"`
	SaleEnums      SaleEnumDTO    `json:"sale_enums"`
}


type BaseEnumDTO struct {
	PaymentTerms     []Enum    `json:"payment_terms"`
	PaymentStatus     []Enum    `json:"payment_status"`
	ShippingTerms     []Enum    `json:"shipping_terms"`
	ShippingMethods     []Enum    `json:"shipping_methods"`
}

type ItemEnumDTO struct {
	ItemType     []Enum    `json:"item_type"`
	ItemStatus     []Enum    `json:"item_status"`
}

type EntityEnumDTO struct {
	EntityType     []Enum    `json:"entity_type"`
	EntityStatus     []Enum    `json:"entity_status"`
}

type InventoryEnumDTO struct {
	InventoryTransactionType     []Enum    `json:"inventory_transaction_type"`
}

type PurchaseEnumDTO struct {
	PurchaseOrderStatus     []Enum    `json:"purchase_order_status"`
}

type SaleEnumDTO struct {
	SalesOrderStatus     []Enum    `json:"sales_order_status"`
	SalesQuotationStatus     []Enum    `json:"sales_quotation_status"`
	SalesShippingStatus     []Enum    `json:"sales_shipping_status"`
}


type Enum struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

/*

--------- FOR REFERENCE OF THE ORIGINAL ENUM TYPES ---------

import (
	"mvrp/data/model/base"
	"mvrp/data/model/item"
	"mvrp/data/model/entity"
	"mvrp/data/model/inventory"
	"mvrp/data/model/purchase"
	"mvrp/data/model/sale"
)

type EnumsDTO struct {
	BaseEnums      BaseEnumDTO    `json:"base_enums"`
	ItemEnums      ItemEnumDTO    `json:"item_enums"`
	EntityEnums      EntityEnumDTO    `json:"entity_enums"`
	InventoryEnums      InventoryEnumDTO    `json:"inventory_enums"`
	PurchaseEnums      PurchaseEnumDTO    `json:"purchase_enums"`
	SaleEnums      SaleEnumDTO    `json:"sale_enums"`
}


type BaseEnumDTO struct {
	PaymentTerms     []base.PaymentTerms    `json:"payment_terms"`
	PaymentStatus     []base.PaymentStatus    `json:"payment_status"`
	ShippingTerms     []base.ShippingTerms    `json:"shipping_terms"`
	ShippingMethods     []base.ShippingMethods    `json:"shipping_methods"`
}

type ItemEnumDTO struct {
	ItemType     []item.ItemType    `json:"item_type"`
	ItemStatus     []item.ItemStatus    `json:"item_status"`
}

type EntityEnumDTO struct {
	EntityType     []entity.EntityType    `json:"entity_type"`
	EntityStatus     []entity.EntityStatus    `json:"entity_status"`
}

type InventoryEnumDTO struct {
	InventoryTransactionType     []inventory.InventoryTransactionType    `json:"inventory_transaction_type"`
}

type PurchaseEnumDTO struct {
	PurchaseOrderStatus     []purchase.PurchaseOrderStatus    `json:"purchase_order_status"`
}

type SaleEnumDTO struct {
	SalesOrderStatus     []sale.SalesOrderStatus    `json:"sales_order_status"`
	SalesQuotationStatus     []sale.SalesQuotationStatus    `json:"sales_quotation_status"`
	SalesShippingStatus     []sale.SalesShippingStatus    `json:"sales_shipping_status"`
}


*/

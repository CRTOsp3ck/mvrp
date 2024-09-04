package enum

import (
	"mvrp/data/model/base"
	"mvrp/data/model/entity"
	"mvrp/data/model/inventory"
	"mvrp/data/model/item"
	"mvrp/data/model/purchase"
	"mvrp/data/model/sale"
)

// -------------------------------------------------------
// BASE
// -------------------------------------------------------

func (r *EnumRepository) GetPaymentTermsEnums() []base.PaymentTerms {
	return base.AllPaymentTerms()
}

func (r *EnumRepository) GetPaymentStatusEnums() []base.PaymentStatus {
	return base.AllPaymentStatus()
}

// -------------------------------------------------------
// ITEM
// -------------------------------------------------------

func (r *EnumRepository) GetItemTypeEnums() []item.ItemType {
	return item.AllItemType()
}

// -------------------------------------------------------
// ENTITY
// -------------------------------------------------------

func (r *EnumRepository) GetEntityTypeEnums() []entity.EntityType {
	return entity.AllEntityType()
}

// -------------------------------------------------------
// INVENTORY
// -------------------------------------------------------

func (r *EnumRepository) GetInventoryTransactionTypeEnums() []inventory.InventoryTransactionType {
	return inventory.AllInventoryTransactionType()
}

// -------------------------------------------------------
// PURCHASE
// -------------------------------------------------------

func (r *EnumRepository) GetPurchaseOrderStatusEnums() []purchase.PurchaseOrderStatus {
	return purchase.AllPurchaseOrderStatus()
}

// -------------------------------------------------------
// SALE
// -------------------------------------------------------

func (r *EnumRepository) GetSalesOrderStatusEnums() []sale.SalesOrderStatus {
	return sale.AllSalesOrderStatus()
}

func (r *EnumRepository) GetSalesQuotationStatusEnums() []sale.SalesQuotationStatus {
	return sale.AllSalesQuotationStatus()
}

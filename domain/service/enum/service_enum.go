package enum

import (
	"context"
	"mvrp/domain/dto"
	"mvrp/util"
)

// LIST ENUM
type ListEnumRequest struct {
	Ctx context.Context
}

func (s *EnumService) NewListEnumRequest(ctx context.Context) *ListEnumRequest {
	return &ListEnumRequest{
		Ctx: ctx,
	}
}

type ListEnumResponse struct {
	Payload dto.EnumsDTO `json:"payload"`
}

func (s *EnumService) NewListEnumResponse(payload dto.EnumsDTO) *ListEnumResponse {
	return &ListEnumResponse{
		Payload: payload,
	}
}

func (s *EnumService) ListEnum(req *ListEnumRequest) (*ListEnumResponse, error) {
	// Base - Payment Terms
	var paymentTerms []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetPaymentTermsEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		paymentTerms = append(paymentTerms, enum)
	}

	// Base - Payment Status
	var paymentStatus []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetPaymentStatusEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		paymentStatus = append(paymentStatus, enum)
	}

	// Base - Shipping Terms
	var shippingTerms []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetShippingTermsEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		shippingTerms = append(shippingTerms, enum)
	}

	// Base - Shipping Method
	var shippingMethod []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetShippingMethodEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		shippingMethod = append(shippingMethod, enum)
	}

	// Item - Item Type
	var itemType []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetItemTypeEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		itemType = append(itemType, enum)
	}

	// Item - Item Status
	var itemStatus []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetItemStatusEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		itemStatus = append(itemStatus, enum)
	}

	// Entity - Entity Type
	var entityType []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetEntityTypeEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		entityType = append(entityType, enum)
	}

	// Entity - Entity Status
	var entityStatus []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetEntityStatusEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		entityStatus = append(entityStatus, enum)
	}

	// Inventory - Inventory Transaction Type
	var inventoryTransactionType []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetInventoryTransactionTypeEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		inventoryTransactionType = append(inventoryTransactionType, enum)
	}

	// Purchase - Purchase Order Status
	var purchaseOrderStatus []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetPurchaseOrderStatusEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		purchaseOrderStatus = append(purchaseOrderStatus, enum)
	}

	// Sale - Sales Order Status
	var salesOrderStatus []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetSalesOrderStatusEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		salesOrderStatus = append(salesOrderStatus, enum)
	}

	// Sale - Sales Quotation Status
	var salesQuotationStatus []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetSalesQuotationStatusEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		salesQuotationStatus = append(salesQuotationStatus, enum)
	}

	// Sale - Sales Shipping Status
	var salesShippingStatus []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetSalesShippingStatusEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		salesShippingStatus = append(salesShippingStatus, enum)
	}

	data := &dto.EnumsDTO{
		BaseEnums: dto.BaseEnumDTO{
			PaymentTerms:    paymentTerms,
			PaymentStatus:   paymentStatus,
			ShippingTerms:   shippingTerms,
			ShippingMethods: shippingMethod,
		},
		ItemEnums: dto.ItemEnumDTO{
			ItemType:   itemType,
			ItemStatus: itemStatus,
		},
		EntityEnums: dto.EntityEnumDTO{
			EntityType:   entityType,
			EntityStatus: entityStatus,
		},
		InventoryEnums: dto.InventoryEnumDTO{
			InventoryTransactionType: inventoryTransactionType,
		},
		PurchaseEnums: dto.PurchaseEnumDTO{
			PurchaseOrderStatus: purchaseOrderStatus,
		},
		SaleEnums: dto.SaleEnumDTO{
			SalesOrderStatus:     salesOrderStatus,
			SalesQuotationStatus: salesQuotationStatus,
			SalesShippingStatus:  salesShippingStatus,
		},
	}
	resp := s.NewListEnumResponse(*data)
	return resp, nil
}

func parseEnumForName(enumStr string) string {
	name := util.Util.NC.ToPascalCase(enumStr)
	name = util.Util.NC.PascalCaseToWords(name)
	name = util.Util.Str.CapitalizeWords(name)
	return name
}

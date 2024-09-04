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
	Payload dto.EnumsDTO
}

func (s *EnumService) NewListEnumResponse(payload dto.EnumsDTO) *ListEnumResponse {
	return &ListEnumResponse{
		Payload: payload,
	}
}

func (s *EnumService) ListEnum(req *ListEnumRequest) (*ListEnumResponse, error) {
	var paymentTerms []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetPaymentTermsEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		paymentTerms = append(paymentTerms, enum)
	}
	var paymentStatus []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetPaymentStatusEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		paymentStatus = append(paymentStatus, enum)
	}
	var itemType []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetItemTypeEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		itemType = append(itemType, enum)
	}
	var entityType []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetEntityTypeEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		entityType = append(entityType, enum)
	}
	var inventoryTransactionType []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetInventoryTransactionTypeEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		inventoryTransactionType = append(inventoryTransactionType, enum)
	}
	var purchaseOrderStatus []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetPurchaseOrderStatusEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		purchaseOrderStatus = append(purchaseOrderStatus, enum)
	}
	var salesOrderStatus []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetSalesOrderStatusEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		salesOrderStatus = append(salesOrderStatus, enum)
	}
	var salesQuotationStatus []dto.Enum
	for _, enumStr := range s.Repo.Enum.GetSalesQuotationStatusEnums() {
		enum := dto.Enum{
			Name:  parseEnumForName(enumStr.String()),
			Value: enumStr.String(),
		}
		salesQuotationStatus = append(salesQuotationStatus, enum)
	}
	data := &dto.EnumsDTO{
		BaseEnums: dto.BaseEnumDTO{
			PaymentTerms:  paymentTerms,
			PaymentStatus: paymentStatus,
		},
		ItemEnums: dto.ItemEnumDTO{
			ItemType: itemType,
		},
		EntityEnums: dto.EntityEnumDTO{
			EntityType: entityType,
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

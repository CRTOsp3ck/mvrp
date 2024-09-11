package sale

import (
	"context"
	"fmt"
	"mvrp/data/model/base"
	"mvrp/data/model/inventory"
	"mvrp/data/model/sale"
	"mvrp/domain/dto"
	entityService "mvrp/domain/service/entity"
	enumService "mvrp/domain/service/enum"
	inventoryService "mvrp/domain/service/inventory"
	saleService "mvrp/domain/service/sale"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func SeedSalesOrder(count int) error {
	// define services
	invSvc := inventoryService.NewInventoryService()
	saleSvc := saleService.NewSaleService()
	entSvc := entityService.NewEntityService()
	enumSvc := enumService.NewEnumService()

	// get all entities
	lsEntReq := entSvc.NewListEntityRequest(context.Background())
	lsEntResp, err := entSvc.ListEntity(lsEntReq)
	if err != nil {
		return err
	}

	// get all inventories
	lsInvReq := invSvc.NewListInventoryByItemTypeRequest(context.Background(), "product")
	lsInvResp, err := invSvc.ListInventory(lsInvReq)
	if err != nil {
		return err
	}
	invs := lsInvResp.Payload

	// get all enums
	lsEnumReq := enumSvc.NewListEnumRequest(context.Background())
	enumResp, err := enumSvc.ListEnum(lsEnumReq)
	if err != nil {
		return err
	}
	// get all shipping terms
	allShipTerms := enumResp.Payload.BaseEnums.ShippingTerms
	// get all shipping methods
	allShipMethods := enumResp.Payload.BaseEnums.ShippingMethods
	// get all payment terms
	allPayTerms := enumResp.Payload.BaseEnums.PaymentTerms
	// get all payment status
	allPayStatus := enumResp.Payload.BaseEnums.PaymentStatus

	// proprietor entity data
	proprietorData := lsEntResp.Payload[0]

	for i := 0; i < count; i++ {
		// get inventories to create sales quotation item
		itemsCount := gofakeit.Number(1, 10)
		selectedInvs := make([]*inventory.Inventory, itemsCount)
		for j := 0; j < itemsCount; j++ {
			invData := getRandomInventoryData(selectedInvs, invs)
			selectedInvs[j] = invData
		}

		// prepare sales quotation items
		crSoItDto := make([]dto.CreateSalesOrderItemDTO, itemsCount)
		for j, inv := range selectedInvs {
			// prepare sales quotation item
			qty := decimal.New(int64(gofakeit.Number(1, 10)*100), 2)
			unitPriceF64, ok := inv.PricePerUnit.Float64()
			if !ok {
				return fmt.Errorf("failed to convert price per unit to float64")
			}
			unitDiscAmt := decimal.New(int64(gofakeit.Number(0, 10)*100), 2)
			unitTaxAmt := decimal.New(int64(6), 2)
			unitSfAmt := decimal.New(int64(unitPriceF64*(gofakeit.Float64Range(5, 15)/100)*100), 2)
			crSoItDto[j] = dto.CreateSalesOrderItemDTO{
				BaseDocumentItem: base.BaseDocumentItem{
					ItemID:             inv.ItemID,
					Quantity:           types.NewNullDecimal(qty),
					UnitPrice:          types.NewNullDecimal(inv.PricePerUnit.Big),
					UnitDiscountAmount: types.NewNullDecimal(unitDiscAmt),
					UnitTaxAmount:      types.NewNullDecimal(unitTaxAmt),
					UnitShippingFees:   types.NewNullDecimal(unitSfAmt),
				},
				SalesOrderItem: sale.SalesOrderItem{},
			}
		}

		// create sales quotation
		additionalDiscAmt := decimal.New(int64(gofakeit.Number(0, 10)*100), 2)
		otherFees := decimal.New(int64(gofakeit.Number(0, 10)*100), 2)
		customAdjAmt := decimal.New(int64(gofakeit.Number(0, 10)*100), 2)
		shipTerm := allShipTerms[gofakeit.Number(0, len(allShipTerms)-1)]
		shipMethod := allShipMethods[gofakeit.Number(0, len(allShipMethods)-1)]
		payTerm := allPayTerms[gofakeit.Number(0, len(allPayTerms)-1)]
		payStatus := allPayStatus[gofakeit.Number(0, len(allPayStatus)-1)]
		shipDate := gofakeit.FutureDate().Add(time.Hour * 24 * 7)

		shipToInfoJson := fmt.Sprintf(`{"address": "%s"}`, gofakeit.Address().Address)
		shipToInfoByte := []byte(shipToInfoJson)
		shipFromInfoJson := fmt.Sprintf(`{"address": "%s"}`, proprietorData.Address.String)
		shipFromInfoByte := []byte(shipFromInfoJson)

		// requestedBy := gofakeit.Person()
		// requestedByJson := fmt.Sprintf(
		// 	`{"name": "%s", "phone": "%s", "email": "%s"}`,
		// 	fmt.Sprintf("%s %s", requestedBy.FirstName, requestedBy.LastName),
		// 	requestedBy.Contact.Phone,
		// 	requestedBy.Contact.Email,
		// )
		// requestedByByte := []byte(requestedByJson)

		crSoDto := dto.CreateSalesOrderDTO{
			BaseDocument: base.BaseDocument{
				IssueDate:                null.TimeFrom(time.Now()),
				AdditionalDiscountAmount: types.NewNullDecimal(additionalDiscAmt),
				OtherFees:                types.NewNullDecimal(otherFees),
				CustomAdjustmentAmount:   types.NewNullDecimal(customAdjAmt),
				ShippingTerms:            base.ShippingTerms(shipTerm.Value),
				ShippingMethod:           base.ShippingMethod(shipMethod.Value),
				ShippingDate:             null.TimeFrom(shipDate),
				PaymentTerms:             base.PaymentTerms(payTerm.Value),
				PaymentStatus:            base.PaymentStatus(payStatus.Value),
				PaymentInstructions:      null.StringFrom(gofakeit.Sentence(5)),
				Remarks:                  null.StringFrom(gofakeit.Sentence(10)),
				TermsAndConditions:       null.StringFrom(gofakeit.Sentence(20)),
			},
			SalesOrder: sale.SalesOrder{
				VendorID:                      null.IntFrom(1),
				CustomerID:                    null.IntFrom(getRandomCustomerData(lsEntResp.Payload).ID),
				SalesRepresentativeEmployeeID: null.IntFrom(getRandomEmployeeData(lsEntResp.Payload).ID),
				ShipToInformation:             null.JSONFrom(shipToInfoByte),
				ShipFromInformation:           null.JSONFrom(shipFromInfoByte),
				PaymentDueDate:                null.TimeFrom(shipDate.Add(time.Hour * 24 * 14)), // 2 weeks after shipping date
				OrderStatus:                   sale.SalesOrderStatusAccepted,
				// RequestedBy:          null.JSONFrom(requestedByByte),
			},
			Items: crSoItDto,
		}
		crSoReq := saleSvc.NewCreateSalesOrderRequest(context.Background(), crSoDto)
		crSoResp, err := saleSvc.CreateSalesOrder(crSoReq)
		if err != nil {
			return err
		}

		fmt.Println("Sales Order created ID: ", crSoResp.Payload.ID)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}

	return nil
}

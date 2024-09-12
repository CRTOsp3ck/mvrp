package sale

import (
	"context"
	"fmt"
	"mvrp/data/model/base"
	"mvrp/data/model/sale"
	"mvrp/domain/dto"
	enumService "mvrp/domain/service/enum"
	saleService "mvrp/domain/service/sale"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jinzhu/copier"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func SeedDeliveryNote() error {
	// define services
	saleSvc := saleService.NewSaleService()
	enumSvc := enumService.NewEnumService()

	// all shipping status enums
	lsShpStReq := enumSvc.NewListEnumRequest(context.Background())
	lsShpStResp, err := enumSvc.ListEnum(lsShpStReq)
	if err != nil {
		return err
	}
	AllShippingStatusEnums := lsShpStResp.Payload.SaleEnums.SalesShippingStatus

	// get all sales orders with even number id to create delivery note
	lsAllSovReq := saleSvc.NewListSalesOrderViewRequest(context.Background())
	lsAllSovResp, err := saleSvc.ListSalesOrderView(lsAllSovReq)
	if err != nil {
		return err
	}
	sovsToDeliver := make([]*sale.SalesOrderView, 0)
	for _, sov := range lsAllSovResp.Payload {
		if sov.ID.Int%2 == 0 {
			sovsToDeliver = append(sovsToDeliver, sov)
		}
	}

	// create delivery note for each sales order
	for _, sov := range sovsToDeliver {
		// sales order view base document
		var sovBd base.BaseDocument
		err = sov.BaseDocument.Unmarshal(&sovBd)
		if err != nil {
			return err
		}
		// fmt.Println("Sales Order View Base Document")
		// util.Util.Json.PrintJson(sovBd)

		// delivery note base document
		var dnBd base.BaseDocument
		err = copier.Copy(&dnBd, &sovBd)
		if err != nil {
			return err
		}
		// fmt.Println("Delivery Note Base Document")
		// util.Util.Json.PrintJson(dnBd)

		// create delivery note item DTOs
		var sovItems sale.SalesOrderItemViewSlice
		err = sov.SalesOrderItems.Unmarshal(&sovItems)
		if err != nil {
			return err
		}
		// fmt.Println("Sales Order View Items")
		// util.Util.Json.PrintJson(sovItems)

		var dnItDtos []dto.CreateDeliveryNoteItemDTO
		for _, sovIt := range sovItems {
			// sales order view item base document item
			var sovItBdi base.BaseDocumentItem
			err = sovIt.BaseDocumentItem.Unmarshal(&sovItBdi)
			if err != nil {
				return err
			}
			// delivery note item base document item
			var dnItBdi base.BaseDocumentItem
			err = copier.Copy(&dnItBdi, &sovItBdi)
			if err != nil {
				return err
			}
			// create delivery note item DTO
			dnItDto := dto.CreateDeliveryNoteItemDTO{
				BaseDocumentItem: dnItBdi,
				DeliveryNoteItem: sale.DeliveryNoteItem{
					GoodsCondition: "Good condition as received",
				},
			}
			dnItDtos = append(dnItDtos, dnItDto)
		}
		// prepare delivery note
		var shipToInfoJRM types.JSON
		err = shipToInfoJRM.Marshal(sov.ShipToInformation.JSON)
		if err != nil {
			return err
		}
		var shipFromInfoJRM types.JSON
		err = shipFromInfoJRM.Marshal(sov.ShipFromInformation.JSON)
		if err != nil {
			return err
		}
		shipPrsInfo := fmt.Sprintf(`{"name": "%s", "phone": "%s", "email": "%s"}`, gofakeit.Name(), gofakeit.Phone(), gofakeit.Email())
		recvByInfo := fmt.Sprintf(`{"name": "%s", "phone": "%s", "email": "%s"}`, gofakeit.Name(), gofakeit.Phone(), gofakeit.Email())
		dn := sale.DeliveryNote{
			SalesOrderID:                 sov.ID.Int,
			VendorID:                     null.IntFrom(1),
			CustomerID:                   sov.CustomerID,
			ShipToInformation:            shipToInfoJRM,
			ShipFromInformation:          shipFromInfoJRM,
			BillToInformation:            shipToInfoJRM,
			ShippingDate:                 dnBd.ShippingDate.Time,
			ShippingPersonnelInformation: null.JSONFrom([]byte(shipPrsInfo)),
			ShippingStatus:               sale.ShippingStatus(AllShippingStatusEnums[gofakeit.Number(0, len(AllShippingStatusEnums)-1)].Value),
			ReceivedBy:                   null.JSONFrom([]byte(recvByInfo)),
			OverallGoodsCondition:        null.StringFrom("Good condition as received"),
		}
		crDnDto := dto.CreateDeliveryNoteDTO{
			BaseDocument: dnBd,
			DeliveryNote: dn,
			Items:        dnItDtos,
		}

		// create delivery note
		crDnReq := saleSvc.NewCreateDeliveryNoteRequest(context.Background(), crDnDto)
		crDnResp, err := saleSvc.CreateDeliveryNote(crDnReq)
		if err != nil {
			return err
		}

		fmt.Println("Delivery Note created ID: ", crDnResp.Payload.ID)
	}

	return nil
}

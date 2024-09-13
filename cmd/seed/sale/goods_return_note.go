package sale

import (
	"context"
	"fmt"
	"mvrp/data/model/base"
	"mvrp/data/model/sale"
	"mvrp/domain/dto"
	entityService "mvrp/domain/service/entity"
	saleService "mvrp/domain/service/sale"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func SeedGoodsReturnNote() error {
	// define services
	saleSvc := saleService.NewSaleService()
	entitySvc := entityService.NewEntityService()

	// entity list
	lsEntReq := entitySvc.NewListEntityRequest(context.Background())
	lsEntResp, err := entitySvc.ListEntity(lsEntReq)
	if err != nil {
		return err
	}
	entities := lsEntResp.Payload

	// get proprietor data
	getEntReq := entitySvc.NewGetEntityRequest(context.Background(), 1)
	getEntResp, err := entitySvc.GetEntity(getEntReq)
	if err != nil {
		return err
	}
	proprietor := getEntResp.Payload

	// get all sales order with delivery notes
	searchDnDto := dto.SearchDeliveryNoteDTO{
		ShippingStatus: "shipped",
		ItemsPerPage:   100,
		Page:           1,
		SortBy:         "desc",
		OrderBy:        "created_at",
	}
	srDnReq := saleSvc.NewSearchDeliveryNoteViewRequest(context.Background(), searchDnDto)
	srDnResp, err := saleSvc.SearchDeliveryNoteView(srDnReq)
	if err != nil {
		return err
	}

	for _, dnv := range srDnResp.Payload {
		// base document
		var bd base.BaseDocument
		err = dnv.BaseDocument.Unmarshal(&bd)
		if err != nil {
			return err
		}

		// delivery note items
		var dnvItems sale.DeliveryNoteItemViewSlice
		err = dnv.DeliveryNoteItems.Unmarshal(&dnvItems)
		if err != nil {
			return err
		}
		var grnItemDtos []dto.CreateGoodsReturnNoteItemDTO
		for _, dnvItem := range dnvItems {
			// base document item
			var bdi base.BaseDocumentItem
			err = dnvItem.BaseDocumentItem.Unmarshal(&bdi)
			if err != nil {
				return err
			}
			// goods return note item
			bdiQty, ok := bdi.Quantity.Float64()
			if !ok {
				return fmt.Errorf("failed to convert quantity to float64")
			}

			returnQty := gofakeit.Number(1, int(bdiQty))
			returnCondition := gofakeit.RandomString([]string{"good", "acceptable", "bad"})
			returnReason := gofakeit.RandomString([]string{"damaged", "wrong item", "not needed"})

			bdi.Quantity = types.NewNullDecimal(decimal.New(int64(returnQty)*100, 2))

			grnItem := sale.GoodsReturnNoteItem{
				ReturnQuantity:  types.NewNullDecimal(decimal.New(int64(returnQty)*100, 2)),
				ReturnCondition: null.StringFrom(returnCondition),
				ReturnReason:    null.StringFrom(returnReason),
			}
			grnItemDtos = append(grnItemDtos, dto.CreateGoodsReturnNoteItemDTO{
				BaseDocumentItem:    bdi,
				GoodsReturnNoteItem: grnItem,
			})
		}

		// get random employee
		emp := getRandomEmployeeData(entities)

		// create goods return note
		recvLoc := fmt.Sprintf(`{"address": "%s"}`, proprietor.Address.String)
		grnDto := dto.CreateGoodsReturnNoteDTO{
			BaseDocument: bd,
			Items:        grnItemDtos,
			GoodsReturnNote: sale.GoodsReturnNote{
				ReturnDate:                   time.Now(),
				ReturnedByCustomerID:         dnv.CustomerID,
				ReceivingLocationInformation: null.JSONFrom([]byte(recvLoc)),
				ReceivedByEmployeeID:         null.IntFrom(emp.ID),
				OverallGoodsCondition:        null.StringFrom(gofakeit.RandomString([]string{"good", "acceptable", "bad"})),
			},
		}
		crGrnReq := saleSvc.NewCreateGoodsReturnNoteRequest(context.Background(), grnDto)
		crGrnResp, err := saleSvc.CreateGoodsReturnNote(crGrnReq)
		if err != nil {
			return err
		}
		fmt.Printf("Goods Return Note created: %d\n", crGrnResp.Payload.ID)
	}
	return nil
}

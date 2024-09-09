package inventory

import (
	"context"
	"fmt"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
	inventoryService "mvrp/domain/service/inventory"
	itemService "mvrp/domain/service/item"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func SeedReturnMerchandiseAuthorization() error {
	// define services
	invSvc := inventoryService.NewInventoryService()

	// get all inventories
	lsInvReq := invSvc.NewListInventoryByItemTypeRequest(context.Background(), "product")
	lsInvResp, err := invSvc.ListInventory(lsInvReq)
	if err != nil {
		return err
	}
	invs := lsInvResp.Payload

	// create return merchandise authorization for each inventory
	for _, inv := range invs {
		// get the item
		itemSvc := itemService.NewItemService()
		itemReq := itemSvc.NewGetItemRequest(context.Background(), inv.ItemID.Int)
		itemResp, err := itemSvc.GetItem(itemReq)
		if err != nil {
			return err
		}
		item := itemResp.Payload

		// prepare return merchandise authorization item
		qty := decimal.New(int64(gofakeit.Number(1, 10)*100), 2)
		crRmaItemDto := dto.CreateReturnMerchandiseAuthorizationItemDTO{
			ReturnMerchandiseAuthorizationItem: inventory.ReturnMerchandiseAuthorizationItem{
				InventoryID: null.IntFrom(inv.ID),
				Quantity:    types.NewDecimal(qty),
				UnitValue:   types.NewDecimal(item.Price.Big),
			},
		}

		// create return merchandise authorization
		crRmaDto := dto.CreateReturnMerchandiseAuthorizationDTO{
			ReturnMerchandiseAuthorization: inventory.ReturnMerchandiseAuthorization{
				RmaDate: null.TimeFrom(time.Now()),
				Notes:   null.StringFrom(gofakeit.Sentence(10)),
			},
			Items: []dto.CreateReturnMerchandiseAuthorizationItemDTO{crRmaItemDto},
		}
		crRmaReq := invSvc.NewCreateReturnMerchandiseAuthorizationRequest(context.Background(), crRmaDto)
		crRmaResp, err := invSvc.CreateReturnMerchandiseAuthorization(crRmaReq)
		if err != nil {
			return err
		}
		rma := crRmaResp.Payload

		// get created return merchandise authorization
		getRmaReq := invSvc.NewGetReturnMerchandiseAuthorizationRequest(context.Background(), rma.ID)
		getRmaResp, err := invSvc.GetReturnMerchandiseAuthorization(getRmaReq)
		if err != nil {
			return err
		}
		rmaItems := make([]dto.UpdateReturnMerchandiseAuthorizationItemDTO, len(getRmaResp.Payload.Items))
		for i, item := range getRmaResp.Payload.Items {
			rmaItems[i] = dto.UpdateReturnMerchandiseAuthorizationItemDTO(item)
		}

		// update return merchandise authorization number
		rmaNumber := fmt.Sprintf("GIN%03d", inv.ID)
		getRmaResp.Payload.ReturnMerchandiseAuthorization.RmaNumber = rmaNumber
		upRmaDto := dto.UpdateReturnMerchandiseAuthorizationDTO{
			ReturnMerchandiseAuthorization: getRmaResp.Payload.ReturnMerchandiseAuthorization,
			Items:                          rmaItems,
		}
		updRmaReq := invSvc.NewUpdateReturnMerchandiseAuthorizationRequest(context.Background(), upRmaDto)
		updRmaResp, err := invSvc.UpdateReturnMerchandiseAuthorization(updRmaReq)
		if err != nil {
			return err
		}

		fmt.Println("Return Merchandise Authorization created ID: ", updRmaResp.Payload.ID)
		time.Sleep(1 * time.Millisecond)
	}
	return nil
}

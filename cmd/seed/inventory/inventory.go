package inventory

import (
	"context"
	"fmt"
	"mvrp/data/model/inventory"
	"mvrp/data/model/item"
	"mvrp/domain/dto"
	inventoryService "mvrp/domain/service/inventory"
	itemService "mvrp/domain/service/item"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func SeedInventory() error {
	// define services
	itemSvc := itemService.NewItemService()
	invSvc := inventoryService.NewInventoryService()

	// get all items
	itemType := item.ItemTypeProduct
	lsItemReq := itemSvc.NewListItemByTypeRequest(context.Background(), &itemType)
	items, err := itemSvc.ListItemByType(lsItemReq)
	if err != nil {
		return err
	}

	// create inventory for each item
	for _, item := range items.Payload {
		qty := decimal.New(int64(gofakeit.Number(100, 1000)*100), 2)
		reorderLevel := decimal.New(int64(gofakeit.Number(10, 50)*100), 2)
		reorderQty := decimal.New(int64(gofakeit.Number(500, 1000)*100), 2)
		crDto := dto.CreateInventoryDTO{
			Inventory: inventory.Inventory{
				ItemID:            null.IntFrom(item.ID),
				QuantityAvailable: types.NewNullDecimal(qty),
				CostPerUnit:       item.Cost,
				PricePerUnit:      item.Price,
				ReorderLevel:      types.NewNullDecimal(reorderLevel),
				ReorderQuantity:   types.NewNullDecimal(reorderQty),
				Remarks:           null.StringFrom(gofakeit.Sentence(10)),
			},
		}
		crInvReq := invSvc.NewCreateInventoryRequest(context.Background(), crDto)
		crInvResp, err := invSvc.CreateInventory(crInvReq)
		if err != nil {
			return err
		}

		// update inventory number using the id
		inv := crInvResp.Payload
		inv.InventoryNumber = fmt.Sprintf("INV%03d", inv.ID)
		updInvReq := invSvc.NewUpdateInventoryRequest(context.Background(), dto.UpdateInventoryDTO{
			Inventory: inv,
		})
		updInvResp, err := invSvc.UpdateInventory(updInvReq)
		if err != nil {
			return err
		}

		fmt.Println("Inventory created ID: ", updInvResp.Payload.ID)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
	return nil
}

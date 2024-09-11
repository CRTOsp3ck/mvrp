package inventory

import (
	"context"
	"fmt"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
	entityService "mvrp/domain/service/entity"
	inventoryService "mvrp/domain/service/inventory"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func SeedStockCountSheet() error {
	// define services
	invSvc := inventoryService.NewInventoryService()

	// get all inventories
	lsInvReq := invSvc.NewListInventoryByItemTypeRequest(context.Background(), "product")
	lsInvResp, err := invSvc.ListInventory(lsInvReq)
	if err != nil {
		return err
	}
	invs := lsInvResp.Payload

	// get all entities
	entSvc := entityService.NewEntityService()
	lsEntReq := entSvc.NewListEntityRequest(context.Background())
	lsEntResp, err := entSvc.ListEntity(lsEntReq)
	if err != nil {
		return err
	}

	// create stock count sheet for each inventory
	for _, inv := range invs {
		invQty, ok := inv.QuantityAvailable.Float64()
		if !ok {
			return fmt.Errorf("failed to convert quantity available to float64")
		}
		rangeBuff := invQty / 10
		if rangeBuff < 1 {
			rangeBuff = 1
		}
		countedInvQty := gofakeit.Number(int(invQty-rangeBuff), int(invQty+rangeBuff))

		// create goods issue note
		crScsDto := dto.CreateStockCountSheetDTO{
			StockCountSheet: inventory.StockCountSheet{
				InventoryID:         null.IntFrom(inv.ID),
				CountDate:           null.TimeFrom(time.Now()),
				CountedByEmployeeID: null.IntFrom(getRandomEmployeeData(lsEntResp.Payload).ID),
				TotalQuantity:       types.NewDecimal(decimal.New(int64(countedInvQty*100), 2)),
			},
		}
		crScsReq := invSvc.NewCreateStockCountSheetRequest(context.Background(), crScsDto)
		crScsResp, err := invSvc.CreateStockCountSheet(crScsReq)
		if err != nil {
			return err
		}
		scs := crScsResp.Payload

		// get created stock count sheet
		getScsReq := invSvc.NewGetStockCountSheetRequest(context.Background(), scs.ID)
		getScsResp, err := invSvc.GetStockCountSheet(getScsReq)
		if err != nil {
			return err
		}

		// update stock count sheet number
		scsNumber := fmt.Sprintf("SCS%03d", inv.ID)
		getScsResp.Payload.SCSNumber = scsNumber
		upScsDto := dto.UpdateStockCountSheetDTO{
			StockCountSheet: getScsResp.Payload,
		}
		updScsReq := invSvc.NewUpdateStockCountSheetRequest(context.Background(), upScsDto)
		updScsResp, err := invSvc.UpdateStockCountSheet(updScsReq)
		if err != nil {
			return err
		}

		fmt.Println("Stock Count Sheet created ID: ", updScsResp.Payload.ID)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
	return nil
}

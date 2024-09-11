package inventory

import (
	"context"
	"fmt"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
	entityService "mvrp/domain/service/entity"
	inventoryService "mvrp/domain/service/inventory"
	itemService "mvrp/domain/service/item"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func SeedGoodsIssueNote() error {
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

	// create goods issue note for each inventory
	for _, inv := range invs {
		// get the item
		itemSvc := itemService.NewItemService()
		itemReq := itemSvc.NewGetItemRequest(context.Background(), inv.ItemID.Int)
		itemResp, err := itemSvc.GetItem(itemReq)
		if err != nil {
			return err
		}
		item := itemResp.Payload

		// prepare goods issue note item
		qty := decimal.New(int64(gofakeit.Number(1, 10)*100), 2)
		crGinItemDto := dto.CreateGoodsIssueNoteItemDTO{
			GoodsIssueNoteItem: inventory.GoodsIssueNoteItem{
				InventoryID: null.IntFrom(inv.ID),
				Quantity:    types.NewDecimal(qty),
				UnitValue:   types.NewDecimal(item.Price.Big),
			},
		}

		// create goods issue note
		crGinDto := dto.CreateGoodsIssueNoteDTO{
			GoodsIssueNote: inventory.GoodsIssueNote{
				IssueDate:    null.TimeFrom(time.Now()),
				Notes:        null.StringFrom(gofakeit.Sentence(10)),
				ReceipientID: null.IntFrom(getRandomCustomerData(lsEntResp.Payload).ID),
			},
			Items: []dto.CreateGoodsIssueNoteItemDTO{crGinItemDto},
		}
		crGinReq := invSvc.NewCreateGoodsIssueNoteRequest(context.Background(), crGinDto)
		crGinResp, err := invSvc.CreateGoodsIssueNote(crGinReq)
		if err != nil {
			return err
		}
		gin := crGinResp.Payload

		// get created goods issue note
		getGinReq := invSvc.NewGetGoodsIssueNoteRequest(context.Background(), gin.ID)
		getGinResp, err := invSvc.GetGoodsIssueNote(getGinReq)
		if err != nil {
			return err
		}
		ginItems := make([]dto.UpdateGoodsIssueNoteItemDTO, len(getGinResp.Payload.Items))
		for i, item := range getGinResp.Payload.Items {
			ginItems[i] = dto.UpdateGoodsIssueNoteItemDTO(item)
		}

		// update goods issue note number
		ginNumber := fmt.Sprintf("GIN%03d", inv.ID)
		getGinResp.Payload.GoodsIssueNote.GinNumber = ginNumber
		upGinDto := dto.UpdateGoodsIssueNoteDTO{
			GoodsIssueNote: getGinResp.Payload.GoodsIssueNote,
			Items:          ginItems,
		}
		updGinReq := invSvc.NewUpdateGoodsIssueNoteRequest(context.Background(), upGinDto)
		updGinResp, err := invSvc.UpdateGoodsIssueNote(updGinReq)
		if err != nil {
			return err
		}

		fmt.Println("Goods Issue Note created ID: ", updGinResp.Payload.ID)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
	return nil
}

package item

import (
	"context"
	"fmt"
	model "mvrp/data/model/item"
	"mvrp/domain/dto"
	"mvrp/domain/service/item"
	"mvrp/util"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func SeedServices(count int) error {
	svc := item.NewItemService()
	for i := 0; i < count; i++ {
		refProd := gofakeit.Product()
		price := gofakeit.Price(500, 1000)
		service := fmt.Sprintf("%s %s %s Service", gofakeit.Verb(), gofakeit.AdjectiveDescriptive(), gofakeit.Noun())
		data := dto.CreateItemDTO{
			Item: model.Item{
				Code:        fmt.Sprintf("SER%03d", i+1),
				Sku:         refProd.UPC,
				Brand:       gofakeit.Car().Brand,
				Category:    refProd.Categories[0],
				Name:        util.Util.Str.CapitalizeWords(service),
				Description: gofakeit.ProductDescription(),
				Origin:      gofakeit.Country(),
				Cost:        types.NewNullDecimal(decimal.New(int64(price)*75, 2)),
				Price:       types.NewNullDecimal(decimal.New(int64(price)*100, 2)),
				Type:        model.ItemTypeService,
				Status:      model.ItemStatusActive,
			},
		}
		ctx := context.Background()
		req := svc.NewCreateItemRequest(ctx, data)
		resp, err := svc.CreateItem(req)
		if err != nil {
			return err
		}

		fmt.Println("Service created ID: ", resp.Payload.ID)
		time.Sleep(1 * time.Millisecond)
	}
	return nil
}

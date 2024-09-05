package item

import (
	"context"
	"fmt"
	model "mvrp/data/model/item"
	"mvrp/domain/dto"
	"mvrp/domain/service/item"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func SeedProducts(count int) error {
	svc := item.NewItemService()

	for i := 0; i < count; i++ {
		fp := gofakeit.Product()
		data := dto.CreateItemDTO{
			Item: model.Item{
				Code:        fmt.Sprintf("PRO%03d", i+1),
				Sku:         fp.UPC,
				Brand:       gofakeit.Car().Brand,
				Category:    fp.Categories[0],
				Name:        fp.Name,
				Description: fp.Description,
				Origin:      gofakeit.Country(),
				Cost:        types.NewNullDecimal(decimal.New(int64(fp.Price)*75, 2)),
				Price:       types.NewNullDecimal(decimal.New(int64(fp.Price)*100, 2)),
				Type:        model.ItemTypeProduct,
				Status:      model.ItemStatusActive,
			},
		}
		ctx := context.Background()
		req := svc.NewCreateItemRequest(ctx, data)
		resp, err := svc.CreateItem(req)
		if err != nil {
			return err
		}

		fmt.Println("Product created ID: ", resp.Payload.ID)
		time.Sleep(10 * time.Millisecond)
	}
	return nil
}

package entity

import (
	"context"
	"fmt"
	model "mvrp/data/model/entity"
	"mvrp/domain/dto"
	"mvrp/domain/service/entity"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/volatiletech/null/v8"
)

func SeedCustomers(count int) error {
	// Create a new service
	svc := entity.NewEntityService()

	for i := 0; i < count; i++ {
		// Create a new request
		data := dto.CreateEntityDTO{
			Entity: model.Entity{
				Code:        fmt.Sprintf("CUS%03d", i+1),
				Name:        gofakeit.Name(),
				Description: gofakeit.AdjectiveDescriptive(),
				Address:     null.StringFrom(gofakeit.Address().Address),
				Phone:       null.StringFrom(gofakeit.Phone()),
				Email:       null.StringFrom(gofakeit.Email()),
				Website:     null.StringFrom(gofakeit.URL()),
				Type:        model.EntityTypeCustomer,
				Status:      model.EntityStatusActive,
			},
		}
		ctx := context.Background()
		req := svc.NewCreateEntityRequest(ctx, data)

		// Create a new customer
		resp, err := svc.CreateEntity(req)
		if err != nil {
			return err
		}

		fmt.Println("Customer created ID: ", resp.Payload.ID)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}

	return nil
}

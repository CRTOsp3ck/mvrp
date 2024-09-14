package sale

import (
	"mvrp/data/model/entity"
	"mvrp/data/model/inventory"

	"github.com/brianvoe/gofakeit/v7"
)

func getRandomInventoryData(currIs inventory.InventorySlice, is inventory.InventorySlice) *inventory.Inventory {
	invData := is[gofakeit.Number(0, len(is)-1)]
	for _, inv := range currIs {
		if inv == nil {
			continue
		}
		if inv.ID == invData.ID {
			return getRandomInventoryData(currIs, is)
		}
	}
	return invData
}

func getRandomCustomerData(es entity.EntitySlice) *entity.Entity {
	cusData := es[gofakeit.Number(0, len(es)-1)]
	// if its not a customer, get another one
	if cusData.EntityType != "customer" {
		return getRandomCustomerData(es)
	}
	return cusData
}

func getRandomEmployeeData(es entity.EntitySlice) *entity.Entity {
	empData := es[gofakeit.Number(0, len(es)-1)]
	// if its not an employee, get another one
	if empData.EntityType != "employee" {
		return getRandomEmployeeData(es)
	}
	return empData
}

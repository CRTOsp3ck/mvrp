package inventory

import (
	"mvrp/data/model/entity"

	"github.com/brianvoe/gofakeit/v7"
)

func getRandomCustomerData(es entity.EntitySlice) *entity.Entity {
	cusData := es[gofakeit.Number(0, len(es)-1)]
	// if its not a customer, get another one
	if cusData.Type != "customer" {
		return getRandomCustomerData(es)
	}
	return cusData
}

func getRandomEmployeeData(es entity.EntitySlice) *entity.Entity {
	empData := es[gofakeit.Number(0, len(es)-1)]
	// if its not an employee, get another one
	if empData.Type != "employee" {
		return getRandomEmployeeData(es)
	}
	return empData
}

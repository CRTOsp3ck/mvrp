package main

import (
	"mvrp/cmd/seed/entity"
	"mvrp/cmd/seed/inventory"
	"mvrp/cmd/seed/item"
	"mvrp/data"
	"mvrp/domain"
	"mvrp/env"
)

func main() {
	Init()
	Seed()
}

func Init() {
	env.Init()

	entity.Init()
	item.Init()
	inventory.Init()

	data.Init()
	domain.Init()
}

func Seed() {
	err := seedEntity()
	if err != nil {
		panic(err)
	}

	err = seedItem()
	if err != nil {
		panic(err)
	}

	err = seedInventory()
	if err != nil {
		panic(err)
	}
}

func seedInventory() error {
	err := inventory.SeedInventory()
	if err != nil {
		return err
	}

	err = inventory.SeedGoodsIssueNote()
	if err != nil {
		return err
	}

	err = inventory.SeedReturnMerchandiseAuthorization()
	if err != nil {
		return err
	}

	err = inventory.SeedStockCountSheet()
	if err != nil {
		return err
	}

	return nil
}

func seedItem() error {
	// Create 100 products
	err := item.SeedProducts(100)
	if err != nil {
		return err
	}

	// Create 25 services
	err = item.SeedServices(25)
	if err != nil {
		return err
	}

	return nil
}

func seedEntity() error {
	// Create 10 customers
	err := entity.SeedCustomers(250)
	if err != nil {
		return err
	}

	// Create 5 suppliers
	err = entity.SeedSuppliers(50)
	if err != nil {
		return err
	}

	// Create 3 employees
	err = entity.SeedEmployees(25)
	if err != nil {
		return err
	}

	return nil
}

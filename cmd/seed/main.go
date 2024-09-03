package main

import (
	"mvrp/cmd/seed/entity"
	"mvrp/cmd/seed/item"
	"mvrp/data"
	"mvrp/domain"
)

func main() {
	Init()
	Seed()
}

func Init() {
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
	err := entity.SeedCustomers(100)
	if err != nil {
		return err
	}

	// Create 5 suppliers
	err = entity.SeedSuppliers(25)
	if err != nil {
		return err
	}

	// Create 3 employees
	err = entity.SeedEmployees(5)
	if err != nil {
		return err
	}

	return nil
}

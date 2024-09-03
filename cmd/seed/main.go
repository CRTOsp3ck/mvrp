package main

import (
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
	// Create 10 customers
	err := seedCustomers(100)
	if err != nil {
		panic(err)
	}

	// Create 5 suppliers
	err = seedSuppliers(25)
	if err != nil {
		panic(err)
	}

	// Create 3 employees
	err = seedEmployees(5)
	if err != nil {
		panic(err)
	}
}

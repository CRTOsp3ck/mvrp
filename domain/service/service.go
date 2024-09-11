package service

import (
	"mvrp/domain/service/entity"
	"mvrp/domain/service/inventory"
	"mvrp/domain/service/invoice"
	"mvrp/domain/service/item"
	"mvrp/domain/service/sale"
)

func Init() {
	entity.Init()
	item.Init()
	inventory.Init()
	sale.Init()
	invoice.Init()
}

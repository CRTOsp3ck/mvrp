package model

import (
	"fmt"
	"mvrp/data/model/base"
	"mvrp/data/model/entity"
	"mvrp/data/model/inventory"
	"mvrp/data/model/invoice"
	"mvrp/data/model/item"
	"mvrp/data/model/purchase"
	"mvrp/data/model/sale"
	"reflect"
)

var ModelInfoSlice []ModelInfo

type ModelInfo struct {
	PackageName string
	ModelName   string
	TableName   string
}

func Init() {
	generateModelInfo()
}

func generateModelInfo() {
	ModelInfoSlice = append(ModelInfoSlice, parseModelInfo(base.TableNames, "base")...)
	ModelInfoSlice = append(ModelInfoSlice, parseModelInfo(base.ViewNames, "base")...)

	ModelInfoSlice = append(ModelInfoSlice, parseModelInfo(entity.TableNames, "entity")...)
	ModelInfoSlice = append(ModelInfoSlice, parseModelInfo(entity.ViewNames, "entity")...)

	ModelInfoSlice = append(ModelInfoSlice, parseModelInfo(inventory.TableNames, "inventory")...)
	ModelInfoSlice = append(ModelInfoSlice, parseModelInfo(inventory.ViewNames, "inventory")...)

	ModelInfoSlice = append(ModelInfoSlice, parseModelInfo(invoice.TableNames, "invoice")...)
	ModelInfoSlice = append(ModelInfoSlice, parseModelInfo(invoice.ViewNames, "invoice")...)

	ModelInfoSlice = append(ModelInfoSlice, parseModelInfo(item.TableNames, "item")...)
	ModelInfoSlice = append(ModelInfoSlice, parseModelInfo(item.ViewNames, "item")...)

	ModelInfoSlice = append(ModelInfoSlice, parseModelInfo(purchase.TableNames, "purchase")...)
	ModelInfoSlice = append(ModelInfoSlice, parseModelInfo(purchase.ViewNames, "purchase")...)

	ModelInfoSlice = append(ModelInfoSlice, parseModelInfo(sale.TableNames, "sale")...)
	ModelInfoSlice = append(ModelInfoSlice, parseModelInfo(sale.ViewNames, "sale")...)
}

func parseModelInfo(s interface{}, pn string) []ModelInfo {
	var mis []ModelInfo
	v := reflect.ValueOf(s)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		mi := ModelInfo{
			PackageName: pn,
			ModelName:   t.Field(i).Name,
			TableName:   v.Field(i).Interface().(string),
		}
		mis = append(mis, mi)
		// DebugModelInfo(mi)
	}
	return mis
}

func DebugModelInfo(mi ModelInfo) {
	fmt.Printf(
		"PackageName: %s, ModelName: %s, TableName: %s\n",
		mi.PackageName,
		mi.ModelName,
		mi.TableName,
	)
}

func DebugTableNames() {
	fmt.Println(base.TableNames)
	fmt.Println(entity.TableNames)
	fmt.Println(inventory.TableNames)
	fmt.Println(invoice.TableNames)
	fmt.Println(item.TableNames)
	fmt.Println(purchase.TableNames)
	fmt.Println(sale.TableNames)
}

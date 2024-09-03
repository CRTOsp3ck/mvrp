package ds

import (
	"fmt"
	"reflect"
)

type IDSUtil interface {
	StructToSlice(input interface{}) []interface{}
	StructToStringSlice(input interface{}) []string
	Contains(slice []interface{}, item interface{}) bool
}

type DSUtil struct{}

func (d *DSUtil) StructToSlice(input interface{}) []interface{} {
	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Struct {
		panic("input is not a struct")
	}

	var result []interface{}
	for i := 0; i < val.NumField(); i++ {
		result = append(result, val.Field(i).Interface())
	}
	return result
}

func (d *DSUtil) StructToStringSlice(input interface{}) []string {
	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Struct {
		panic("input is not a struct")
	}

	var result []string
	for i := 0; i < val.NumField(); i++ {
		result = append(result, fmt.Sprintf("%v", val.Field(i).Interface()))
	}
	return result
}

// Contains checks if a slice contains a specific item.
func (d *DSUtil) Contains(slice []interface{}, item interface{}) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

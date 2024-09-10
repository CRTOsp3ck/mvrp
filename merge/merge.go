package merge

import (
	"errors"
	"reflect"
	"strings"
)

// MergeNilOrEmptyValueFields takes two structs of the same type and updates the destination's empty fields
// (nil or empty string) with the values from the source struct, while skipping specified fields.
func MergeNilOrEmptyValueFields(src, dest interface{}, defaultSkipFields bool, skipAdditionalFields ...string) error {
	srcVal := reflect.ValueOf(src)
	destVal := reflect.ValueOf(dest)

	// Check if src and dest are pointers to structs
	if srcVal.Kind() != reflect.Ptr || destVal.Kind() != reflect.Ptr {
		return ErrNonPointerStruct
	}
	if srcVal.Elem().Kind() != reflect.Struct || destVal.Elem().Kind() != reflect.Struct {
		return ErrNonStruct
	}

	srcElem := srcVal.Elem()
	destElem := destVal.Elem()

	for i := 0; i < srcElem.NumField(); i++ {
		srcField := srcElem.Field(i)
		destField := destElem.Field(i)

		// Get the field name
		fieldName := srcElem.Type().Field(i).Name

		defaultFieldsToSkip := []string{"ID", "CreatedAt", "UpdatedAt", "DeletedAt"}
		if defaultSkipFields {
			skipAdditionalFields = append(skipAdditionalFields, defaultFieldsToSkip...)
		}

		// Check if the field is in the skip list
		if contains(skipAdditionalFields, fieldName) {
			continue // Skip this field
		}

		// Check for nil or empty string in destination fields
		if isZeroValue(destField) {
			if destField.CanSet() {
				destField.Set(srcField)
			}
		}
	}
	return nil
}

// isZeroValue checks if a value is the zero value (nil for pointers or interfaces, "" for strings)
func isZeroValue(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.Ptr, reflect.Interface:
		return val.IsNil()
	case reflect.String:
		return val.String() == ""
	default:
		return reflect.DeepEqual(val.Interface(), reflect.Zero(val.Type()).Interface())
	}
}

// contains checks if a string exists in a slice of strings (case-insensitive)
func contains(list []string, item string) bool {
	for _, elem := range list {
		if strings.EqualFold(elem, item) {
			return true
		}
	}
	return false
}

// Custom errors
var (
	ErrNonPointerStruct = errors.New("src and dest must be pointers to structs")
	ErrNonStruct        = errors.New("both src and dest must be structs")
)

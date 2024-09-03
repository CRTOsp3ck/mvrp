package nc

import (
	"strings"
	"unicode"
)

type INCUtil interface {
	ToSnakeCase(s string) string
	ToPascalCase(s string) string
	ToCamelCase(s string) string
}

type NCUtil struct{}

func (n *NCUtil) ToSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) && i != 0 {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}

func (n *NCUtil) ToPascalCase(s string) string {
	var result strings.Builder
	upperNext := true
	for _, r := range s {
		if r == '_' || r == ' ' {
			upperNext = true
			continue
		}
		if upperNext {
			result.WriteRune(unicode.ToUpper(r))
			upperNext = false
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}
	return result.String()
}

func (n *NCUtil) ToCamelCase(s string) string {
	var result strings.Builder
	upperNext := false
	for i, r := range s {
		if r == '_' || r == ' ' {
			upperNext = true
			continue
		}
		if i == 0 {
			result.WriteRune(unicode.ToLower(r))
		} else if upperNext {
			result.WriteRune(unicode.ToUpper(r))
			upperNext = false
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}
	return result.String()
}

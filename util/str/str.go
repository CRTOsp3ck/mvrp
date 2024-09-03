package str

import "strings"

type IStrUtil interface {
	IsEmpty(str string) bool
	ToString(val interface{}) string
	CapitalizeWords(str string) string
}

type StrUtil struct{}

func (s *StrUtil) IsEmpty(str string) bool {
	return str == ""
}

func (s *StrUtil) ToString(val interface{}) string {
	return val.(string)
}

func (s *StrUtil) CapitalizeWords(str string) string {
	words := strings.Fields(str)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

package str

type IStrUtil interface {
}

type StrUtil struct{}

func (s *StrUtil) IsEmpty(str string) bool {
	return str == ""
}

func (s *StrUtil) ToString(val interface{}) string {
	return val.(string)
}

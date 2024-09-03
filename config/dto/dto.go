package dto

import (
	"fmt"
	"mvrp/errors"
	"mvrp/util"
	"path/filepath"
)

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type DTO struct {
	Name                     string        `json:"name"`
	Inherits                 []Inheritance `json:"inherits"`
	AdditionalFields         []Field       `json:"additional_fields"`
	ExcludeFieldsFromInherit []string      `json:"exclude_fields_from_inherit"`
	GroupQueryFields         []Field       `json:"group_query_fields,omitempty"`
	SearchQueryFields        []Field       `json:"search_query_fields,omitempty"`
}

type Inheritance struct {
	Package string `json:"package"`
	Model   string `json:"model"`
}

type Package struct {
	Package string `json:"package"`
	DTOs    []DTO  `json:"dtos"`
}

type Root struct {
	Data []Package `json:"data"`
}

var Config *Root

func GetConfig() (*Root, error) {
	if Config != nil {
		return Config, nil
	}

	rootDir, err := util.Util.FS.FindProjectRoot("go.mod")
	if err != nil {
		return nil, err
	}
	var root Root
	ptrIface, err := util.Util.Json.ParseJsonFile(
		filepath.Join(rootDir, "config", "dto", "dto.json"), &root)
	if err != nil {
		return nil, err
	}
	rootPtr, ok := ptrIface.(*Root)
	if !ok {
		return nil, errors.WrapError(errors.ErrTypeAssertion,
			fmt.Sprintf("expected type: %T, got: %T", root, ptrIface))
	}
	Config = rootPtr
	return Config, nil
}

func GetDTO(dtoName string) (*DTO, error) {
	config, err := GetConfig()
	if err != nil {
		return nil, err
	}
	for _, pkg := range config.Data {
		for _, dto := range pkg.DTOs {
			if dto.Name == dtoName {
				return &dto, nil
			}
		}
	}
	return nil, errors.WrapError(errors.ErrTypeNotFound,
		fmt.Sprintf("DTO with name %s not found", dtoName))
}

func GetDTOs() ([]DTO, error) {
	config, err := GetConfig()
	if err != nil {
		return nil, err
	}
	var dtos []DTO
	for _, pkg := range config.Data {
		dtos = append(dtos, pkg.DTOs...)
	}
	return dtos, nil
}

func IfDtoExists(dtoName string) bool {
	dtos, err := GetDTOs()
	if err != nil {
		return false
	}
	for _, dto := range dtos {
		if dto.Name == dtoName {
			return true
		}
	}
	return false
}

func GetSearchQueryFields(dtoName string) ([]string, error) {
	dto, err := GetDTO(dtoName)
	if err != nil {
		return nil, err
	}
	var fields []string
	for _, field := range dto.SearchQueryFields {
		fields = append(fields, field.Name)
	}
	return fields, nil
}

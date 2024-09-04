package enums

import (
	"fmt"
	"mvrp/errors"
	"mvrp/util"
	"path/filepath"
)

type Enum struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Package struct {
	Package string `json:"package"`
	Enums   []Enum `json:"enums"`
}

type Root struct {
	Data []Package `json:"data"`
}

func GetConfig() (*Root, error) {
	rootDir, err := util.Util.FS.FindProjectRoot("go.mod")
	if err != nil {
		return nil, err
	}
	var root Root
	ptrIface, err := util.Util.Json.ParseJsonFile(
		filepath.Join(rootDir, "config", "enums", "enums.json"), &root)
	if err != nil {
		return nil, err
	}
	rootPtr, ok := ptrIface.(*Root)
	if !ok {
		return nil, errors.WrapError(errors.ErrTypeAssertion,
			fmt.Sprintf("expected type: %T, got: %T", root, ptrIface))
	}
	return rootPtr, nil
}

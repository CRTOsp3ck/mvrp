package models

import (
	"fmt"
	"mvrp/errors"
	"mvrp/util"
	"path/filepath"
)

type Model struct {
	Name       string `json:"name"`
	PluralName string `json:"plural_name"`
}

type View struct {
	Name       string `json:"name"`
	PluralName string `json:"plural_name"`
}

type Package struct {
	Package string  `json:"package"`
	Models  []Model `json:"models"`
	Views   []View  `json:"views,omitempty"`
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
		filepath.Join(rootDir, "config", "models", "models.json"), &root)
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

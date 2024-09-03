package handlers

import (
	"fmt"
	"mvrp/errors"
	"mvrp/util"
	"path/filepath"
)

type Route struct {
	Method  string `json:"method"`
	Path    string `json:"path"`
	Handler string `json:"handler"`
	Type    string `json:"type"`
}

type Handler struct {
	Name   string  `json:"name"`
	Routes []Route `json:"routes"`
}

type Package struct {
	Package  string    `json:"package"`
	Handlers []Handler `json:"handlers"`
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
		filepath.Join(rootDir, "config", "handlers", "handlers.json"), &root)
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

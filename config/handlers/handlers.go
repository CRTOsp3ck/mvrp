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
	IsView bool    `json:"is_view"`
}

type Package struct {
	Package   string    `json:"package"`
	BaseRoute string    `json:"base_route"`
	Handlers  []Handler `json:"handlers"`
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

	updateIsViewField(rootPtr)

	return rootPtr, nil
}

func updateIsViewField(config *Root) {
	for i, pkg := range config.Data {
		for j, handlerGroup := range pkg.Handlers {
			// check if the name ends with "_view"
			config.Data[i].Handlers[j].IsView = util.Util.Str.EndsWith(handlerGroup.Name, "_view")
		}
	}
}

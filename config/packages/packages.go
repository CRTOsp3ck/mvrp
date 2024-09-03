package packages

import (
	"fmt"
	"mvrp/errors"
	"mvrp/util"
	"path/filepath"
)

type Packages struct {
	Packages []string `json:"packages"`
}

func GetConfig() (*Packages, error) {
	rootDir, err := util.Util.FS.FindProjectRoot("go.mod")
	if err != nil {
		return nil, err
	}
	var pkgs Packages
	ptrIface, err := util.Util.Json.ParseJsonFile(
		filepath.Join(rootDir, "config", "packages", "packages.json"), &pkgs)
	if err != nil {
		return nil, err
	}
	pkgsPtr, ok := ptrIface.(*Packages)
	if !ok {
		return nil, errors.WrapError(errors.ErrTypeAssertion,
			fmt.Sprintf("expected type: %T, got: %T", pkgs, ptrIface))
	}
	return pkgsPtr, nil
}

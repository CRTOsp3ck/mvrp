package packages

import (
	"mvrp/util"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetConfig(t *testing.T) {
	pkgs, err := GetConfig()
	require.NoError(t, err)
	assert.NotNil(t, pkgs)
	assert.NotEmpty(t, pkgs.Packages)
	util.Util.Json.PrintJson(pkgs)
}

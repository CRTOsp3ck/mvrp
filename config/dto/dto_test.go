package dto

import (
	"mvrp/util"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetConfig(t *testing.T) {
	root, err := GetConfig()
	require.NoError(t, err)
	assert.NotNil(t, root)
	assert.NotEmpty(t, root.Data)
	util.Util.Json.PrintJson(root)
}

func TestGetDTOConfigs(t *testing.T) {
	dtos, err := GetDTOs()
	require.NoError(t, err)
	assert.NotEmpty(t, dtos)
	util.Util.Json.PrintJson(dtos)
}

func TestIfDtoExists(t *testing.T) {
	exists := IfDtoExists("SearchEntityDTO")
	assert.True(t, exists)
	exists = IfDtoExists("UserDTO1")
	assert.False(t, exists)
}

package models

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

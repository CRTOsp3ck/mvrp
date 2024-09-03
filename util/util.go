package util

import (
	"mvrp/util/ds"
	"mvrp/util/fs"
	"mvrp/util/json"
	"mvrp/util/nc"
	"mvrp/util/ptr"
	"mvrp/util/str"
)

var Util *UtilContainer

type UtilContainer struct {
	FS   fs.FSUtil
	DS   ds.DSUtil
	NC   nc.NCUtil
	Json json.JsonUtil
	Ptr  ptr.PtrUtil
	Str  str.StrUtil
}

func init() {
	Util = &UtilContainer{
		FS:   fs.FSUtil{},
		DS:   ds.DSUtil{},
		NC:   nc.NCUtil{},
		Json: json.JsonUtil{},
		Ptr:  ptr.PtrUtil{},
		Str:  str.StrUtil{},
	}
}

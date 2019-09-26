package gart

import (
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

var exports = map[string]interface{}{
	"Vector2d":          NewVector2d,
	"Vector2dFromPolar": NewVector2dFromPolar,
	"Palettes":          Palettes,
}

func luaLibLoad(L *lua.LState) int {
	// register functions to the table
	mod := L.NewTable()

	// register other stuff
	for name, value := range exports {
		L.SetField(mod, name, luar.New(L, value))
	}

	// returns the module
	L.Push(mod)

	return 1
}

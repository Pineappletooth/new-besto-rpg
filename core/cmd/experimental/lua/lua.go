package main

import (
	"github.com/yuin/gopher-lua"
	"layeh.com/gopher-luar"
)

type User struct {
	Name  string
	Token func(string) string
}

type Wrapper struct {
	UserW User
}

const script = `
local Name = "KK"
a.Token = function(c) return c..Name..1 end
`
const script2 = `
local Name = "GG"
a.Token = function(c) return c..Name..2 end
`

func Example_basic() {
	L := lua.NewState()
	defer L.Close()

	user := getUser(L, script)
	user2 := getUser(L, script2)
//	getUser(L, script2)
	if user2.Token != nil {
		println(user2.Token("xxx"))
	}
	println(user.Token("Hello from Lua, "))
	

	// Output:
	// Hello from Lua, Tim!
	// Lua set your token to: 12345
}

func getUser(L *lua.LState, ss string) User {
	wrapper := &User{}
	L.SetGlobal("a", luar.New(L, wrapper))
	if err := L.DoString(ss); err != nil {
		panic(err)
	}
	return *wrapper
}
func main() {
	Example_basic()
}

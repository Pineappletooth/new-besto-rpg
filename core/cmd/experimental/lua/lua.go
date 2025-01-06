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
return {
    Name = "10",
    Token = function(a) return a.."qwe" end
}
`
const script2 = `
return {
    Name = "10",
}
`
const loader =`
sss.UserW = aaa
`
func Example_basic() {
    L := lua.NewState()
    defer L.Close()


   user := getUser(L,script)
   user2 := getUser(L,script2)
   if user.Token != nil {
    println(user.Token("111"))

}
    if user2.Token != nil {
        println(user2.Token("444"))

    }

 
    // Output:
    // Hello from Lua, Tim!
    // Lua set your token to: 12345
}

func getUser(L *lua.LState,ss string) User {
    if err := L.DoString(ss); err != nil {
        panic(err)
    }
    wrapper := &Wrapper{
    }
    L.SetGlobal("aaa", L.Get(-1))
    L.Pop(1)
    L.SetGlobal("sss",luar.New(L,wrapper))
    if err := L.DoString(loader); err != nil {
        panic(err)
    }
    user := wrapper.UserW
    return user
}
func main() {
    Example_basic()
}
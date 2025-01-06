package main

import (
	"fmt"
	"time"

	"github.com/yuin/gopher-lua"
	"layeh.com/gopher-luar"
)

type Fib struct {
    Token func(int) int
}
const script0 = `
function fib0(n, a, b)
   if n == 0 then
		return a
	elseif n == 1 then
		return b
	end
	return fib0(n-1, b, a+b)
end
`

const script = `
function fib(n)
    if n < 2 then return n end
    return fib(n - 2) + fib(n - 1)
end
`

const script2 = `
uu.Token = function(n)
    if n < 2 then return n end
    return fib(n - 2) + fib(n - 1)
end
`

const script3 = `
function fib3(n)
    uu.Token(n)
end
`

const loader =`
sss.UserW = aaa
`
func Example_basic() {
    L := lua.NewState()
    defer L.Close()

    if err := L.DoString(script0); err != nil {
        panic(err)
    }
    
    for i := 0; i < 5; i++ {
        ts := time.Now()
        if err := L.CallByParam(lua.P{
            Fn: L.GetGlobal("fib0"),
            NRet: 1,
            Protect: true,
            }, lua.LNumber(35),lua.LNumber(0),lua.LNumber(1)); err != nil {
                panic(err)
            }
        fmt.Println("cost:", time.Since(ts))
    }


    if err := L.DoString(script); err != nil {
        panic(err)
    }
    
    for i := 0; i < 0; i++ {
        ts := time.Now()
        if err := L.CallByParam(lua.P{
            Fn: L.GetGlobal("fib"),
            NRet: 1,
            Protect: true,
            }, lua.LNumber(35)); err != nil {
                panic(err)
            }
        fmt.Println("cost:", time.Since(ts))
    }


    fib := Fib{}
    L.SetGlobal("uu",luar.New(L,&fib))
    if err := L.DoString(script2); err != nil {
        panic(err)
    }
    for i := 0; i < 0; i++ {
        ts := time.Now()
        fib.Token(35)
        fmt.Println("cost:", time.Since(ts))
    }
    if err := L.DoString(script3); err != nil {
        panic(err)
    }

    for i := 0; i < 0; i++ {
        ts := time.Now()
        if err := L.CallByParam(lua.P{
            Fn: L.GetGlobal("fib3"),
            NRet: 1,
            Protect: true,
            }, lua.LNumber(35)); err != nil {
                panic(err)
            }
        fmt.Println("cost:", time.Since(ts))
    }
    fib4 := Fib{}
    fib4.Token = func(n int) int {
            if n < 2 {
                return n
            }
            return fib4.Token(n - 2) + fib4.Token(n - 1)
        }
    
    L.SetGlobal("uu",luar.New(L,&fib4))
    for i := 0; i <4; i++ {
        ts := time.Now()
        if err := L.CallByParam(lua.P{
            Fn: L.GetGlobal("fib3"),
            NRet: 1,
            Protect: true,
            }, lua.LNumber(35)); err != nil {
                panic(err)
            }
        fmt.Println("cost:", time.Since(ts))
    }



    for i := 0; i < 4; i++ {
        ts := time.Now()
        fib4.Token(35)
        fmt.Println("cost:", time.Since(ts))
    }

    for i := 0; i < 0; i++ {
        ts := time.Now()
        a := fibbb(35)
        print(a)
        fmt.Println("cost:", time.Since(ts))
    }


}

func fibbb(n int) int {
    if n < 2 {
        return n
    }
    return fibbb(n - 2) + fibbb(n - 1)
}

func main() {
    Example_basic()
}
package main

import "fmt"

func main() {
    a := 20
    f := func() {
        fmt.Println(a)
        a = 30
    }
    f()            // 20
    fmt.Println(a) // 30
}

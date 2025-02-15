package main

import "fmt"

func main() {
    x := make([]string, 0, 5)
    x = append(x, "a", "b", "c", "d")
    y := x[:2:2]
    z := x[2:4:4]
    fmt.Println(cap(y)) // 2 - 0 = 2
    fmt.Println(cap(z)) // 4 - 2 = 2
}

package main

import "fmt"

var (
    add = func(i, j int) int { return i + j }
    sub = func(i, j int) int { return i - j }
    mul = func(i, j int) int { return i * j }
    div = func(i, j int) int { return i / j }
)

func changeAdd() {
    add = func(i, j int) int { return i + j + j }
}

func main() {
    x := add(2, 3)
    fmt.Println(x) // 5
    changeAdd()
    y := add(2, 3)
    fmt.Println(y) // 8
}

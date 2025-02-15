package main

import "fmt"

func main() {
    x := []int{1, 2, 3, 4}
    y := make([]int, 4)
    num := copy(y, x)
    fmt.Println(y, num) // [1 2 3 4] 4

    m := []int{1, 2, 3, 4}
    n := make([]int, 3)
    num = copy(n[:3], m[1:])
    fmt.Println(n, num) // [2 3 4] 3
}

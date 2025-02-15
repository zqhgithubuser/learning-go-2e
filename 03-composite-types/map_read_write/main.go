package main

import "fmt"

func main() {
    totalWins := map[string]int{}
    totalWins["Orcas"] = 1
    totalWins["Lions"] = 2
    fmt.Println(totalWins["Orcas"])   // 1
    fmt.Println(totalWins["Kittens"]) // 0

    totalWins["Kittens"]++
    fmt.Println(totalWins["Kittens"]) // 1

    totalWins["Lions"] = 3
    fmt.Println(totalWins["Lions"]) // 3

    fmt.Println(totalWins) // map[Kittens:1 Lions:3 Orcas:1]
}

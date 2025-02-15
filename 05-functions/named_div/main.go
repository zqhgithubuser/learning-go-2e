package main

import (
    "errors"
    "fmt"
)

func divAndRemainder(num, denom int) (result int, remainder int, err error) {
    if denom == 0 {
        err = errors.New("cannot divide by zero")
        return result, remainder, err
    }
    return num / denom, num % denom, nil
}

func main() {
    x, y, z := divAndRemainder(5, 2)
    fmt.Println(x, y, z)
}

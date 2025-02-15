package main

import "fmt"

type MyInt int

func typeAssert() {
    var i any
    var mine MyInt = 20
    i = mine
    i2 := i.(MyInt)
    fmt.Println(i2 + 1)
}

func typeAssertPanicWrongType() {
    defer func() {
        if m := recover(); m != nil {
            fmt.Println(m) // prints out because a panic happens
        }
    }()
    var i any
    var mine MyInt = 20
    i = mine
    i2 := i.(string)
    fmt.Println(i2)
}

func typeAssertPanicTypeNotIdentical() {
    defer func() {
        if m := recover(); m != nil {
            fmt.Println(m) // prints out because a panic happens
        }
    }()
    var i any
    var mine MyInt = 20
    i = mine
    i2 := i.(int)
    fmt.Println(i2 + 1)
}

func typeAssertCommaOK() error {
    var i any
    var mine MyInt = 20
    i = mine
    i2, ok := i.(int)
    if !ok {
        return fmt.Errorf("unexpected type for %v", i)
    }
    fmt.Println(i2 + 1)
    return nil
}

func main() {
    typeAssert()                      // 21
    typeAssertPanicWrongType()        // interface conversion: interface {} is main.MyInt, not string
    typeAssertPanicTypeNotIdentical() // interface conversion: interface {} is main.MyInt, not int
    err := typeAssertCommaOK()
    if err != nil {
        fmt.Println(err) // unexpected type for 20
    }
}

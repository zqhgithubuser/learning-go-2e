package main

import (
    "fmt"
    "time"
)

type Counter struct {
    total       int
    lastUpdated time.Time
}

func (c *Counter) Increment() {
    c.total++
    c.lastUpdated = time.Now()
}

func (c Counter) String() string {
    return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

type Incrementer interface {
    Increment()
}

func main() {
    var myStringer fmt.Stringer
    var myIncrementer Incrementer
    pointerCounter := &Counter{}
    valueCounter := Counter{}

    // Counter 实现了 String 方法
    myStringer = pointerCounter
    myStringer = valueCounter
    // Counter 实现了 Increment 方法
    myIncrementer = pointerCounter
    //myIncrementer = valueCounter // compile-time error!

    fmt.Println(myStringer, myIncrementer)
}

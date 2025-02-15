package main

import "fmt"

func main() {
	a := 20
	f := func() {
		fmt.Println(a) // 20
		a := 30
		fmt.Println(a) // 30
	}
	f()
	fmt.Println(a) // 20
}

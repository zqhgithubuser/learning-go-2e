package main

import "fmt"

func main() {
	// Literals
	fmt.Println(1_234)   // 1234
	fmt.Println(6.03e23) // 6.03e+23
	fmt.Printf("%T %T %T %T %T\n", 'a', '\141', '\x61', '\u0061', '\U00000061')
	// int32 int32 int32 int32 int32
	fmt.Printf("%T\n", "Greetings and Salutations") // string

	// bool
	var flag bool
	var isAwesome = true
	fmt.Println(flag, isAwesome) // false true

	// integer
	var x = 10
	fmt.Println(x * 2) // 20

	// float
	var f float64 = 10
	var z float64 = 0
	fmt.Println(f / z) // +Inf
	fmt.Println(z / z) // NaN

	// rune
	var myFirstInitial rune = 'J'
	var myLastInitial int32 = 'B'
	fmt.Printf("%T %T\n", myFirstInitial, myLastInitial) // int32 int32
}

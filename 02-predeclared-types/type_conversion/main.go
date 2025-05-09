package main

import "fmt"

func main() {
	var x = 10
	var y = 30.2
	var sum1 = float64(x) + y
	var sum2 = x + int(y)
	fmt.Println(sum1, sum2) // 40.2 40

	var b byte = 100
	var sum3 = x + int(b)
	var sum4 = byte(x) + b
	fmt.Println(sum3, sum4) // 110 110
}

package main

import "fmt"

func failedUpdate(px *int) {
	x2 := 20
	px = &x2 // px 是指针的拷贝
}

func update(px *int) {
	*px = 20
}

func main() {
	x := 10
	failedUpdate(&x)
	fmt.Println(x) // 10
	update(&x)
	fmt.Println(x) // 20
}

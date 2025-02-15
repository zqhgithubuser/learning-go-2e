package main

import "fmt"

func main() {
	var s = "Hello, 😺"
	fmt.Println(len(s)) // 11
	var bs = []byte(s)
	var rs = []rune(s)
	fmt.Println(bs) // [72 101 108 108 111 44 32 240 159 152 186]
	fmt.Println(rs) // [72 101 108 108 111 44 32 128570]
}

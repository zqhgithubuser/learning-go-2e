package main

import "fmt"

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func Convert[T1, T2 Integer](in T1) T2 {
	return T2(in)
}

func main() {
	var a int = 10
	// 显式指定类型
	b := Convert[int, int64](a)
	fmt.Println(b) // 10
}

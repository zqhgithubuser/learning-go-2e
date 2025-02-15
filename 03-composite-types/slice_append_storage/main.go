package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	fmt.Println(cap(x), cap(y)) // 4 4
	y = append(y, "z")
	fmt.Println("x:", x) // [a b z d]
	fmt.Println("y:", y) // [a b z]
	fmt.Println("----------------------------------------")

	m := []string{"a", "b", "c", "d"}
	n := m[2:]                  // cap(n) = cap(m) - 2 = 2
	fmt.Println(cap(m), cap(n)) // 4 2
	n = append(n, "z")
	fmt.Println(cap(m), cap(n)) // 4 4
	fmt.Println("m:", m)        // m: [a b c d]
	fmt.Println("n:", n)        // n: [c d z]
}

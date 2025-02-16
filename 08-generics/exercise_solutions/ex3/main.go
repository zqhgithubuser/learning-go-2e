package main

import "fmt"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

// Add adds a new element to the end of the linked list
func (l *List[T]) Add(t T) {
	n := &Node[T]{
		Value: t,
	}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = l.Tail.Next
}

// Insert adds an element at the specified position in the linked list
func (l *List[T]) Insert(t T, pos int) {
	n := &Node[T]{
		Value: t,
	}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	if pos <= 0 {
		n.Next = l.Head
		l.Head = n
		return
	}
	curNode := l.Head
	for i := 1; i < pos; i++ {
		// len(list) < pos
		if curNode.Next == nil {
			curNode.Next = n
			l.Tail = curNode.Next
			return
		}
		curNode = curNode.Next
	}
	n.Next = curNode.Next
	curNode.Next = n
	if l.Tail == curNode {
		l.Tail = n
	}
}

// Index returns the position of the supplied value, -1 if it's not present
func (l *List[T]) Index(t T) int {
	i := 0
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		if curNode.Value == t {
			return i
		}
		i++
	}
	return -1
}

func main() {
	l := &List[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))  // 0
	fmt.Println(l.Index(10)) // 1
	fmt.Println(l.Index(20)) // -1

	l.Insert(100, 0)
	fmt.Println(l.Index(5))   // 1
	fmt.Println(l.Index(10))  // 2
	fmt.Println(l.Index(20))  // -1
	fmt.Println(l.Index(100)) // 0

	l.Insert(200, 1)
	fmt.Println(l.Index(5))   // 2
	fmt.Println(l.Index(10))  // 3
	fmt.Println(l.Index(200)) // 1
	fmt.Println(l.Index(20))  // -1
	fmt.Println(l.Index(100)) // 0

	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Printf("%d ", curNode.Value) // 100 200 5 10
	}
	fmt.Println()

	l.Insert(300, 10)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Printf("%d ", curNode.Value) // 100 200 5 10 300
	}
	fmt.Println()

	l.Add(400)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Printf("%d ", curNode.Value) // 100 200 5 10 300 400
	}
	fmt.Println()

	l.Insert(500, 6)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Printf("%d ", curNode.Value) // 100 200 5 10 300 400 500
	}
	fmt.Println()
}

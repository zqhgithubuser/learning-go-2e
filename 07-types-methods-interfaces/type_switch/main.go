package main

import (
	"errors"
	"fmt"
)

type treeNode struct {
	val    treeVal
	lchild *treeNode
	rchild *treeNode
}

type treeVal interface {
	isToken()
}

type number int

func (number) isToken() {}

type operator func(int, int) int

func (operator) isToken() {}

func (o operator) process(n1, n2 int) int {
	return o(n1, n2)
}

var operators = map[string]operator{
	"+": func(n1, n2 int) int {
		return n1 + n2
	},
	"-": func(n1, n2 int) int {
		return n1 - n2
	},
	"*": func(n1, n2 int) int {
		return n1 * n2
	},
	"/": func(n1, n2 int) int {
		return n1 / n2
	},
}

func walkTree(t *treeNode) (int, error) {
	switch val := t.val.(type) {
	case nil:
		return 0, errors.New("invalid expression")
	case number:
		return int(val), nil
	case operator:
		left, err := walkTree(t.lchild)
		if err != nil {
			return 0, err
		}
		right, err := walkTree(t.rchild)
		if err != nil {
			return 0, err
		}
		return val.process(left, right), nil
	default:
		return 0, errors.New("unknown node type")
	}
}

func parse(s string) (*treeNode, error) {
	return &treeNode{
		val: operators["+"],
		lchild: &treeNode{
			val:    operators["*"],
			lchild: &treeNode{val: number(5)},
			rchild: &treeNode{val: number(10)},
		},
		rchild: &treeNode{val: number(20)},
	}, nil
}

func main() {
	parseTree, err := parse("5*10+20")
	if err != nil {
		panic(err)
	}
	result, err := walkTree(parseTree)
	fmt.Println(result, err)
}

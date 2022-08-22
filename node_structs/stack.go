package main

import "fmt"

type node struct {
	val  int
	next *node
}

type Stack struct {
	head *node
}

func newnode(val int) *node {
	return &node{val: val}
}

func newStack() *Stack {
	return &Stack{}
}

func (stack *Stack) push(item int) {
	currentHead := stack.head
	toPush := newnode(item)

	toPush.next = currentHead
	stack.head = toPush
}

func (stack *Stack) pop() (int, bool) {
	currentHead := stack.head

	if currentHead == nil {
		return 0, false
	}

	nextnode := currentHead.next
	stack.head = nextnode

	return currentHead.val, true
}

func main() {
	s := newStack()
	s.push(3)
	s.push(7)
	s.push(8)

	s.pop()
	fmt.Println(s.head.val)
}

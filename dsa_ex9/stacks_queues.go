package main

import (
	"fmt"
	"strings"
)

// Queue implementation
type Queue struct {
	store []string
}

func (s *Queue) deque(str string) string {
	item := s.store[0]
	s.store = s.store[1:]

	return item
}

func (s *Queue) enque(str string) int {
	s.store = append(s.store, str)

	return len(s.store)
}

func (s *Queue) initialize(list []string) {
	s.store = append(s.store, list...)
}

// Stack implementation
type Stack struct {
	store []string
}

func (s *Stack) push(item string) {
	s.store = append(s.store, item)
}

func (s *Stack) pop() string {
	lastIdx := len(s.store) - 1
	popped := s.store[lastIdx]

	s.store = s.store[:lastIdx]

	return popped
}

func reverse(str string) string {
	chars := &Stack{}

	for _, e := range str {
		chars.push(string(e))
	}

	builder := strings.Builder{}

	for i := 0; i < len(str); i++ {
		builder.WriteString(chars.pop())
	}

	return builder.String()
}

func main() {
	fmt.Println(reverse("hello"))
	fmt.Println("vim-go")
}

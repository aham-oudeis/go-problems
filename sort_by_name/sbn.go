package main

import (
	"fmt"
	"sort"
)

var Conversion map[int]string = map[int]string{
	0: "zero",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

type Number string
type byName []int

func (s byName) Len() int {
	return len(s)
}

func (s byName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byName) Less(i, j int) bool {
	return Conversion[s[i]] < Conversion[s[j]]
}

func main() {
	col := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(Number("Hello"))
	fmt.Println(byName(col))
	sort.Sort(byName(col))
	fmt.Println(col)
}

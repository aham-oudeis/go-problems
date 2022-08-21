package main

import "fmt"

func simpleMap(list []int) map[int]bool {
	seen := map[int]bool{}
	for _, e := range list {
		if _, ok := seen[e]; !ok {
			seen[e] = true
		}
	}

	return seen
}

func union(l1, l2 []int) []int {
	seen := simpleMap(l1)
	all := l1

	for _, item := range l2 {
		_, ok := seen[item]

		if !ok {
			all = append(all, item)
		}
	}

	return all
}
func uniqItems(list []int) []int {
	seen := map[int]bool{}
	coll := []int{}

	for _, item := range list {
		if _, ok := seen[item]; !ok {
			coll = append(coll, item)
		}

		seen[item] = true
	}

	return coll
}

func main() {
	l1 := []int{1, 3, 4, 5}
	l2 := []int{2, 4, 8, 9}
	l3 := []int{2, 4, 7, 2, 8, 4, 9}

	fmt.Println(union(l1, l2))
	fmt.Println(uniqItems(l3))
}

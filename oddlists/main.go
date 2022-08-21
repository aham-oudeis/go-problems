package main

import "fmt"

func predicate(_, i int) bool {
	return i%2 == 0
}

func filter(list []int, predicate func(val, idx int) bool) []int {
	collection := []int{}

	for i, e := range list {
		if predicate(list[i], i) {
			collection = append(collection, e)
		}
	}

	return collection
}

func main() {
	list := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(filter(list, predicate))
}

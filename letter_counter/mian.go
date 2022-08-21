package main

import (
	"fmt"
	"strings"
)

func transform(list []string, change func(val string) int) []int {
	collection := make([]int, len(list))

	for i := range list {
		collection[i] = change(list[i])
	}

	return collection
}

func countMap(iterable []int) map[int]int {
	count := map[int]int{}

	for i := range iterable {
		e := iterable[i]

		if value, ok := count[e]; ok == true {
			count[e] = value + 1
		} else {
			count[e] = 1
		}
	}

	return count
}

func wordSizes(str string) map[int]int {
	length := func(word string) int {
		return len(word)
	}

	collection := transform(strings.Split(str, " "), length)

	return countMap(collection)

}
func main() {
	fmt.Println(wordSizes("Four score and seven"))
}

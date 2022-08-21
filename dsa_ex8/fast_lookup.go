package main

import (
	"fmt"
	"strings"
)

type Key interface {
	int | string
}

func simpleMap[K Key, V bool](list []K) map[K]V {
	seen := map[K]V{}
	for _, e := range list {
		if _, ok := seen[e]; !ok {
			seen[e] = true
		}
	}

	return seen
}

func countMap[K Key, V int](list []K) map[K]V {
	seen := map[K]V{}

	for _, e := range list {
		if _, ok := seen[e]; ok {
			seen[e] += 1
		} else {
			seen[e] = 1
		}
	}

	return seen
}

func common(l1, l2 []int) []int {
	seen := simpleMap(l1)

	col := []int{}

	for _, num := range l2 {
		if _, ok := seen[num]; ok == true {
			col = append(col, num)
		}
	}

	return col
}

func firstDupliate(chars []string) string {
	seen := countMap(chars)

	fmt.Println(seen)
	for _, item := range chars {
		val, ok := seen[item]

		if ok == true && val == 2 {
			return item
		}
	}

	return ""
}

func missingLetter(str string) string {
	alphabets := "abcdefghijklmnopqrstuvqxyz"

	charsList := strings.Split(str, "")
	seen := simpleMap(charsList)

	for idx := range alphabets {
		char := string(alphabets[idx])

		if _, ok := seen[char]; !ok {
			return char
		}
	}

	return ""
}

func findNonDuplicate(str string) string {
	chars := strings.Split(str, "")
	count := countMap(chars)

	for _, char := range chars {
		val, _ := count[char]

		if val == 1 {
			return char
		}
	}

	return ""
}

func main() {
	l1 := []int{1, 2, 3, 4, 5, 6, 7}
	l2 := []int{11, 2, 3, 44, 8, 6, 7}

	strs := []string{"a", "b", "c", "e", "d", "c", "b", "a"}
	fmt.Println(firstDupliate(strs))

	chars := "abcdefghijklmnopqrstuwxyz"

	fmt.Println("The missing letter is: ")
	fmt.Println(missingLetter(chars))

	fmt.Println(common(l1, l2))
	fmt.Println("vim-go")

	fmt.Println(findNonDuplicate(strings.Join(strs, "")))
}

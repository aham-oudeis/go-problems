package main

import "fmt"

func multisum(ceil int) int {
	var total int = 0
	var col []int

	var i int
	for i = 1; i < ceil; i++ {
		if i%3 == 0 || i%5 == 0 {
			col = append(col, i)
			total += i
		}
	}

	fmt.Println(col)
	return total
}

func main() {
	fmt.Println(multisum(3))
	fmt.Println(multisum(9))
	fmt.Println(multisum(13))
	fmt.Println(multisum(20))
	fmt.Println(multisum(148))
}

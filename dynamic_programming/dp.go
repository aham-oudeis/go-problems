package main

import "fmt"

func addUntill100(list []int) int {
	if len(list) == 0 {
		return 0
	}

	sum := list[0] + addUntill100(list[1:])

	if sum > 100 {
		return sum - list[0]
	} else {
		return sum
	}
}

func addTillHundred(list []int) int {
	var total int

	var limitSum func([]int, int) int

	limitSum = func(arr []int, total int) int {
		if len(arr) == 0 {
			return total
		}

		sum := arr[0] + total

		if sum > 100 {
			return limitSum(arr[1:], total)
		} else {
			return limitSum(arr[1:], sum)
		}
	}

	return limitSum(list, total)
}

func golomb(n int) int {
	cache := map[int]int{}
	cache[1] = 1

	var g func(int) int

	g = func(num int) int {
		value, ok := cache[num]

		if !ok {
			value = 1 + g(num-g(g(num-1)))
			cache[num] = value
		}

		return value
	}

	return g(n)
}

func main() {
	l := []int{1, 4, 5, 22, 21, 4, 100, 8, 7}
	fmt.Println(addUntill100(l))
	fmt.Println(addTillHundred(l))

	fmt.Println(golomb(2))
	fmt.Println("vim-go")
}

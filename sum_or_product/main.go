package main

import "fmt"

func sum(n1, n2 float64) float64 {
	total := float64(0)

	for n1 <= n2 {
		total += n1
		n1++
	}

	return total
}

func product(n1, n2 float64) float64 {
	total := float64(1)

	for n1 <= n2 {
		total *= n1
	}

	return total
}

func main() {
	var n1 int
	var oper string

	fmt.Println("Enter a number greater than 0: ")
	fmt.Scanln(&n1)

	fmt.Println("Enter 's' to sum or 'm' to multiply all natural numbers upto ", n1)
	fmt.Scanln(&oper)

	fmt.Println(n1, oper)

	if n1 < 1 {
		panic("The number is not valid")
	}

	if oper == string('s') {
		fmt.Println("The sum of numbers upto", n1, "is")
		fmt.Println(sum(1, float64(n1)))
		fmt.Println()
	}

	if oper == string('m') {
		fmt.Println(product(1, float64(n1)))
	}

}

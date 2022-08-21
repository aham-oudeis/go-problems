package main

import "fmt"

func isLeapYear(year int) bool {
	var zero int
	var julianCalendarBegin int = 1752

	if year < julianCalendarBegin {
		return year%4 == 0
	}

	switch zero {
	case year % 400:
		return true
	case year % 100:
		return false
	case year % 4:
		return true
	}

	return false
}

func main() {
	fmt.Println(isLeapYear(2016), true)    // true
	fmt.Println(isLeapYear(2015), false)   // false
	fmt.Println(isLeapYear(2100), false)   // false
	fmt.Println(isLeapYear(2400), true)    // true
	fmt.Println(isLeapYear(240000), true)  // true
	fmt.Println(isLeapYear(240001), false) // false
	fmt.Println(isLeapYear(2000), true)    // true
	fmt.Println(isLeapYear(1900), false)   // false
	fmt.Println(isLeapYear(1752))          // true
	fmt.Println(isLeapYear(1700))          // false
	fmt.Println(isLeapYear(1))             // false
	fmt.Println(isLeapYear(100))           // false
	fmt.Println(isLeapYear(400))           // true
}

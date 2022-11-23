package main

import "fmt"

func main() {
	if IsLeapYear(1996) {
		fmt.Println("1996 is a leap year")
	} else {
		fmt.Println("1996 is not a leap year")
	}
	if IsLeapYear(2000) {
		fmt.Println("2000 is a leap year")
	} else {
		fmt.Println("2000 is not a leap year")
	}
	if IsLeapYear(1900) {
		fmt.Println("1900 is a leap year")
	} else {
		fmt.Println("1900 is not a leap year")
	}
	if IsLeapYear(1997) {
		fmt.Println("1997 is a leap year")
	} else {
		fmt.Println("1997 is not a leap year")
	}
}

func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
		// if year%4 == 0 {
		// if b_year100 == 0 {
		// 	if b_year400 != 0 {
		// 		return false
		// 	}
		// 	return true
		// }
		// return true
	}
	return false
}

// https://theweeklychallenge.org/blog/perl-weekly-challenge-191/
// You are given a string with alphabetic characters only: A..Z and a..z.

// Write a script to find out if the usage of Capital is appropriate if it satisfies at least one of the following rules:

// 1) Only first letter is capital and all others are small.
// 2) Every letter is small.
// 3) Every letter is capital.
package main

import (
	"fmt"
	"unicode"
)

func main() {
	input := "Perl" // return true
	fmt.Printf("%s:%t\n", input, isValid(input))

	input = "TPF" // return true
	fmt.Printf("%s:%t\n", input, isValid(input))

	input = "PyThon" // return false
	fmt.Printf("%s:%t\n", input, isValid(input))

	input = "raku" // return true
	fmt.Printf("%s:%t\n", input, isValid(input))

	input = "rAku" // return true
	fmt.Printf("%s:%t\n", input, isValid(input))

	// test switch statement
	input2 := "Perl" // return true
	fmt.Printf("%s:%t\n", input2, isValid(input2))

	input2 = "TPF" // return true
	fmt.Printf("%s:%t\n", input2, isValid(input2))

	input2 = "PyThon" // return false
	fmt.Printf("%s:%t\n", input2, isValid(input2))

	input2 = "raku" // return true
	fmt.Printf("%s:%t\n", input2, isValid(input2))

	input2 = "rAku" // return true
	fmt.Printf("%s:%t\n", input2, isValid(input2))
}

// Package acronym implements a utility for creating an acronym from
// a given phrase
func isValid(word string) bool {
	var isUpper bool
	var allUpper bool
	var allLower bool
	// var lastLetter rune
	isUpper = false
	for i, j := range word {
		// check next character is a letter
		if !unicode.IsLetter(j) {
			return false
		}

		// is first letter upper or lower
		if i == 0 {
			if unicode.IsUpper(j) {
				isUpper = true
			}
			continue
		}

		// if first letter is lower and second is upper then return false
		if i == 1 && !isUpper && unicode.IsUpper(j) {
			return false
		}

		// if first and second letters are upper then all letters are be upper
		if i == 1 && isUpper && unicode.IsUpper(j) {
			allUpper = true
		}

		// if first and second letters are lower then all letters are be lower
		if i == 1 && !isUpper && unicode.IsLower(j) {
			allLower = true
		}

		// tests perform on letter afters the second....:
		// 1. if character is lower but should be upper return false
		if i > 1 && unicode.IsLower(j) && allUpper {
			return false
		}
		// 2. if character is upper but should be lower return false
		if i > 1 && unicode.IsUpper(j) && allLower {
			return false
		}
		// 3. if character is upper but not all characters are not upper then
		// return false
		if i > 1 && unicode.IsUpper(j) && !allUpper {
			return false
		}
	}
	return true
}

// Package acronym implements a utility for creating an acronym from
// a given phrase
func isValidSwitch(word string) bool {
	var isUpper bool
	var allUpper bool
	var allLower bool
	// var lastLetter rune
	isUpper = false
	for i, j := range word {
		switch {
		case i == 0: // first letter
			// check next character is a letter
			if !unicode.IsLetter(j) {
				return false
			}

			// is first letter upper or lower
			if unicode.IsUpper(j) {
				isUpper = true
			}

		case i == 1: // second letter
			// if first letter is lower and second is upper then return false
			if !isUpper && unicode.IsUpper(j) {
				return false
			}

			// if first and second letters are upper then all letters are be upper
			if isUpper && unicode.IsUpper(j) {
				allUpper = true
			}

			// if first and second letters are lower then all letters are be lower
			if !isUpper && unicode.IsLower(j) {
				allLower = true
			}

		case i > 1: // rest of the letters
			// tests perform on letter afters the second....:
			// 1. if character is lower but should be upper return false
			if unicode.IsLower(j) && allUpper {
				return false
			}
			// 2. if character is upper but should be lower return false
			if unicode.IsUpper(j) && allLower {
				return false
			}
			// 3. if character is upper but not all characters are not upper then
			// return false
			if unicode.IsUpper(j) && !allUpper {
				return false
			}
		}
	}
	return true
}

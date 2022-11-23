// https://theweeklychallenge.org/blog/perl-weekly-challenge-191/

package main

import (
	"fmt"
	"sort"
)

func main() {
	aSlice := []int{1, 2, 3, 4}
	if is_larget_double(aSlice) {
		fmt.Println("%+q", aSlice)
	}

	aSlice = []int{1, 2, 0, 5}
	if is_larget_double(aSlice) {
		fmt.Println("%+q", aSlice) // TODO: what is this??
	}

	aSlice = []int{2, 6, 3, 1}
	if is_larget_double(aSlice) { // TODO: what what
		fmt.Println("%+q", aSlice)
	}

	aSlice = []int{4, 5, 2, 3}
	if is_larget_double(aSlice) {
		fmt.Println("%+q", aSlice)
	}
}

// Package acronym implements a utility for creating an acronym from
// a given phrase
func is_larget_double(aSlice []int) bool {
	// Write some code here to pass the test suite.
	sort.Slice(aSlice, func(i, j int) bool {
		return aSlice[i] > aSlice[j] // sort descending
	})
	highest := 0
	for _, j := range aSlice {
		if highest == 0 {
			highest = j
			continue
		}
		if j*2 > highest {
			return false
		}
	}
	return true
}

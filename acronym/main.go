package main

import (
	"fmt"
	"strings"
)

// import (
// 	"github.com/funkynunu/go_rest_api/pkg/leapyear"
// )

func main() {
	fmt.Println(Abbreviate("Testing One Two"))
	fmt.Println(Abbreviate(""))
	fmt.Println(Abbreviate("Portable Network Graphics"))
	fmt.Println(Abbreviate("Ruby on Rails"))
	fmt.Println(Abbreviate("First In, First Out"))
	fmt.Println(Abbreviate("GNU Image Manipulation Program"))
	fmt.Println(Abbreviate("Complementary metal-oxide semiconductor"))
	fmt.Println(Abbreviate("Rolling On The Floor Laughing So Hard That My Dogs Came Over And Licked Me"))
	fmt.Println(Abbreviate("Something - I made up from thin air"))
	fmt.Println(Abbreviate("Halley's Comet"))
	fmt.Println(Abbreviate("The Road _Not_ Taken"))
}

// Package acronym implements a utility for creating an acronym from
// a given phrase
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	s = strings.ToUpper(s)
	var anagram string
	var first string

	for _, l := range s {
		if first == "" && string(l) != " " && string(l) != "-" && string(l) != "_" {
			first = string(l)
		}

		if string(l) == " " || string(l) == "-" {
			anagram += string(first)
			first = ""
		}
	}
	if first != "" {
		anagram += string(first)
	}

	return anagram
}

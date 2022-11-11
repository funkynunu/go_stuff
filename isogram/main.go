package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	if IsIsogram("lumberjacks") {
		fmt.Println("lumberjacks is a isogram")
	} else {
		fmt.Println("lumberjacks is not a isogram")
	}
	if IsIsogram("background") {
		fmt.Println("background is a isogram")
	} else {
		fmt.Println("background is not a isogram")
	}
	if IsIsogram("downstream") {
		fmt.Println("downstream is a isogram")
	} else {
		fmt.Println("downstream is not a isogram")
	}
	if IsIsogram("six-year-old") {
		fmt.Println("six-year-old is a isogram")
	} else {
		fmt.Println("six-year-old is not a isogram")
	}
	if IsIsogram("isograms") {
		fmt.Println("isograms is a isogram")
	} else {
		fmt.Println("isograms is not a isogram")
	}
	if IsIsogram("") {
		fmt.Println(" is a isogram")
	} else {
		fmt.Println(" is not a isogram")
	}
	if IsIsogram("eleven") {
		fmt.Println("eleven is a isogram")
	} else {
		fmt.Println("eleven is not a isogram")
	}
	if IsIsogram("zzyzx") {
		fmt.Println("zzyzx is a isogram")
	} else {
		fmt.Println("zzyzx is not a isogram")
	}
	if IsIsogram("subdermatoglyphic") {
		fmt.Println("subdermatoglyphic is a isogram")
	} else {
		fmt.Println("subdermatoglyphic is not a isogram")
	}
	if IsIsogram("Alphabet") {
		fmt.Println("Alphabet is a isogram")
	} else {
		fmt.Println("Alphabet is not a isogram")
	}
	if IsIsogram("alphAbet") {
		fmt.Println("alphAbet is a isogram")
	} else {
		fmt.Println("alphAbet is not a isogram")
	}
	if IsIsogram("thumbscrew-japingly") {
		fmt.Println("thumbscrew-japingly is a isogram")
	} else {
		fmt.Println("thumbscrew-japingly is not a isogram")
	}
	if IsIsogram("thumbscrew-jappingly") {
		fmt.Println("thumbscrew-jappingly is a isogram")
	} else {
		fmt.Println("thumbscrew-jappingly is not a isogram")
	}
	if IsIsogram("Emily Jung Schwartzkopf") {
		fmt.Println("Emily Jung Schwartzkopf is a isogram")
	} else {
		fmt.Println("Emily Jung Schwartzkopf is not a isogram")
	}
	if IsIsogram("accentor") {
		fmt.Println("accentor is a isogram")
	} else {
		fmt.Println("accentor is not a isogram")
	}
	if IsIsogram("angola") {
		fmt.Println("angola is a isogram")
	} else {
		fmt.Println("angola is not a isogram")
	}
}

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	aSlice := []string{}
	for _, c := range word {
		if unicode.IsLetter(c) == false {
			continue
		}
		for _, d := range aSlice {
			if d == string(c) {
				return false
			}
		}
		aSlice = append(aSlice, string(c))
	}
	return true
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	// {"a", 1},                           // lowercase letter
	// {"A", 1},                           // uppercase letter
	// {"f", 4},                           // valuable letter
	// {"at", 2},                          // short word
	// {"zoo", 12},                        // short, valuable word
	// {"street", 6},                      // medium word
	// {"quirky", 22},                     // medium, valuable word
	// {"OxyphenButazone", 41},            // long, mixed-case word
	// {"pinata", 8},                      // english-like word
	// {"", 0},                            // empty input
	// {"abcdefghijklmnopqrstuvwxyz", 87}, // entire alphabet available

	fmt.Println(Score("a"))
	fmt.Println(Score("A"))
	fmt.Println(Score("f"))
	fmt.Println(Score("at"))
	fmt.Println(Score("zoo"))
	fmt.Println(Score("street"))
	fmt.Println(Score("quirky"))
	fmt.Println(Score("OxyphenButazone"))
	fmt.Println(Score("pinata"))
	fmt.Println(Score(""))
	fmt.Println(Score("abcdefghijklmnopqrstuvwxyz"))
}

func Score(word string) int {
	word = strings.ToUpper(word)
	m := make(map[int]string)
	m[1] = "AEIOULNRST"
	m[2] = "DG"
	m[3] = "BCMP"
	m[4] = "FHVWY"
	m[5] = "K"
	m[8] = "JX"
	m[10] = "QZ"
	var score int
	score = 0
	for _, letter := range word {
		for key, element := range m {
			for _, l := range element {
				if letter == l {
					score += key
				}
			}
		}
	}
	return score
}

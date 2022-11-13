package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var s_output string
	// s_output = "aaaaaaaaaaaabbbbbbbcdefffffffff"
	s_output = "  hsqq qww  "
	fmt.Println(s_output)
	s_output = RunLengthEncode(s_output)
	fmt.Println(s_output)
	s_output = RunLengthDecode(s_output)
	fmt.Println(s_output)
}

func RunLengthEncode(input string) string {
	var output string
	var first_letter string
	count := 1

	for _, l := range input {
		if first_letter == "" {
			first_letter = string(l)
			continue
		}
		if first_letter != string(l) {
			if count == 1 {
				output += first_letter
			} else {
				output += strconv.FormatInt(int64(count), 10) + first_letter
			}
			first_letter = string(l)
			count = 1
			continue
		} else {
			count++
		}
	}

	if first_letter != "" {
		if count == 1 {
			fmt.Println(count)
			output += first_letter
		} else {
			output += strconv.FormatInt(int64(count), 10) + first_letter
		}
	}

	return output
}

func RunLengthDecode(input string) string {
	// panic("Please implement the RunLengthDecode function")
	var output string
	var count string

	for _, l := range input {
		if !unicode.IsNumber(l) && count == "" {
			output += string(l)
		} else if !unicode.IsNumber(l) && count != "" {
			s_int, err := strconv.ParseInt(count, 10, 64)
			if err != nil {
				return "Error"
			}
			output += strings.Repeat(string(l), int(s_int))
			count = ""
		} else {
			count += string(l)
		}
	}
	return output
}

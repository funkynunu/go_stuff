package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// import (
// 	"github.com/funkynunu/go_rest_api/pkg/leapyear"
// )

// {"0", 0, true},
// {"1", 1, true},
// {"2", 2, true},
// {"10", 3, true},
// {"201", 19, true},
// {"0201", 19, true},
// {"0000000000000000000000000000000000000000201", 19, true},
// {"2021110011022210012102010021220101220221", 9223372036854775807, true},
// {"2021110011022210012102010021220101220222", 0, false},

func main() {
	result, err := ParseTrinary("102012")
	// fmt.Println(result)
	// result, err = ParseTrinary("0")
	// fmt.Println(result)
	// result, err = ParseTrinary("1")
	// fmt.Println(result)
	// result, err = ParseTrinary("2")
	// fmt.Println(result)
	// result, err = ParseTrinary("10")
	// fmt.Println(result)
	// result, err = ParseTrinary("201")
	// fmt.Println(result)
	// result, err = ParseTrinary("0201")
	// fmt.Println(result)
	// result, err = ParseTrinary("0000000000000000000000000000000000000000201")
	// fmt.Println(result)
	// result, err = ParseTrinary("2021110011022210012102010021220101220221")
	// fmt.Println(result)
	// result, err = ParseTrinary("2021110011022210012102010021220101220222")
	// fmt.Println(result)
	// result, err = ParseTrinary("2021110011022210012102010021220101220223")
	fmt.Println(result)
	fmt.Println(err)
}

// Package acronym implements a utility for creating an acronym from
// a given phrase
func ParseTrinary(s string) (int64, error) {
	// check if s is valid trinary number
	for _, c := range s {
		if c != '0' && c != '1' && c != '2' {
			return 0, errors.New("Invalid trinary")
		}
	}

	var length = len(s)
	var value int64
	for _, l := range s {
		length--
		if string(l) == "0" {
			continue
		}
		s_int, err := strconv.ParseInt(string(l), 10, 64)
		if err != nil {
			return 0, err
		}

		// add check for overflow
		temp_value := value + (int64(s_int) * int64(math.Pow(3, float64(length))))
		if temp_value < value {
			return math.MaxInt64, nil
		}
		value += int64(s_int) * int64(math.Pow(3, float64(length)))

	}

	return value, nil
}

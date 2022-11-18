// You are given an encoded string consisting of a sequence of numeric characters: 0..9, $s.

// Write a script to find the all valid different decodings in sorted order.

// Encoding is simply done by mapping A,B,C,D,… to 1,2,3,4,… etc.

// Example 1
// Input: $s = 11
// Output: AA, K

// 11 can be decoded as (1 1) or (11) i.e. AA or K
// Example 2
// Input: $s = 1115
// Output: AAAE, AAO, AKE, KAE, KO

// Possible decoded data are:
// (1 1 1 5) => (AAAE)
// (1 1 15)  => (AAO)
// (1 11 5)  => (AKE)
// (11 1 5)  => (KAE)
// (11 15)   => (KO)
// Example 3
// Input: $s = 127
// Output: ABG, LG

// Possible decoded data are:
// (1 2 7) => (ABG)
// (12 7)  => (LG)
package main

import (
	"fmt"
	"strings"
)

var code_ref map[string]string

// m := make(map[string]string)

func main() {
	code_ref["1"] = "A"
	code_ref["2"] = "B"
	code_ref["3"] = "C"
	code_ref["4"] = "D"
	code_ref["5"] = "E"
	code_ref["6"] = "F"
	code_ref["7"] = "G"
	code_ref["8"] = "H"
	code_ref["9"] = "I"
	code_ref["10"] = "J"
	code_ref["11"] = "K"
	code_ref["12"] = "L"
	code_ref["13"] = "M"
	code_ref["14"] = "N"
	code_ref["15"] = "O"
	code_ref["16"] = "P"
	code_ref["17"] = "Q"
	code_ref["18"] = "R"
	code_ref["19"] = "S"
	code_ref["20"] = "T"
	code_ref["21"] = "U"
	code_ref["22"] = "V"
	code_ref["23"] = "W"
	code_ref["24"] = "X"
	code_ref["25"] = "Y"
	code_ref["26"] = "Z"
	fmt.Println(Decoded_rec("1115")) // AAAE, AAO, AKE, KAE, KO

}

// Package acronym implements a utility for creating an acronym from
// a given phrase
func Decoded_rec(s string) {
	// perl solution TODO: translate to go
	return $_[0] eq '' ? '' : $_[0] eq '0' ? () : chr(64 + $_[0]) if 2 > length $_[0];
	my($f,$s,$r) = split m{}, $_[0], 3;
	$r ||= '';
	( $f      && $s                ? ( map { chr(           $f + 64 ) . $_ } Decoded_rec($s.$r) ) : (),
	  $f == 1 || $f == 2 && $s < 7 ? ( map { chr( $f * 10 + $s + 64 ) . $_ } Decoded_rec($r   ) ) : () );
  }



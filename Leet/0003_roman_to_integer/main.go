package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("IX"))
}

func romanToInt(s string) int {
	roman := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	romanSubst := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}

	var prevChar string
	var subNumber int
	var output int

	for _, rune := range s {
		char := string(rune)
		if prevChar != "" {
			subNumber = romanSubst[prevChar+char]

			if subNumber != 0 {
				output -= roman[prevChar]
				output += subNumber
				continue
			}
		}

		prevChar = char
		output += roman[char]

	}

	return output
}

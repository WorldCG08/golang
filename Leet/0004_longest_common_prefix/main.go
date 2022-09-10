package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"tes3t1", "test2", "test3"}))
}

func longestCommonPrefix(strs []string) string {
	shortWord := getShortWord(strs)

	var chars string
	match := make(map[string]string)
	counter := 0

	for i := 0; i < shortWord; i++ {
		var char string
		for _, str := range strs {
			char = string(str[counter])
			match[char] = char
		}

		if len(match) == 1 {
			chars += match[char]
		} else {
			break
		}

		// clear map
		for k := range match {
			delete(match, k)
		}
		counter++
	}

	return chars
}

func getShortWord(strArr []string) int {
	shortWord := len(strArr[0])

	for _, str := range strArr {
		if len(str) < shortWord {
			shortWord = len(str)
		}
	}

	return shortWord
}

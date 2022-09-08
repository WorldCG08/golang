package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(95))
	fmt.Println(isPalindrome(10102))
	fmt.Println(isPalindrome(3301033))
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	newInt := 0
	y := x
	for y > 0 {
		remainder := y % 10
		newInt *= 10
		newInt += remainder
		y /= 10
	}

	return newInt == x
}

package main

import "fmt"

func main() {
	var nums []int = twoNums([]int{1, 5, 10, 3, 2}, 3)

	fmt.Println(nums)
}

func twoNums(integers []int, target int) []int {
	var rem, rem2 int
	var find bool
	var result []int
	for i, _ := range integers {
		rem = i
		for j, _ := range integers {
			if j == rem {
				continue
			}
			if (integers[rem] + integers[j]) == target {
				rem2 = j
				find = true
				break
			}
		}
		if find == true {
			break
		}
	}
	result = append(result, rem, rem2)

	return result
}

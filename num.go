package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	sumExpected := (n * (n + 1)) / 2
	sumActual := 0

	for _, num := range nums {
		sumActual += num
	}

	return sumExpected - sumActual
}

func main() {
	nums := []int{3, 0, 1, 2, 5}
	result := missingNumber(nums)
	fmt.Println("The missing number is:", result)
}

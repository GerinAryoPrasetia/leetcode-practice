package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
	// fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	num := make(map[int]int)

	for i, n := range nums {
		pairResult := target - n
		if idx, ok := num[pairResult]; ok {
			return []int{idx, i}
		}
		num[n] = i
	}

	return nil
}

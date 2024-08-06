package main

import "fmt"

func main() {

	// fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{5, -3, 9, 1, 7, 7, 9, 10, 2, 2, 10, 10, 3, -1, 3, 7, -9, -1, 3, 3}, 3))
}

func topKFrequent(nums []int, k int) []int {
	num := make(map[int]int)
	result := make([]int, 0)
	for _, n := range nums {
		num[n]++
	}

	for key := range num {
		fmt.Println(key)
		if len(result) < k {
			result = append(result, key)
			continue
		}
		// fmt.Println(result)
		for idx, n := range result {
			if num[key] > num[n] {
				// fmt.Printf("num[key] : %d-%d, num[n]: %d-%d\n", num[key], key, num[n], n)
				result = append(result[:idx], result[idx+1:]...)
				result = append(result, key)
			}

		}
	}
	fmt.Println(num)
	return result
}

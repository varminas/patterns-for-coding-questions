package main

import (
	"fmt"
)

/**
 * Maximum Sum Subarray of Size K (easy)
 */
func main() {
	fmt.Println("Sliding Window")
	fmt.Println("Maximum Sum Subarray of Size K (easy)")
	executeFindMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2})
	executeFindMaxSumSubArray(2, []int{2, 3, 4, 1, 5})
}

func executeFindMaxSumSubArray(k int, arr []int) {
	result := findMaxSumSubArray(k, arr)
	fmt.Println("1.---------")
	fmt.Printf("Input: %d, k: %d\n", arr, k)
	fmt.Printf("Output: %d\n", result)
}

func findMaxSumSubArray(k int, arr []int) int {
	var windowStart = 0
	var maxSum = 0
	var sum = 0
	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		sum += arr[windowEnd]
		if windowEnd >= k-1 {
			maxSum = Max(sum, maxSum)
			sum -= arr[windowStart]
			windowStart++
		}
	}

	return maxSum
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

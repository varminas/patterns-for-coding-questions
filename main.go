package main

import (
	"fmt"
	"github.com/varminas/patterns-for-coding-questions/slidingwindow"
)

func main() {
	// Sliding window
	fmt.Println("Sliding Window")
	fmt.Println("Maximum Sum Subarray of Size K (easy)")
	executeFindMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2})
	executeFindMaxSumSubArray(2, []int{2, 3, 4, 1, 5})
	// =======================================================
	fmt.Println("Smallest Subarray With a Greater Sum (easy)")
	executeFindMinSubArray(7, []int{2, 1, 5, 2, 3, 2})
}

func executeFindMaxSumSubArray(k int, arr []int) {
	result := slidingwindow.FindMaxSumSubArray(k, arr)
	fmt.Println("1.---------")
	fmt.Printf("Input: %d, k: %d\n", arr, k)
	fmt.Printf("Output: %d\n", result)
}

func executeFindMinSubArray(k int, arr []int) {
	result := slidingwindow.FindMinSubArray(k, arr)
	fmt.Println("1.---------")
	fmt.Printf("Input: %d, k: %d\n", arr, k)
	fmt.Printf("Output: %d\n", result)
}

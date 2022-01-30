// Pattern: Sliding Window
package slidingwindow

import (
	"github.com/varminas/patterns-for-coding-questions/math"
)

// 1. Maximum Sum Subarray of Size K (easy)
func FindMaxSumSubArray(k int, arr []int) int {
	var windowStart = 0
	var maxSum = 0
	var sum = 0
	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		sum += arr[windowEnd]
		if windowEnd >= k-1 {
			maxSum = math.Max(sum, maxSum)
			sum -= arr[windowStart]
			windowStart++
		}
	}

	return maxSum
}

// 2. Smallest Subarray With a Greater Sum (easy)
func FindMinSubArray(S int, arr []int) int {
	var windowStart = 0
	var sum = 0
	var minInterval = len(arr)
	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		//fmt.Printf("window [%v-%v]\n", windowStart, windowEnd)
		sum += arr[windowEnd]

		for sum >= S {
			//fmt.Printf("sum:%v\n", sum)
			minInterval = math.Min((windowEnd - windowStart + 1), minInterval)
			sum -= arr[windowStart]
			windowStart++
		}
	}
	return minInterval
}

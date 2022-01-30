// Pattern: Sliding Window
// Maximum Sum Subarray of Size K (easy)

package slidingwindow

import (
	"github.com/varminas/patterns-for-coding-questions/math"
)

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

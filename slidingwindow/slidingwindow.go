// Pattern: Sliding Window
package slidingwindow

import (
	"github.com/cornelk/hashmap"
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

// 3. Longest Substring with maximum K Distinct Characters (medium)
func FindLongestSubstringWithMaxKDistChars(str string, k int) int {
	var windowStart = 0
	var maxLength = 0
	var uniqueChars = &hashmap.HashMap{}

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		var rightChar = str[windowEnd]
		var count = 1
		currentCount, ok := uniqueChars.Get(rightChar)
		if ok {
			count = currentCount.(int)
		}
		uniqueChars.Set(rightChar, count)

		for uniqueChars.Len() > k {
			var leftChar = str[windowStart]
			count = 0
			currentLeftCount, ok := uniqueChars.Get(leftChar)
			if ok {
				count = currentLeftCount.(int)
			}
			uniqueChars.Set(leftChar, count-1)

			currentLeftCount, ok = uniqueChars.Get(leftChar)
			if ok && currentLeftCount == 0 {
				uniqueChars.Del(leftChar)
			}
			windowStart++
		}

		maxLength = math.Max(maxLength, (windowEnd - windowStart + 1))
	}
	return maxLength
}

// 4. Fruits into Baskets (medium)
func FindMaxFruitCountOf2Types(input []string) int {
	windowStart := 0
	uniqueChars := &hashmap.HashMap{}
	maxSize := 0
	for windowEnd := 0; windowEnd < len(input); windowEnd++ {

		rightChar := input[windowEnd]
		count := 0
		currentRightCount, ok := uniqueChars.Get(rightChar)
		if ok {
			count = currentRightCount.(int)
		}
		uniqueChars.Set(rightChar, count+1)

		for uniqueChars.Len() > 2 {
			leftChar := input[windowStart]
			count = 0
			currentLeftCount, ok := uniqueChars.Get(leftChar)
			if ok {
				count = currentLeftCount.(int)
			}
			uniqueChars.Set(leftChar, count-1)
			currentLeftCount, ok = uniqueChars.Get(leftChar)
			if ok && currentLeftCount == 0 {
				uniqueChars.Del(leftChar)
			}
			windowStart++
		}

		maxSize = math.Max(maxSize, windowEnd-windowStart+1)
	}
	return maxSize
}

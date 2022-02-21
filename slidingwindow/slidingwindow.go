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
		uniqueChars.Set(rightChar, getOrDefault(uniqueChars, rightChar, 0)+1)

		for uniqueChars.Len() > k {
			var leftChar = str[windowStart]
			uniqueChars.Set(leftChar, getOrDefault(uniqueChars, leftChar, 0)-1)

			currentLeftCount, ok := uniqueChars.Get(leftChar)
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
		uniqueChars.Set(rightChar, getOrDefault(uniqueChars, rightChar, 0)+1)

		for uniqueChars.Len() > 2 {
			leftChar := input[windowStart]

			uniqueChars.Set(leftChar, getOrDefault(uniqueChars, leftChar, 0)-1)
			currentLeftCount, ok := uniqueChars.Get(leftChar)
			if ok && currentLeftCount == 0 {
				uniqueChars.Del(leftChar)
			}
			windowStart++
		}

		maxSize = math.Max(maxSize, windowEnd-windowStart+1)
	}
	return maxSize
}

// 5. Longest Substring with Distinct Characters (hard)
func FindNoRepeatSubstring(input string) int {
	windowStart := 0
	maxLength := 0
	charIndexMap := &hashmap.HashMap{}
	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		rightChar := input[windowEnd]
		val, ok := charIndexMap.Get(rightChar)
		if ok {
			windowStart = math.Max(windowStart, val.(int)+1)
		} else {
			charIndexMap.Set(rightChar, windowEnd)
			maxLength = math.Max(maxLength, windowEnd-windowStart+1)
		}

	}
	return maxLength
}

// 6. Longest Substring with Same Letters after Replacement (hard)
func FindCharacterReplacementLength(inputStr string, k int) int {
	input := []rune(inputStr)
	windowStart := 0
	maxLength := 0
	maxRepeatedLettersCount := 0
	charsMap := &hashmap.HashMap{}

	for windowEnd, rightChar := range input {
		val := getOrDefault(charsMap, rightChar, 0) + 1
		charsMap.Set(rightChar, val)
		maxRepeatedLettersCount = math.Max(maxRepeatedLettersCount, val)

		if windowEnd-windowStart+1-maxRepeatedLettersCount > k {
			leftChar := input[windowStart]
			charsMap.Set(leftChar, getOrDefault(charsMap, leftChar, 0)-1)
			windowStart++
		}
		maxLength = math.Max(maxLength, windowEnd-windowStart+1)
	}
	return maxLength
}

// 7. Longest Subarray with Ones after Replacement (hard)
func FindReplacingOnesLength(arr []int, k int) int {
	maxLength := 0
	windowStart := 0
	maxRepeatedValuesCount := 0

	for windowEnd, rightValue := range arr {
		if rightValue == 1 {
			maxRepeatedValuesCount++
		}

		if windowEnd-windowStart+1-maxRepeatedValuesCount > k {
			leftValue := arr[windowStart]
			if leftValue == 1 {
				maxRepeatedValuesCount--
			}
			windowStart++
		}

		maxLength = math.Max(maxLength, windowEnd-windowStart+1)
	}
	return maxLength
}

// 8. Problem Challenge 1. Permutation in a String (hard)
func HasStringPermutation(inputStr string, pattern string) bool {
	input := []rune(inputStr)
	windowStart := 0
	matchingCharsCount := 0
	//matchingCharsMap := make(map[string]int)

	patternCharsCountMap := make(map[string]int)
	for _, patternValue := range pattern {
		char := string(patternValue)
		patternCharsCountMap[char] = getOrDefaultOnMap(patternCharsCountMap, char, 0) + 1
	}

	for _, inputCharAsRune := range input {
		inputChar := string(inputCharAsRune)
		_, ok := patternCharsCountMap[inputChar]
		if ok {
			//matchingCharsMap[inputChar] = getOrDefaultOnMap(matchingCharsMap, inputChar, 0) + 1
			matchingCharsCount++
		} else {
			windowStart++
			if matchingCharsCount > 0 {
				matchingCharsCount--
			}
		}

		if matchingCharsCount == len(pattern) {
			return true
		}
	}
	return false
}

// 9. String Anagrams (hard)
func FindListOfStartingIndicesOfAnagrams(inputStr string, pattern string) []int {
	input := []rune(inputStr)
	windowStart := 0
	patternCharsMap := make(map[string]int)
	matchingCharsCount := 0
	patternLength := len(pattern)
	startingIndices := []int{}
	for _, charRune := range pattern {
		char := string(charRune)
		patternCharsMap[char] = getOrDefaultOnMap(patternCharsMap, char, 0) + 1
	}

	for windowEnd, charRightRune := range input {
		charRight := string(charRightRune)
		charRightCount, ok := patternCharsMap[charRight]
		if ok {
			patternCharsMap[charRight] = charRightCount - 1
			matchingCharsCount++
			if charRightCount == 0 {
				for start := windowStart; start < windowEnd; start++ {
					charLeft := string(input[start])
					charLeftCount, _ := patternCharsMap[charLeft]
					patternCharsMap[charLeft] = charLeftCount + 1
				}
				matchingCharsCount = 1
				windowStart = windowEnd
			}
		}

		if matchingCharsCount == patternLength {
			startingIndices = append(startingIndices, windowStart)
			patternCharsMap[string(input[windowStart])]++
			matchingCharsCount--
			windowStart++
		}
	}

	return startingIndices
}

// 10. Smallest Window containing Substring (hard)
func FindSmallestWindowContainingSubstring(input string, patternAsString string) string {
	patternArray := []rune(patternAsString)
	patternCharsMap := make(map[string]int)

	windowStart := 0
	matched := 0
	minLength := len(input) + 1
	subStrStart := 0

	for _, char := range patternArray {
		charAsString := string(char)
		patternCharsMap[charAsString]++
	}

	for windowEnd, rightCharAsRune := range input {
		rightChar := string(rightCharAsRune)

		rightCharCount, ok := patternCharsMap[rightChar]
		if ok {
			rightCharCount--
			patternCharsMap[rightChar] = rightCharCount
			if rightCharCount >= 0 {
				matched++
			}
		}

		for matched == len(patternAsString) {
			if minLength > windowEnd-windowStart+1 {
				minLength = windowEnd - windowStart + 1
				subStrStart = windowStart
			}

			leftChar := string(input[windowStart])
			windowStart++
			leftCharCount, leftCharFound := patternCharsMap[leftChar]
			if leftCharFound {
				if leftCharCount == 0 {
					matched--
				}
				patternCharsMap[leftChar] = leftCharCount + 1
			}
		}
	}

	if minLength > len(input) {
		return ""
	} else {
		return input[subStrStart : subStrStart+minLength]
	}
}

func getOrDefaultOnMap(inputMap map[string]int, str string, defaultValue int) int {
	currentLeftCount, exists := inputMap[str]
	if exists {
		return currentLeftCount
	} else {
		return defaultValue
	}
}

func getOrDefault(inputMap *hashmap.HashMap, str interface{}, defaultValue int) int {
	currentLeftCount, ok := inputMap.Get(str)
	if ok {
		return currentLeftCount.(int)
	} else {
		return defaultValue
	}
}

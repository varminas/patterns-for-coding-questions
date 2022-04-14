package pointers

import (
	"fmt"
	"math"
	"sort"
)

// 1. Pair with Target Sum (easy)
func PairWithTargetSum(input []int, target int) []int {
	beginPointer := 0
	endPointer := len(input) - 1
	for i := 0; i < len(input)/2+1; i++ {
		sum := input[beginPointer] + input[endPointer]
		if sum == target {
			return []int{beginPointer, endPointer}
		} else if sum > target {
			endPointer--
			continue
		} else {
			beginPointer++
			continue
		}
	}
	return nil
}

// 2. Remove Duplicates (easy)
func RemoveDuplicates(input []int) int {
	if len(input) == 1 {
		return 1
	}

	nextNonDuplicate := 1

	for next := 0; next < len(input); next++ {
		if input[next] != input[nextNonDuplicate-1] {
			input[nextNonDuplicate] = input[next]
			nextNonDuplicate++
		}
	}
	return nextNonDuplicate
}

// 2.1. Remove Duplicates (easy) - unsorted array and target
func RemoveDuplicatesFromUnsorted(input []int, key int) int {
	if len(input) == 1 {
		return 1
	}

	nextElement := 0

	for next := 0; next < len(input); next++ {
		if input[next] != key {
			input[nextElement] = input[next]
			nextElement++
		}
	}
	return nextElement
}

// 3. Squaring a Sorted Array (easy)
func SquaringSortedArray(input []int) []int {
	squares := make([]int, len(input))

	begin := 0
	end := len(input) - 1
	highestSquareIndex := len(input) - 1

	for begin <= end {
		beginSquare := input[begin] * input[begin]
		endSquare := input[end] * input[end]
		if beginSquare > endSquare {
			squares[highestSquareIndex] = beginSquare
			highestSquareIndex--
			begin++
		} else {
			squares[highestSquareIndex] = endSquare
			highestSquareIndex--
			end--
		}
	}
	return squares
}

// 4. Triplet Sum to Zero (medium)
func TripletSumToZero(arr []int) [][]int {
	sort.Ints(arr)

	triplets := make([][]int, 0)
	for i := 0; i < len(arr)-2; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		searchPair(arr, -arr[i], i+1, &triplets)
	}
	return triplets
}

func searchPair(arr []int, targetSum int, left int, triplets *[][]int) {
	right := len(arr) - 1
	for left < right {
		currentSum := arr[left] + arr[right]
		if currentSum == targetSum {
			*triplets = append(*triplets, []int{-targetSum, arr[left], arr[right]})
			left++
			right--
			for left < right && arr[left] == arr[left-1] {
				left++
			}
			for left < right && arr[right] == arr[right+1] {
				right--
			}
		} else if targetSum > currentSum {
			left++
		} else {
			right--
		}
	}
}

//	5. Triplet Sum Close to Target (medium)
func TripletSumCloseToTarget(arr []int, targetSum int) int {
	sort.Ints(arr)

	smallestDifference := math.MaxInt32
	for i := 0; i < len(arr)-2; i++ {
		left := i + 1
		right := len(arr) - 1
		for left < right {
			targetDiff := targetSum - arr[i] - arr[left] - arr[right]
			if targetDiff == 0 {
				return targetSum
			}

			if math.Abs(float64(targetDiff)) < math.Abs(float64(smallestDifference)) ||
				(math.Abs(float64(targetDiff)) == math.Abs(float64(smallestDifference)) &&
					targetDiff > smallestDifference) {
				smallestDifference = targetDiff
			}

			if targetDiff > 0 {
				left++
			} else {
				right--
			}
		}
	}
	return targetSum - smallestDifference
}

// 6. Triplets with Smaller Sum (medium)
func TripletsWithSmallerSum(arr []int, target int) int {
	sort.Ints(arr)

	smallerCount := 0
	for i := 0; i < len(arr)-2; i++ {
		left := i + 1
		right := len(arr) - 1
		targetSum := target - arr[i]
		for left < right {
			if arr[left]+arr[right] < targetSum {
				fmt.Printf("[%d %d %d]\n", arr[i], arr[left], arr[right])
				smallerCount += right - left
				left++
			} else {
				right--
			}
		}
	}
	fmt.Println("")
	return smallerCount
}

// 7. Subarrays with Product Less than a Target (medium)
func SubarraysWithProductLessThanTarget(arr []int, target int) [][]int {

	var result = make([][]int, 0)
	product := 1
	left := 0
	for right := 0; right < len(arr); right++ {
		product *= arr[right]
		for product >= target && left < len(arr) {
			product /= arr[left]
			left++
		}
		var tempList = make([]int, 0)
		for i := right; i >= left; i-- {
			tempList = append(tempList, arr[i])
			result = append(result, tempList)
		}
	}
	return result
}

// 8. Dutch National Flag Problem (medium)
func DutchNationalFlagProblem(arr []int) []int {
	low := 0
	high := len(arr) - 1
	for i := 0; i <= high; {
		if arr[i] == 0 {
			swap(arr, i, low)
			i++
			low++
		} else if arr[i] == 1 {
			i++
		} else {
			swap(arr, i, high)
			high--
		}
	}
	return arr
}

func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

// 9. Quadruple Sum to Target (medium)
func QuadrupleSumToTarget(arr []int, target int) [][]int {
	sort.Ints(arr)

	quadruplets := make([][]int, 0)
	for i := 0; i < len(arr)-3; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		for j := i + 1; j < len(arr)-2; j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}
			searchPairs(arr, target, i, j, &quadruplets)
		}
	}
	return quadruplets
}

func searchPairs(arr []int, targerSum int, first int, second int, quadruplets *[][]int) {
	left := second + 1
	right := len(arr) - 1
	for left < right {
		sum := arr[first] + arr[second] + arr[left] + arr[right]
		if sum == targerSum {
			*quadruplets = append(*quadruplets, []int{arr[first], arr[second], arr[left], arr[right]})
			left++
			right--
			for left < right && arr[left] == arr[left-1] {
				left++
			}
			for left < right && arr[right] == arr[right+1] {
				right--
			}
		} else if sum < targerSum {
			left++
		} else {
			right--
		}
	}
}

// 10. Comparing Strings containing Backspaces (medium)
func ComparingStringsContainingBackspaces(str1 string, str2 string) bool {
	return true
}

// 11. Minimum Window Sort (medium)
func MinimumWindowSort(arr []int) int {
	return 0
}

package pointers

import (
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

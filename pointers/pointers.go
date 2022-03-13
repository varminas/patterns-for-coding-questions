package pointers

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
	// TODO
	return nil
}

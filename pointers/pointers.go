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
	// TODO
	return 0
}

// 3. Squaring a Sorted Array (easy)
func SquaringSortedArray(input []int) []int {
	// TODO
	return nil
}

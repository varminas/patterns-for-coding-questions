package slidingwindow

import (
	"fmt"
	"testing"
)

func TestSlidingWindow_maxSumSubarrayOfSizeK3(t *testing.T) {
	want := 9
	input := []int{2, 1, 5, 1, 3, 2}

	got := FindMaxSumSubArray(3, input)
	if got != want {
		t.Errorf(errorString(input, got, want))
	}
}

func TestSlidingWindow_maxSumSubarrayOfSizeK2(t *testing.T) {
	want := 7
	input := []int{2, 3, 4, 1, 5}

	got := FindMaxSumSubArray(2, input)
	if got != want {
		t.Errorf(errorString(input, got, want))
	}
}

func TestSlidingWindow_FindMinSubArray_2(t *testing.T) {
	want := 2
	input := []int{2, 1, 5, 2, 3, 2}

	got := FindMinSubArray(7, input)
	if got != want {
		t.Errorf(errorString(input, got, want))
	}
}

func TestSlidingWindow_FindMinSubArray_1(t *testing.T) {
	want := 1
	input := []int{2, 1, 5, 2, 8}

	got := FindMinSubArray(7, input)
	if got != want {
		t.Errorf(errorString(input, got, want))
	}
}

func TestSlidingWindow_FindMinSubArray_3(t *testing.T) {
	want := 3
	input := []int{3, 4, 1, 1, 6}

	got := FindMinSubArray(8, input)
	if got != want {
		t.Errorf(errorString(input, got, want))
	}
}

func errorString(input []int, got int, want int) string {
	return fmt.Sprintf("findMaxSumSubArray(%#v) \n\ngot= %#v \n\nwant=%#v", input, got, want)
}

func errorStringArr(input []int, got []int, want []int) string {
	return fmt.Sprintf("findMaxSumSubArray(%#v) \n\ngot= %#v \n\nwant=%#v", input, got, want)
}

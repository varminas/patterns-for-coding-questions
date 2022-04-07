package pointers

import (
	"fmt"
	"reflect"
	"testing"
)

// 1. Pair with Target Sum (easy)
func TestPairWithTargetSum_v1(t *testing.T) {
	want := []int{1, 3}
	input := []int{1, 2, 3, 4, 6}

	got := PairWithTargetSum(input, 6)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("PairWithTargetSum", input, got, want))
	}
}

func TestPairWithTargetSum_v2(t *testing.T) {
	want := []int{0, 2}
	input := []int{2, 5, 9, 11}

	got := PairWithTargetSum(input, 11)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("PairWithTargetSum", input, got, want))
	}
}

// 2. Remove Duplicates (easy)
func TestRemoveDuplicates_v1(t *testing.T) {
	want := 4
	input := []int{2, 3, 3, 3, 6, 9, 9}

	got := RemoveDuplicates(input)
	if got != want {
		t.Errorf(errorString("RemoveDuplicates", input, got, want))
	}
}

func TestRemoveDuplicates_v2(t *testing.T) {
	want := 2
	input := []int{2, 2, 2, 11}

	got := RemoveDuplicates(input)
	if got != want {
		t.Errorf(errorString("RemoveDuplicates", input, got, want))
	}
}

// 2.1. Remove Duplicates (easy) - unsorted array and target
func TestRemoveDuplicatesFromUnsorted_v1(t *testing.T) {
	want := 4
	input := []int{3, 2, 3, 6, 3, 10, 9, 3}

	got := RemoveDuplicatesFromUnsorted(input, 3)
	if got != want {
		t.Errorf(errorString("RemoveDuplicatesFromUnsorted", input, got, want))
	}
}

func TestRemoveDuplicatesFromUnsorted_v2(t *testing.T) {
	want := 2
	input := []int{2, 11, 2, 2, 1}

	got := RemoveDuplicatesFromUnsorted(input, 2)
	if got != want {
		t.Errorf(errorString("RemoveDuplicatesFromUnsorted", input, got, want))
	}
}

// 3. Squaring a Sorted Array (easy)
func TestSquaringSortedArray_v1(t *testing.T) {
	want := []int{0, 1, 4, 4, 9}
	input := []int{-2, -1, 0, 2, 3}

	got := SquaringSortedArray(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("SquaringSortedArray", input, got, want))
	}
}

func TestSquaringSortedArray_v2(t *testing.T) {
	want := []int{0, 1, 1, 4, 9}
	input := []int{-3, -1, 0, 1, 2}

	got := SquaringSortedArray(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("SquaringSortedArray", input, got, want))
	}
}

// 4. Triplet Sum to Zero (medium)
func TestTripletSumToZero_v1(t *testing.T) {
	input := []int{-3, 0, 1, 2, -1, 1, -2}
	want := [][]int{{-3, 1, 2}, {-2, 0, 2}, {-2, 1, 1}, {-1, 0, 1}}

	got := TripletSumToZero(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("TripletSumToZero", input, got, want))
	}
}

func TestTripletSumToZero_v2(t *testing.T) {
	input := []int{-5, 2, -1, -2, 3}
	want := [][]int{{-5, 2, 3}, {-2, -1, 3}}

	got := TripletSumToZero(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("TripletSumToZero", input, got, want))
	}
}

func TestTripletSumCloseToTarget_v1(t *testing.T) {
	input := []int{-2, 0, 1, 2}
	want := 1

	got := TripletSumCloseToTarget(input, 2)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("TripletSumCloseToTarget", input, got, want))
	}
}

func TestTripletSumCloseToTarget_v2(t *testing.T) {
	input := []int{-3, -1, 1, 2}
	want := 0

	got := TripletSumCloseToTarget(input, 1)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("TripletSumCloseToTarget", input, got, want))
	}
}

func TestTripletSumCloseToTarget_v3(t *testing.T) {
	input := []int{1, 0, 1, 1}
	want := 3

	got := TripletSumCloseToTarget(input, 100)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("TripletSumCloseToTarget", input, got, want))
	}
}

// 6. Triplets with Smaller Sum (medium)
func TestTripletsWithSmallerSum_v1(t *testing.T) {
	input := []int{-1, 0, 2, 3}
	want := 2

	got := TripletsWithSmallerSum(input, 3)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("TripletsWithSmallerSum", input, got, want))
	}
}

func TestTripletsWithSmallerSum_v2(t *testing.T) {
	input := []int{-1, 4, 2, 1, 3}
	want := 4

	got := TripletsWithSmallerSum(input, 5)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("TripletsWithSmallerSum", input, got, want))
	}
}

//// 7. Subarrays with Product Less than a Target (medium)
//func TestSubarraysWithProductLessThanTarget_v1(t *testing.T) {
//	input := []int{2, 5, 3, 10}
//	want := [][]int{{2}, {5}, {2, 5}, {3}, {5, 3}, {10}}
//
//	got := SubarraysWithProductLessThanTarget(input, 30)
//	if !reflect.DeepEqual(got, want) {
//		t.Errorf(errorString("SubarraysWithProductLessThanTarget", input, got, want))
//	}
//}

//func TestSubarraysWithProductLessThanTarget_v2(t *testing.T) {
//	input := []int{8, 2, 6, 5}
//	want := [][]int{{8}, {2}, {8, 2}, {6}, {2, 6}, {5}, {6, 5}}
//
//	got := SubarraysWithProductLessThanTarget(input, 50)
//	if !reflect.DeepEqual(got, want) {
//		t.Errorf(errorString("SubarraysWithProductLessThanTarget", input, got, want))
//	}
//}

//// 8. Dutch National Flag Problem (medium)
//func TestDutchNationalFlagProblem_v2(t *testing.T) {
//	input := []int{8, 2, 6, 5}
//	want := [][]int{{8}, {2}, {8, 2}, {6}, {2, 6}, {5}, {6, 5}}
//
//	got := DutchNationalFlagProblem(input)
//	if !reflect.DeepEqual(got, want) {
//		t.Errorf(errorString("DutchNationalFlagProblem", input, got, want))
//	}
//}

func errorString(funcName string, input interface{}, got interface{}, want interface{}) string {
	return fmt.Sprintf("%v(%#v) \n\ngot= %#v \n\nwant=%#v", funcName, input, got, want)
}

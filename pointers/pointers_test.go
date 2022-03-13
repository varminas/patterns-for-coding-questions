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

//// 3. Squaring a Sorted Array (easy)
//func TestSquaringSortedArray_v1(t *testing.T) {
//	want := []int{0, 1, 4, 4, 9}
//	input := []int{-2, -1, 0, 2, 3}
//
//	got := SquaringSortedArray(input)
//	if !reflect.DeepEqual(got,want) {
//		t.Errorf(errorString("SquaringSortedArray", input, got, want))
//	}
//}
//
//func TestSquaringSortedArray_v2(t *testing.T) {
//	want := []int{0, 1, 1, 4, 9}
//	input := []int{-3, -1, 0, 1, 2}
//
//	got := SquaringSortedArray(input)
//	if !reflect.DeepEqual(got,want) {
//		t.Errorf(errorString("SquaringSortedArray", input, got, want))
//	}
//}

func errorString(funcName string, input interface{}, got interface{}, want interface{}) string {
	return fmt.Sprintf("%v(%#v) \n\ngot= %#v \n\nwant=%#v", funcName, input, got, want)
}

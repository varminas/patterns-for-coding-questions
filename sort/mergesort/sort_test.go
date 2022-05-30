package mergesort

import (
	"fmt"
	"github.com/varminas/patterns-for-coding-questions/utils"
	"reflect"
	"testing"
)

func TestMergesort_v1(t *testing.T) {
	input := []int{38, 27, 43, 3, 9, 82, 10}
	got := Mergesort(input)
	fmt.Printf("Quicksort result: %v", input)
	want := []int{3, 9, 10, 27, 38, 43, 82}

	if !reflect.DeepEqual(got, want) {
		t.Errorf(utils.PrintErrorString("Mergesort", input, got, want))
	}
}

func TestMergesort_v2(t *testing.T) {
	input := []int{10, 6, 8, 5, 7, 3, 4}
	got := Mergesort(input)
	fmt.Printf("Quicksort result: %v", input)
	want := []int{3, 4, 5, 6, 7, 8, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf(utils.PrintErrorString("Mergesort", input, got, want))
	}
}

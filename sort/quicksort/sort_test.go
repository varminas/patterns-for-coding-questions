package quicksort

import (
	"fmt"
	"github.com/varminas/patterns-for-coding-questions/utils"
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	input := []int{10, 80, 30, 90, 40, 50, 70}
	got := Quicksort(input)
	fmt.Printf("Quicksort result: %v", input)
	want := []int{10, 30, 40, 50, 70, 80, 90}

	if !reflect.DeepEqual(got, want) {
		t.Errorf(utils.PrintErrorString("Sort", input, got, want))
	}
}

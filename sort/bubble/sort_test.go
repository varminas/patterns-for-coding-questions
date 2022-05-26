package bubble

import (
	"fmt"
	"github.com/varminas/patterns-for-coding-questions/utils"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	input := []int{5, 1, 4, 2, 8}
	got := Sort(input)
	fmt.Printf("Bubble sort result: %v", input)
	want := []int{1, 2, 4, 5, 8}

	if !reflect.DeepEqual(got, want) {
		t.Errorf(utils.PrintErrorString("Sort", input, got, want))
	}
}

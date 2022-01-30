package math

import (
	"testing"
)

func TestFindMax(t *testing.T) {
	got := Max(1, 2)

	if got != 2 {
		t.Errorf("Should be 2")
	}
}

package slidingwindow

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func TestSlidingWindow_maxSumSubarrayOfSizeK3(t *testing.T) {
	want := 9
	input := []int{2, 1, 5, 1, 3, 2}

	got := FindMaxSumSubArray(3, input)
	if got != want {
		t.Errorf(errorString("FindMaxSumSubArray", input, got, want))
	}
}

func TestSlidingWindow_maxSumSubarrayOfSizeK2(t *testing.T) {
	want := 7
	input := []int{2, 3, 4, 1, 5}

	got := FindMaxSumSubArray(2, input)
	if got != want {
		t.Errorf(errorString("FindMaxSumSubArray", input, got, want))
	}
}

func TestSlidingWindow_FindMinSubArray_2(t *testing.T) {
	want := 2
	input := []int{2, 1, 5, 2, 3, 2}

	got := FindMinSubArray(7, input)
	if got != want {
		t.Errorf(errorString("FindMinSubArray", input, got, want))
	}
}

func TestSlidingWindow_FindMinSubArray_1(t *testing.T) {
	want := 1
	input := []int{2, 1, 5, 2, 8}

	got := FindMinSubArray(7, input)
	if got != want {
		t.Errorf(errorString("FindMinSubArray", input, got, want))
	}
}

func TestSlidingWindow_FindMinSubArray_3(t *testing.T) {
	want := 3
	input := []int{3, 4, 1, 1, 6}

	got := FindMinSubArray(8, input)
	if got != want {
		t.Errorf(errorString("FindMinSubArray", input, got, want))
	}
}

func TestSlidingWindow_FindLongestSubstringWithMaxKDistChars_K2(t *testing.T) {
	want := 4
	input := "araaci"

	got := FindLongestSubstringWithMaxKDistChars(input, 2)
	if got != want {
		t.Errorf(errorString("FindLongestSubstringWithMaxKDistChars", input, got, want))
	}
}

func TestSlidingWindow_FindLongestSubstringWithMaxKDistChars_K1(t *testing.T) {
	want := 2
	input := "araaci"

	got := FindLongestSubstringWithMaxKDistChars(input, 1)
	if got != want {
		t.Errorf(errorString("FindLongestSubstringWithMaxKDistChars", input, got, want))
	}
}

func TestSlidingWindow_FindLongestSubstringWithMaxKDistChars_K3(t *testing.T) {
	want := 5
	input := "cbbebi"

	got := FindLongestSubstringWithMaxKDistChars(input, 3)
	if got != want {
		t.Errorf(errorString("FindLongestSubstringWithMaxKDistChars", input, got, want))
	}
}

func TestSlidingWindow_FindLongestSubstringWithMaxKDistChars_K10(t *testing.T) {
	want := 6
	input := "cbbebi"

	got := FindLongestSubstringWithMaxKDistChars(input, 10)
	if got != want {
		t.Errorf(errorString("FindLongestSubstringWithMaxKDistChars", input, got, want))
	}
}

func TestSlidingWindow_FindMaxFruitCountOf2Types_V1(t *testing.T) {
	want := 3
	input := []string{"A", "B", "C", "A", "C"}

	got := FindMaxFruitCountOf2Types(input)
	if got != want {
		t.Errorf(errorString("FindMaxFruitCountOf2Types", input, got, want))
	}
}

func TestSlidingWindow_FindMaxFruitCountOf2Types_V2(t *testing.T) {
	want := 5
	input := []string{"A", "B", "C", "B", "B", "C"}

	got := FindMaxFruitCountOf2Types(input)
	if got != want {
		t.Errorf(errorString("FindMaxFruitCountOf2Types", input, got, want))
	}
}

func TestFindNoRepeatSubstring_V1(t *testing.T) {
	want := 3
	input := "aabccbb"

	got := FindNoRepeatSubstring(input)
	if got != want {
		t.Errorf(errorString("FindNoRepeatSubstring", input, got, want))
	}
}

func TestFindNoRepeatSubstring_V2(t *testing.T) {
	want := 2
	input := "abbbb"

	got := FindNoRepeatSubstring(input)
	if got != want {
		t.Errorf(errorString("FindNoRepeatSubstring", input, got, want))
	}
}

func TestFindNoRepeatSubstring_V3(t *testing.T) {
	want := 3
	input := "abccde"

	got := FindNoRepeatSubstring(input)
	if got != want {
		t.Errorf(errorString("FindNoRepeatSubstring", input, got, want))
	}
}

func TestFindNoRepeatSubstring_V4(t *testing.T) {
	want := 25
	input := "abcdefghijklmnoprstqvwxyz"

	got := FindNoRepeatSubstring(input)
	if got != want {
		t.Errorf(errorString("FindNoRepeatSubstring", input, got, want))
	}
}

func TestFindCharacterReplacementLength_V1(t *testing.T) {
	want := 5
	input := "aabccbb"

	got := FindCharacterReplacementLength(input, 2)
	if got != want {
		t.Errorf(errorString("FindCharacterReplacementLength", input, got, want))
	}
}

func TestFindCharacterReplacementLength_V2(t *testing.T) {
	want := 4
	input := "abbcb"

	got := FindCharacterReplacementLength(input, 1)
	if got != want {
		t.Errorf(errorString("FindCharacterReplacementLength", input, got, want))
	}
}

func TestFindCharacterReplacementLength_V3(t *testing.T) {
	want := 3
	input := "abccde"

	got := FindCharacterReplacementLength(input, 1)
	if got != want {
		t.Errorf(errorString("FindCharacterReplacementLength", input, got, want))
	}
}

func TestFindReplacingOnesLength_V1(t *testing.T) {
	want := 6
	input := []int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}

	got := FindReplacingOnesLength(input, 2)
	if got != want {
		t.Errorf(errorString("FindReplacingOnesLength", input, got, want))
	}
}

func TestFindReplacingOnesLength_V2(t *testing.T) {
	want := 9
	input := []int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}

	got := FindReplacingOnesLength(input, 3)
	if got != want {
		t.Errorf(errorString("FindReplacingOnesLength", input, got, want))
	}
}

func TestHasStringPermutation_V1(t *testing.T) {
	want := true
	input := "oidbcaf"

	got := HasStringPermutation(input, "abc")
	if got != want {
		t.Errorf(errorString("HasStringPermutation", input, got, want))
	}
}

func TestHasStringPermutation_V2(t *testing.T) {
	want := false
	input := "odicf"

	got := HasStringPermutation(input, "dc")
	if got != want {
		t.Errorf(errorString("HasStringPermutation", input, got, want))
	}
}

func TestHasStringPermutation_V3(t *testing.T) {
	want := true
	input := "bcdxabcdy"

	got := HasStringPermutation(input, "bcdyabcdx")
	if got != want {
		t.Errorf(errorString("HasStringPermutation", input, got, want))
	}
}

func TestHasStringPermutation_V4(t *testing.T) {
	want := true
	input := "aaacb"

	got := HasStringPermutation(input, "bca")
	if got != want {
		t.Errorf(errorString("HasStringPermutation", input, got, want))
	}
}

// 9. String Anagrams (hard)
func TestFindListOfStartingIndicesOfAnagrams_V1(t *testing.T) {
	want := []int{1, 2}
	input := "ppqp"

	got := FindListOfStartingIndicesOfAnagrams(input, "pq")
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("FindListOfStartingIndicesOfAnagrams", input, got, want))
	}
}

func TestFindListOfStartingIndicesOfAnagrams_V2(t *testing.T) {
	want := []int{2, 3, 4}
	input := "abbcabc"

	got := FindListOfStartingIndicesOfAnagrams(input, "abc")
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("FindListOfStartingIndicesOfAnagrams", input, got, want))
	}
}

// 10. Smallest Window containing Substring (hard)
func TestFindSmallestWindowContainingSubstring_V1(t *testing.T) {
	want := "abdec"
	input := "aabdec"

	got := FindSmallestWindowContainingSubstring(input, "abc")
	if got != want {
		t.Errorf(errorString("FindSmallestWindowContainingSubstring", input, got, want))
	}
}

func TestFindSmallestWindowContainingSubstring_V2(t *testing.T) {
	want := "aabdec"
	input := "aabdec"

	got := FindSmallestWindowContainingSubstring(input, "abac")
	if got != want {
		t.Errorf(errorString("FindSmallestWindowContainingSubstring", input, got, want))
	}
}

func TestFindSmallestWindowContainingSubstring_V3(t *testing.T) {
	want := "bca"
	input := "abdbca"

	got := FindSmallestWindowContainingSubstring(input, "abc")
	if got != want {
		t.Errorf(errorString("FindSmallestWindowContainingSubstring", input, got, want))
	}
}

func TestFindSmallestWindowContainingSubstring_V4(t *testing.T) {
	want := ""
	input := "adcad"

	got := FindSmallestWindowContainingSubstring(input, "abc")
	if got != want {
		t.Errorf(errorString("FindSmallestWindowContainingSubstring", input, got, want))
	}
}

func errorString(funcName string, input interface{}, got interface{}, want interface{}) string {
	return fmt.Sprintf("%v(%#v) \n\ngot= %#v \n\nwant=%#v", funcName, input, got, want)
}

func getFuncName(handler interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
}

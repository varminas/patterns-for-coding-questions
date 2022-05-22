package fastslowpointers

import (
	"fmt"
	"reflect"
	"testing"
)

// 1. LinkedList Cycle
func TestLinkedListCycle_v1(t *testing.T) {
	head := ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next.Next.Next = &ListNode{6, nil}
	want := false

	got := LinkedListCycle(head)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("LinkedListCycle", head, got, want))
	}
}

func TestLinkedListCycle_v2(t *testing.T) {
	head := ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next.Next.Next = &ListNode{6, nil}
	head.Next.Next.Next.Next.Next.Next = head.Next.Next
	want := true

	got := LinkedListCycle(head)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("LinkedListCycle", head, got, want))
	}
}

func TestLinkedListCycle_v3(t *testing.T) {
	head := ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next.Next.Next = &ListNode{6, nil}
	head.Next.Next.Next.Next.Next.Next = head.Next.Next.Next
	want := true

	got := LinkedListCycle(head)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("LinkedListCycle", head, got, want))
	}
}

// 1.2 Given the head of a LinkedList with a cycle, find the length of the cycle
func TestLinkedListCycleLength_v1(t *testing.T) {
	head := ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next.Next.Next = &ListNode{6, nil}
	head.Next.Next.Next.Next.Next.Next = head.Next.Next
	want := 4

	got := LinkedListCycleLength(head)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("LinkedListCycleLength", head, got, want))
	}
}

func TestLinkedListCycleLength_v2(t *testing.T) {
	head := ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next.Next.Next = &ListNode{6, nil}
	head.Next.Next.Next.Next.Next.Next = head.Next.Next.Next
	want := 3

	got := LinkedListCycleLength(head)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("LinkedListCycleLength", head, got, want))
	}
}

// 2. Start of LinkedList Cycle (medium)
func TestStartOfLinkedListCycle_v1(t *testing.T) {
	head := ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next.Next.Next = &ListNode{6, nil}
	head.Next.Next.Next.Next.Next.Next = head.Next.Next
	want := 3

	got := StartOfLinkedListCycle(head)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("StartOfLinkedListCycle", head, got, want))
	}
}

func TestStartOfLinkedListCycle_v2(t *testing.T) {
	head := ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next.Next.Next = &ListNode{6, nil}
	head.Next.Next.Next.Next.Next.Next = head.Next.Next.Next
	want := 4

	got := StartOfLinkedListCycle(head)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("StartOfLinkedListCycle", head, got, want))
	}
}

func TestStartOfLinkedListCycle_v3(t *testing.T) {
	head := ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next.Next.Next = &ListNode{6, nil}
	head.Next.Next.Next.Next.Next.Next = &head
	want := 1

	got := StartOfLinkedListCycle(head)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("StartOfLinkedListCycle", head, got, want))
	}
}

// 3. Happy Number (medium)
func TestHappyNumber_v1(t *testing.T) {
	input := 23
	got := HappyNumber(input)
	want := true
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("HappyNumber", input, got, want))
	}
}

func TestHappyNumber_v2(t *testing.T) {
	input := 12
	got := HappyNumber(input)
	want := false
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("HappyNumber", input, got, want))
	}
}

// 4. Middle of the LinkedList
func TestMiddleOfTheLinkedList_v1(t *testing.T) {
	head := ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	want := 3

	got := MiddleOfTheLinkedList(head)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("MiddleOfTheLinkedList", head, got, want))
	}
}

// 5. Palindrome LinkedList (medium)
func TestPalindromeLinkedList_v1(t *testing.T) {
	head := ListNode{2, nil}
	head.Next = &ListNode{4, nil}
	head.Next.Next = &ListNode{6, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{2, nil}
	want := true

	got := PalindromeLinkedList(head)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("PalindromeLinkedList", head, got, want))
	}
}

func TestPalindromeLinkedList_v2(t *testing.T) {
	head := ListNode{2, nil}
	head.Next = &ListNode{4, nil}
	head.Next.Next = &ListNode{6, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{2, nil}
	head.Next.Next.Next.Next.Next = &ListNode{2, nil}
	want := false

	got := PalindromeLinkedList(head)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("PalindromeLinkedList", head, got, want))
	}
}

// 6.

func TestMiddleOfTheLinkedList_v2(t *testing.T) {
	head := ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next.Next.Next = &ListNode{6, nil}
	want := 4

	got := MiddleOfTheLinkedList(head)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("MiddleOfTheLinkedList", head, got, want))
	}
}

func TestMiddleOfTheLinkedList_v3(t *testing.T) {
	head := ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next.Next.Next = &ListNode{6, nil}
	head.Next.Next.Next.Next.Next.Next = &ListNode{7, nil}
	want := 4

	got := MiddleOfTheLinkedList(head)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("MiddleOfTheLinkedList", head, got, want))
	}
}

// 6. Rearrange a LinkedList (medium)
func TestRearrangeLinkedList_v1(t *testing.T) {
	head := ListNode{2, nil}
	head.Next = &ListNode{4, nil}
	head.Next.Next = &ListNode{6, nil}
	head.Next.Next.Next = &ListNode{8, nil}
	head.Next.Next.Next.Next = &ListNode{10, nil}
	head.Next.Next.Next.Next.Next = &ListNode{12, nil}
	want := "2 -> 12 -> 4 -> 10 -> 6 -> 8 -> nil"

	got := RearrangeLinkedList(head)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("RearrangeLinkedList", head, got, want))
	}
}

func TestRearrangeLinkedList_v2(t *testing.T) {
	head := ListNode{2, nil}
	head.Next = &ListNode{4, nil}
	head.Next.Next = &ListNode{6, nil}
	head.Next.Next.Next = &ListNode{8, nil}
	head.Next.Next.Next.Next = &ListNode{10, nil}
	want := "2 -> 10 -> 4 -> 8 -> 6 -> nil"

	got := RearrangeLinkedList(head)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorString("RearrangeLinkedList", head, got, want))
	}
}

// 7. Cycle in a Circular Array
// TODO
//func TestCycleInCircularArray_v1(t *testing.T) {
//	input := []int{1, 2, -1, 2, 2}
//	want := true
//
//	got := CycleInCircularArray(input)
//	if !reflect.DeepEqual(got, want) {
//		t.Errorf(errorString("CycleInCircularArray", input, got, want))
//	}
//}
//
//func TestCycleInCircularArray_v2(t *testing.T) {
//	input := []int{2, 2, -1, 2}
//	want := true
//
//	got := CycleInCircularArray(input)
//	if !reflect.DeepEqual(got, want) {
//		t.Errorf(errorString("CycleInCircularArray", input, got, want))
//	}
//}
//
//func TestCycleInCircularArray_v3(t *testing.T) {
//	input := []int{2, 1, -1, -2}
//	want := false
//
//	got := CycleInCircularArray(input)
//	if !reflect.DeepEqual(got, want) {
//		t.Errorf(errorString("CycleInCircularArray", input, got, want))
//	}
//}

func errorString(funcName string, input interface{}, got interface{}, want interface{}) string {
	return fmt.Sprintf("%v(%#v) \n\ngot= %#v \n\nwant=%#v", funcName, input, got, want)
}

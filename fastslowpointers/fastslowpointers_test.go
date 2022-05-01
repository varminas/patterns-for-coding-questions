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

func errorString(funcName string, input interface{}, got interface{}, want interface{}) string {
	return fmt.Sprintf("%v(%#v) \n\ngot= %#v \n\nwant=%#v", funcName, input, got, want)
}

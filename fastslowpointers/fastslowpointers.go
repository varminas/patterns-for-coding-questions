package fastslowpointers

// Pattern: Fast & Slow pointers

type ListNode struct {
	Value int
	Next  *ListNode
}

// 1. LinkedList Cycle
func LinkedListCycle(head ListNode) bool {
	slow := &head
	fast := &head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

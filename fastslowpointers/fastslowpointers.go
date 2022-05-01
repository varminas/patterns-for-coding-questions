package fastslowpointers

// Pattern: Fast & Slow pointers

type ListNode struct {
	Value int
	Next  *ListNode
}

// 1.1 LinkedList Cycle
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

// 1.2 Given the head of a LinkedList with a cycle, find the length of the cycle
func LinkedListCycleLength(head ListNode) int {
	slow := &head
	fast := &head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return calculateCycleLength(slow)
		}
	}
	return 0
}

func calculateCycleLength(slow *ListNode) int {
	current := slow
	cycleLength := 0
	for {
		current = current.Next
		cycleLength++
		if current == slow {
			break
		}
	}
	return cycleLength
}

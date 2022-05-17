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

// 2. Start of LinkedList Cycle (medium)
func StartOfLinkedListCycle(head ListNode) int {
	cycleLength := LinkedListCycleLength(head)
	pointer1 := &head
	pointer2 := &head
	for i := 0; i < cycleLength; i++ {
		pointer2 = pointer2.Next
	}
	for {
		if pointer1.Value == pointer2.Value {
			break
		}
		pointer1 = pointer1.Next
		pointer2 = pointer2.Next
	}
	return pointer1.Value
}

// 3. Happy Number (medium)
func HappyNumber(num int) bool {
	slow := num
	fast := num

	for {
		slow = findSquareSum(slow)
		fast = findSquareSum(findSquareSum(fast))
		if fast == slow {
			break
		}
	}

	return slow == 1
}

func findSquareSum(num int) int {
	digit := 0
	sum := 0
	for num > 0 {
		digit = num % 10
		sum += digit * digit
		num /= 10
	}
	return sum
}

// 4. Middle of the LinkedList
func MiddleOfTheLinkedList(head ListNode) int {
	slow := &head
	fast := &head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow.Value
}

// 5. Palindrome LinkedList (medium)
func PalindromeLinkedList(head ListNode) bool {
	slow := &head
	fast := &head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse list
	headSecondHalf := reverse(slow)
	copyOfSecondHalf := headSecondHalf

	//fmt.Printf("Middle is %v\n", slow.Value)
	//fmt.Printf("Head second half is %v\n", headSecondHalf.Value)

	// compare the first and the second half
	for &head != nil && headSecondHalf != nil {
		if head.Value != headSecondHalf.Value {
			break // not a palindrome
		}
		if head.Next == nil {
			break
		}
		head = *head.Next
		headSecondHalf = headSecondHalf.Next
	}

	reverse(copyOfSecondHalf)
	if head.Next == nil || headSecondHalf == nil {
		return true
	}
	return false
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

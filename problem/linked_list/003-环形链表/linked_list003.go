package linked_list003

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if nil == head {
		return false
	}

	fast := head.Next
	// if head != nil && fast != nil && fast.Next != nil {
	if fast != nil && fast.Next != nil {
		if fast == head {
			return true
		}
		fast = fast.Next.Next
		head = head.Next
	}
	return false
}

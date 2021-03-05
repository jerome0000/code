package linked_list

type ListNode struct {
	Val  string
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preNode := head
	curNode := head
	for i := 0; i < n; i++ {
		curNode = curNode.Next
	}
	if curNode == nil {
		return preNode.Next
	}
	for curNode.Next != nil {
		preNode = preNode.Next
		curNode = curNode.Next
	}
	preNode.Next = preNode.Next.Next
	return head
}

package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		beg = nil
		mid = head
		end = head.Next
	)
	for {
		mid.Next = beg
		if end == nil {
			break
		}
		beg = mid
		mid = end
		end = end.Next
	}
	head = mid
	return head
}

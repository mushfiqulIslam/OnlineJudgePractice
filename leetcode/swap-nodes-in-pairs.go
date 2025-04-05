/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	secondNode := head.Next
	if secondNode == nil {
		return head
	}

	dummy := ListNode{}
	current := &dummy
	for head != nil && secondNode != nil {
		head.Next = secondNode.Next
		secondNode.Next = head
		current.Next = secondNode
		
		current = current.Next
		current = current.Next

		head = head.Next
		if head != nil {
			secondNode = head.Next
		} else {
			secondNode = nil
		}
	}
	if head != nil && secondNode == nil {
		current.Next = head
	}
	return dummy.Next

}

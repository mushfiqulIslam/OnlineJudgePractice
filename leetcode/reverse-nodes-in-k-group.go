func reverseKGroup(head *ListNode, k int) *ListNode {
	var toReverseHead, reversedHead *ListNode
	dummy := ListNode{}
	current := &dummy
	count := 1

	for head != nil {
		next := head.Next
		if count == 1 {
			toReverseHead = head
		} else if count == k {
			last := toReverseHead
			reversedHead = toReverseHead
			toReverseHead = toReverseHead.Next
			reversedHead.Next = nil
			count--
			for count > 0 {
				tmp := toReverseHead
				toReverseHead = toReverseHead.Next
				tmp.Next = reversedHead
				reversedHead = tmp
				count--
			}
			current.Next = reversedHead
			current = last
		}

		head = next
		count++
	}

	if count > 0 {
		current.Next = toReverseHead
	}

	return dummy.Next
}

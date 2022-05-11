func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sumVal, carry int
	sumHead := new(ListNode)
	cursor := sumHead
	temp1 := l1
	temp2 := l2
	for temp1 != nil || temp2 != nil || carry != 0 {
		sumVal = carry
		if temp1 != nil {
			sumVal += temp1.Val
            temp1 = temp1.Next
            
		}

		if temp2 != nil {
			sumVal += temp2.Val
            temp2 = temp2.Next
		}

		carry = sumVal / 10
        cursor.Next = &ListNode{
            Val:  sumVal % 10,
            Next: nil,
        }
		
		cursor = cursor.Next
	}
	return sumHead.Next

}

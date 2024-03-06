package main

func hasCycle(head *ListNode) bool {
	visited := make(map[*ListNode]bool)
	for head != nil {
		if visited[head] {
			return true
		}
		visited[head] = true
		head = head.Next
	}
	return false
}

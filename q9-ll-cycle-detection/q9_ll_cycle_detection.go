package q9_ll_cycle_detection

// leetcode #142
// Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
// Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed).
// It is -1 if there is no cycle. Note that pos is not passed as a parameter.

type ListNode struct {
	Val  int
	Next *ListNode
}

// Iterate every node and remember all previous nodes in a set(golang map)
// and if the current node is already in a set then there is cycle detected
func detectCycleBruteForce(head *ListNode) *ListNode { // Time: O(n), Space: O(n)
	curr := head
	seenNodes := map[int]struct{}{}

	for ok := false; !ok; _, ok = seenNodes[curr.Val] {
		if curr.Next == nil {
			return nil // no cycle
		}
		seenNodes[curr.Val] = struct{}{}
		curr = curr.Next
	}

	return curr
}

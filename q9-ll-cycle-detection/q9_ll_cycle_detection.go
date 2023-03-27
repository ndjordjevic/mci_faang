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

// Uses Floyd's Tortoise and Hare algorithm with two pointers, Tortoise moves one node at the time, Hare moves two nodes
// We instantiate both pointers and set to the head
// Then infinite loop until we break either when we found they are met, or Hare is nil or points to nil
// If they meet we have a cycle
// then we have to find at what node they are met
// instantiate two pointers, p1 and p2, one points to the head, the other points to the either Tortoise or Hare
// and move them one step at a time until they meet and that would be the node the cycle has started
func detectCycleFloydTortoiseHair(head *ListNode) *ListNode { // Time: O(n), Space: O(1)
	if head == nil {
		return nil
	}

	tortoise, hare := head, head

	for {
		hare = hare.Next
		tortoise = tortoise.Next

		if hare == nil || hare.Next == nil {
			return nil
		}
		hare = hare.Next

		if tortoise == hare {
			break
		}
	}

	p1, p2 := head, tortoise
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}

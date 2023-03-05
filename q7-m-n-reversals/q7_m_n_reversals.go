package q7_m_n_reversals

// leetcode #92
// Given the head of a singly linked list and two integers m and n where m <= n
// reverse the nodes of the list from position m to position n, and return the reversed list.
// 1. Constraints: 1 <= m <= n <= length of linked list
// 2. Test cases:
// 1->2->3->4->5->nil, m=2, n=4: 1->4->3->2->5->nil
// 1->2->3->4->5->nil, m=1, n=5: 5->4->3->2->1->nil
// 1->nil, m=1, n=1: 1->nil
// nil, m=0, n=0: nil

type ListNode struct {
	Val  int
	Next *ListNode
}

// we need to track: m-1, m, n, n+1
// we are receiving head, and need a var to track position and the current node
// start at head and check if current node is one from the: m-1, m, n, n+1
// no: just advance to the next node and update current and position
// yes, == to m-1: record last value before we start reversion into start var and advance
// yes, == to m: current value needs to be recorded into tail var
// no we are in the portion of the ll that needs to be reversed (see the img.png), we got new var: newList (named prev in the other reversed ll challenge)
// we need to stop reversing when the position var is > n var
// when we are done with portion reversing attach newList to start's next and current to tail's next
// and return the head!
// 1->2->3->4->5->nil, m=2, n=4: 1->4->3->2->5->nil
func reverseBetween(head *ListNode, m, n int) *ListNode { // Space: O(1), Time: O(n), head: 1->2->3->4->5->nil, m=2, n=4
	if head == nil {
		return nil
	}

	position := 1       // position=5
	currentNode := head // currentNode=5->nil
	startNode := head   // startNode=1->4->3->2->nil

	// find the start node (m-1) first, here at the end of this loop we get start == node before reverse should begin, so we have start node
	for position < m {
		startNode = currentNode
		currentNode = currentNode.Next
		position++
	}

	var newList *ListNode   // newList: 4->3->2->nil
	tailNode := currentNode // tailNode: 2->5->nil
	for position <= n {
		nextNode := currentNode.Next // nextNode: 5->nil
		currentNode.Next = newList
		newList = currentNode
		currentNode = nextNode
		position++
	}
	startNode.Next = newList
	tailNode.Next = currentNode

	if m > 1 {
		return head
	}

	return newList
}

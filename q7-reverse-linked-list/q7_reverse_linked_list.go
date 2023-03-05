package q7_reverse_linked_list

// leetcode #206
// Given the head of a singly linked list, reverse the list, and return the reversed list.
// 1. Constraints: if null or a single el linked list is passed return null and that element respectively
// 2. Test cases:
// 1->2->3->4->5->nil: 5->4->3->2->1->nil
// 3->nil: 3->nil
// nil: nil

type ListNode struct {
	Val  int
	Next *ListNode
}

// Two vars prev & current and one local inside while loop called next
// prev init to nil (will be the new list we will build)
// current init to head
// while current != nil
// set local lop var next to current.next
// current.next = prev
// prev = current
// current = next
// end of while
// return prev
func reverseLinkedList(head *ListNode) *ListNode { // head: 1->2->3->4->5->nil Space: O(1), Time: O(n)
	var prev *ListNode // prev: 5->4->3->2->1->nil
	current := head    // current: nil

	for current != nil {
		next := current.Next // next: nil
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

package q8_merge_multilevel_doubly_linked_list

// leetcode #430
// You are given a doubly linked list, which contains nodes that have a next pointer, a previous pointer, and an additional child pointer.
// This child pointer may or may not point to a separate doubly linked list, also containing these special nodes.
// These child lists may have one or more children of their own, and so on, to produce a multilevel data structure.
// Given the head of the first level of the list, flatten the list so that all the nodes appear in a single-level, doubly linked list.
// Let curr be a node with a child list. The nodes in the child list should appear after curr and before curr.next in the flattened list.
// Return the head of the flattened list. The nodes in the list must have all of their child pointers set to null.
//
// 1. Constraints:
// Every node can have multiple levels of children
// After flattening set the child property to nil
// Test cases:
// see: img.png
// nil: nil
// 1->2: 1->2

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// Merging from the top to bottom
// iterate from the head node and check if it has a child node
// if yes remember the current node from the main flow
// then start traversing the child nodes till the tail is found (next is null)
// when the tail is found merge the child flow into the main flow with the following steps in the specific order
// fist merge child's tail
// then merge child's head and set current's child to null
// 4 nodes are important to track: current node, current's next, current's child and the tail of the child's linked list
func flatten(head *Node) *Node { // Space: O(1), Time: O(n)
	if head == nil {
		return head
	}

	currentNode := head

	for currentNode != nil {
		if currentNode.Child == nil {
			currentNode = currentNode.Next
			continue
		}

		tail := currentNode.Child

		for tail.Next != nil {
			tail = tail.Next
		}

		tail.Next = currentNode.Next
		if tail.Next != nil {
			tail.Next.Prev = tail
		}

		currentNode.Next = currentNode.Child
		currentNode.Next.Prev = currentNode
		currentNode.Child = nil
	}

	return head
}

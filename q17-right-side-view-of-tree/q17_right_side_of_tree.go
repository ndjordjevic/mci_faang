package q17_right_side_view_of_tree

// leetcode #199: binary-tree-right-side-view
// Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
// Test cases: see the img.png

// We are going to use BFS to solve this problem
// Define result array and a queue
// Add root node to the queue
// While queue is not empty
//   - Get the length of the queue in a length variable
//   - Get the count of the nodes in a count variable
//   - Define a currentNode variable
//   - While count < length
//   - Get the first node from the queue and remove it from the queue
//   - If the node has a left child add it to the queue
//   - If the node has a right child add it to the queue
//   - Increment the count
//
// Add the currentNode to the result array
func rightSideViewBFS(root *TreeNode) []int { // Time: O(n), Space: O(n)
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		length, count := len(queue), 0
		var currentNode *TreeNode
		for count < length {
			currentNode, queue = queue[0], queue[1:]
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
			count++
		}
		result = append(result, currentNode.Val)
	}

	return result
}

// We are going to use DFS to solve this problem with NRL (Node Right Left) Pre-Order approach
// Define result array
// Call the recursive function with root node, currentDepth = 0 and result array
// Return the result array
func rightSideViewDFS(root *TreeNode) []int { // Time: O(n), Space: O(n)
	var result []int
	rightSideViewDFSRecursive(root, 0, &result)
	return result
}

// Define the recursive function with node, currentDepth and result array
// If node is nil return
// If currentDepth >= len(result) then append the node value to the result array (this means we are at the right side of the tree)
// Increment the currentDepth
// Call the recursive function with node.Right, currentDepth and result array
// Call the recursive function with node.Left, currentDepth and result array
func rightSideViewDFSRecursive(node *TreeNode, currentDepth int, result *[]int) {
	if node == nil {
		return
	}

	if currentDepth >= len(*result) {
		*result = append(*result, node.Val)
	}

	currentDepth++
	rightSideViewDFSRecursive(node.Right, currentDepth, result)
	rightSideViewDFSRecursive(node.Left, currentDepth, result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeFromArrayBFS(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	values = values[1:]
	queue := []*TreeNode{root}
	var i int

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Left == nil {
			if values[i] != -1 {
				current.Left = &TreeNode{Val: values[i]}
			}
			i++
			if i >= len(values) {
				return root
			}
		}
		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right == nil {
			if values[i] != -1 {
				current.Right = &TreeNode{Val: values[i]}
			}
			i++
			if i >= len(values) {
				return root
			}
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return root
}

func arrayFromTreeBFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	currentNode := root
	var list []int
	var queue []*TreeNode

	queue = append(queue, currentNode)
	for len(queue) > 0 {
		currentNode, queue = queue[0], queue[1:]
		if currentNode == nil {
			list = append(list, -1)
			continue
		}
		list = append(list, currentNode.Val)
		queue = append(queue, currentNode.Left)
		queue = append(queue, currentNode.Right)
	}

	return list
}

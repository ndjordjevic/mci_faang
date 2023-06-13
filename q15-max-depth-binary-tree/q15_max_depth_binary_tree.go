package q15_max_depth_binary_tree

// leetcode #104
// Given the root of a binary tree, return its maximum depth.
// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
// Constraints: if the tree is empty return 0
// Test cases:
// best case where we have nodes return 5
// just a root node when we return 1
// empty we return 0
func maxDepth(root *TreeNode) int { // Time: O(n), Space: best: O(logn), worst: O(n)
	return maxDepthRecursive(root, 0)
}

func maxDepthRecursive(node *TreeNode, currentDepth int) int {
	if node == nil {
		return currentDepth
	}
	currentDepth++
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	return max(maxDepthRecursive(node.Left, currentDepth), maxDepthRecursive(node.Right, currentDepth))
}

type BinarySearchTree struct {
	rootNode *TreeNode
}

func (tree *BinarySearchTree) insertElement(value int) {
	newNode := &TreeNode{Val: value}
	if tree.rootNode == nil {
		tree.rootNode = newNode
	} else {
		insertNodeBST(tree.rootNode, newNode)
	}
}

func insertNodeBST(parentNode *TreeNode, newNode *TreeNode) {
	if newNode.Val < parentNode.Val {
		if parentNode.Left == nil {
			parentNode.Left = newNode
		} else {
			insertNodeBST(parentNode.Left, newNode)
		}
	}

	if newNode.Val > parentNode.Val {
		if parentNode.Right == nil {
			parentNode.Right = newNode
		} else {
			insertNodeBST(parentNode.Right, newNode)
		}
	}
}

func (tree *BinarySearchTree) breathFirstSearch() []int {
	currentNode := tree.rootNode
	var list []int
	var queue []*TreeNode

	queue = append(queue, currentNode)
	for len(queue) > 0 {
		currentNode, queue = queue[0], queue[1:]
		list = append(list, currentNode.Val)
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}

	return list
}

func (tree *BinarySearchTree) depthFirstSearchPreOrder() []int {
	var list []int
	return traversePreOrder(tree.rootNode, list)
}
func (tree *BinarySearchTree) depthFirstSearchInOrder() []int {
	var list []int
	return traverseInOrder(tree.rootNode, list)
}
func (tree *BinarySearchTree) depthFirstSearchPostOrder() []int {
	var list []int
	return traversePostOrder(tree.rootNode, list)
}

func traversePreOrder(node *TreeNode, list []int) []int {
	list = append(list, node.Val)
	if node.Left != nil {
		list = traversePreOrder(node.Left, list)
	}
	if node.Right != nil {
		list = traversePreOrder(node.Right, list)
	}
	return list
}
func traverseInOrder(node *TreeNode, list []int) []int {
	if node.Left != nil {
		list = traverseInOrder(node.Left, list)
	}
	list = append(list, node.Val)
	if node.Right != nil {
		list = traverseInOrder(node.Right, list)
	}
	return list
}
func traversePostOrder(node *TreeNode, list []int) []int {
	if node.Left != nil {
		list = traversePostOrder(node.Left, list)
	}
	if node.Right != nil {
		list = traversePostOrder(node.Right, list)
	}
	list = append(list, node.Val)
	return list
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

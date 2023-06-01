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
	return max(maxDepthRecursive(node.leftNode, currentDepth), maxDepthRecursive(node.rightNode, currentDepth))
}

type BinarySearchTree struct {
	rootNode *TreeNode
}

func (tree *BinarySearchTree) insertElement(value int) {
	newNode := &TreeNode{value: value}
	if tree.rootNode == nil {
		tree.rootNode = newNode
	} else {
		insertNodeBST(tree.rootNode, newNode)
	}
}

func insertNodeBST(parentNode *TreeNode, newNode *TreeNode) {
	if newNode.value < parentNode.value {
		if parentNode.leftNode == nil {
			parentNode.leftNode = newNode
		} else {
			insertNodeBST(parentNode.leftNode, newNode)
		}
	}

	if newNode.value > parentNode.value {
		if parentNode.rightNode == nil {
			parentNode.rightNode = newNode
		} else {
			insertNodeBST(parentNode.rightNode, newNode)
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
		list = append(list, currentNode.value)
		if currentNode.leftNode != nil {
			queue = append(queue, currentNode.leftNode)
		}
		if currentNode.rightNode != nil {
			queue = append(queue, currentNode.rightNode)
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
	list = append(list, node.value)
	if node.leftNode != nil {
		list = traversePreOrder(node.leftNode, list)
	}
	if node.rightNode != nil {
		list = traversePreOrder(node.rightNode, list)
	}
	return list
}
func traverseInOrder(node *TreeNode, list []int) []int {
	if node.leftNode != nil {
		list = traverseInOrder(node.leftNode, list)
	}
	list = append(list, node.value)
	if node.rightNode != nil {
		list = traverseInOrder(node.rightNode, list)
	}
	return list
}
func traversePostOrder(node *TreeNode, list []int) []int {
	if node.leftNode != nil {
		list = traversePostOrder(node.leftNode, list)
	}
	if node.rightNode != nil {
		list = traversePostOrder(node.rightNode, list)
	}
	list = append(list, node.value)
	return list
}

type TreeNode struct {
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

// {1, 2, 3, 4, nil, nil, 5, 6, nil, nil, nil, nil, nil, nil, 7}
// DOESN'T WORK AS EXPECTED
func (n *TreeNode) createBTFromList(values []*int) {
	*n = TreeNode{
		value: *values[0],
	}
	queue := []*TreeNode{n}
	i := 1
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		// left
		if current.leftNode == nil {
			if values[i] != nil {
				current.leftNode = &TreeNode{value: *values[i]}
			}
			i++
			if i >= len(values) {
				return
			}
		}
		if current.leftNode != nil {
			queue = append(queue, current.leftNode)
		}

		// right
		if current.rightNode == nil {
			if values[i] != nil {
				current.rightNode = &TreeNode{value: *values[i]}
			}
			i++
			if i >= len(values) {
				return
			}
		}
		if current.rightNode != nil {
			queue = append(queue, current.rightNode)
		}
	}
}

// {1, 2, 3, 4, nil, nil, 5, 6, nil, nil, nil, nil, nil, nil, 7}
// DOESN'T WORK AS EXPECTED
func (n *TreeNode) createBTFromList2(array []*int) {
	if array == nil || len(array) == 0 {
		return
	}
	var tnq []*TreeNode
	var iq []*int

	*n = TreeNode{
		value: *array[0],
	}
	tnq = append(tnq, n)

	for i := 1; i < len(array); i++ {
		iq = append(iq, array[i])
	}

	for len(iq) > 0 {
		var lv, rv *int
		if len(iq) > 0 {
			lv = iq[0]
			iq = iq[1:]
		}
		if len(iq) > 0 {
			rv = iq[0]
			iq = iq[1:]
		}
		current := tnq[0]
		tnq = tnq[1:]

		if lv != nil {
			l := &TreeNode{value: *lv}
			current.leftNode = l
			tnq = append(tnq, l)
		}
		if rv != nil {
			r := &TreeNode{value: *rv}
			current.rightNode = r
			tnq = append(tnq, r)
		}
	}
}

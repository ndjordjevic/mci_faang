package q16_binary_tree_level_order

// leetcode #102
// Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level)
//

func levelOrder(root *TreeNode) [][]int { // Time: O(n), Space: O(n)
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length, count := len(queue), 0
		var currentLevelValues []int
		for count < length {
			currentNode := queue[0]
			queue = queue[1:]
			currentLevelValues = append(currentLevelValues, currentNode.Val)
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
			count++
		}
		result = append(result, currentLevelValues)
	}

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

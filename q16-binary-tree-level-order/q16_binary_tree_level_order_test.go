package q16_binary_tree_level_order

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	root := treeFromArrayBFS([]int{3, 9, 20, -1, -1, 15, 7})
	want := [][]int{{3}, {9, 20}, {15, 7}}
	if got := levelOrder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("levelOrder() = %v, want %v", got, want)
	}
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

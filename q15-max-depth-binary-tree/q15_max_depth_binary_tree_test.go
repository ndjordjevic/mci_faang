package q15_max_depth_binary_tree

import (
	"reflect"
	"testing"
)

func TestBreathFirstSearch(t *testing.T) {
	var bst BinarySearchTree
	bst.insertElement(9)
	bst.insertElement(4)
	bst.insertElement(6)
	bst.insertElement(20)
	bst.insertElement(170)
	bst.insertElement(15)
	bst.insertElement(1)

	got := bst.breathFirstSearch()
	want := []int{9, 4, 20, 1, 6, 15, 170}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("bst.breathFirstSearh()=%v, but want=%v", got, want)
	}
}

func TestDepthFirstSearchPreOrder(t *testing.T) {
	var bst BinarySearchTree
	bst.insertElement(9)
	bst.insertElement(4)
	bst.insertElement(6)
	bst.insertElement(20)
	bst.insertElement(170)
	bst.insertElement(15)
	bst.insertElement(1)

	got := bst.depthFirstSearchPreOrder()
	want := []int{9, 4, 1, 6, 20, 15, 170}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("bst.depthFirstSearchPreOrder()=%v, but want=%v", got, want)
	}
}

func TestDepthFirstSearchInOrder(t *testing.T) {
	var bst BinarySearchTree
	bst.insertElement(9)
	bst.insertElement(4)
	bst.insertElement(6)
	bst.insertElement(20)
	bst.insertElement(170)
	bst.insertElement(15)
	bst.insertElement(1)

	got := bst.depthFirstSearchInOrder()
	want := []int{1, 4, 6, 9, 15, 20, 170}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("bst.depthFirstSearchInOrder()=%v, but want=%v", got, want)
	}
}

func TestDepthFirstSearchPostOrder(t *testing.T) {
	var bst BinarySearchTree
	bst.insertElement(9)
	bst.insertElement(4)
	bst.insertElement(6)
	bst.insertElement(20)
	bst.insertElement(170)
	bst.insertElement(15)
	bst.insertElement(1)

	got := bst.depthFirstSearchPostOrder()
	want := []int{1, 6, 4, 15, 170, 20, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("bst.depthFirstSearchPostOrder()=%v, but want=%v", got, want)
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
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		} else {
			queue = append(queue, nil)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		} else {
			queue = append(queue, nil)
		}
	}

	return list
}

func Test_maxDepth(t *testing.T) {
	array := []int{3, 9, 20, -1, -1, 15, 7}

	root := treeFromArrayBFS(array)
	want := 3

	if got := maxDepth(root); got != want {
		t.Errorf("maxDepth() = %v, want %v", got, want)
	}
}

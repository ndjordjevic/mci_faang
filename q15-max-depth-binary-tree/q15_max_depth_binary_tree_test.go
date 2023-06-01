package q15_max_depth_binary_tree

import (
	"fmt"
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

func TestTreeNode_createBTFromList(t *testing.T) {
	intPt := func(i int) *int {
		pt := new(int)
		*pt = i
		return pt
	}

	root := new(TreeNode)
	values := []*int{intPt(1), intPt(2), intPt(3), intPt(4), nil, nil, intPt(5), intPt(6), nil, nil, nil, nil, nil, nil, intPt(7)}

	root.createBTFromList(values)
	fmt.Println(root)
}

func TestTreeNode_createBTFromList2(t *testing.T) {
	intPt := func(i int) *int {
		pt := new(int)
		*pt = i
		return pt
	}

	root := new(TreeNode)
	values := []*int{intPt(1), intPt(2), intPt(3), intPt(4), nil, nil, intPt(5), intPt(6), nil, nil, nil, nil, nil, nil, intPt(7)}

	root.createBTFromList2(values)
	fmt.Println(root)
}

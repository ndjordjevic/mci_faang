package q17_right_side_view_of_tree

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_rightSideViewBFS(t *testing.T) {
	tree := treeFromArrayBFS([]int{1, 2, 3, 4, 5, -1, 6, -1, 7, -1, -1, -1, -1, 8, -1, -1, -1})
	want := []int{1, 3, 6, 7, 8}

	if got := rightSideViewBFS(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("rightSideViewBFS() = %v, want %v", got, want)
	}
}

func Test_rightSideViewDFS(t *testing.T) {
	tree := treeFromArrayBFS([]int{1, 2, 3, 4, 5, -1, 6, -1, 7, -1, -1, -1, -1, 8, -1, -1, -1})
	want := []int{1, 3, 6, 7, 8}

	if got := rightSideViewDFS(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("rightSideViewDFS() = %v, want %v", got, want)
	}
}

func Test_treeFromArrayBFS_ArrayFromTreeBFS(t *testing.T) {
	inputArray := []int{1, 2, 3, 4, 5, -1, 6, -1, 7, -1, -1, -1, -1, 8, -1, -1, -1}

	gotTree := treeFromArrayBFS(inputArray)

	gotArray := arrayFromTreeBFS(gotTree)
	if !reflect.DeepEqual(gotArray, inputArray) {
		t.Errorf("arrayFromTreeBFS() = %v, want %v", gotArray, inputArray)
	}

	fmt.Println(gotTree)
}

package q13_kth_largest_element_array

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		array []int
		left  int
		right int
		want  []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should quick sort an array",
			args: args{
				array: []int{5, 3, 1, 6, 4, 2},
				left:  0,
				right: 5,
				want:  []int{1, 2, 3, 4, 5, 6},
			},
		},
		{
			name: "Should quick sort an array2",
			args: args{
				array: []int{2, 3, 1, 2, 4, 2},
				left:  0,
				right: 5,
				want:  []int{1, 2, 2, 2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.array, tt.args.left, tt.args.right)
			fmt.Println(tt.args.array)
			if !reflect.DeepEqual(tt.args.array, tt.args.want) {
				t.Errorf("quicksort()=%v, want: %v", tt.args.array, tt.args.want)
			}
		})
	}
}

func Test_findKthLargest(t *testing.T) {
	type args struct {
		array []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5",
			args: args{
				array: []int{5, 3, 1, 6, 4, 2},
				k:     2,
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				array: []int{2, 3, 1, 2, 4, 2},
				k:     4,
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				array: []int{3},
				k:     1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.array, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findKthLargestQuickSelect(t *testing.T) {
	type args struct {
		array []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5",
			args: args{
				array: []int{5, 3, 1, 6, 4, 2},
				k:     2,
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				array: []int{2, 3, 1, 2, 4, 2},
				k:     4,
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				array: []int{3},
				k:     1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargestQuickSelect(tt.args.array, tt.args.k); got != tt.want {
				t.Errorf("findKthLargestQuickSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}

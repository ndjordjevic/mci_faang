package q14_start_end_of_target

import (
	"reflect"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	type args struct {
		array        []int
		l, r, target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should found at index 3",
			args: args{
				array:  []int{1, 2, 3, 4, 5, 6, 7, 8},
				l:      0,
				r:      7,
				target: 4,
			},
			want: 3,
		},
		{
			name: "should found at index 3, odd no of els",
			args: args{
				array:  []int{1, 2, 3, 4, 5, 6, 7},
				l:      0,
				r:      6,
				target: 4,
			},
			want: 3,
		},
		{
			name: "should return -1",
			args: args{
				array:  []int{1, 2, 3, 4, 5, 6, 7},
				l:      0,
				r:      6,
				target: 10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.array, tt.args.l, tt.args.r, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Should return a range",
			args: args{
				nums:   []int{1, 3, 3, 5, 5, 5, 8, 9},
				target: 5,
			},
			want: []int{3, 5},
		},
		{
			name: "Should return a range2",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6},
				target: 4,
			},
			want: []int{3, 3},
		},
		{
			name: "Should not return a range",
			args: args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 9,
			},
			want: []int{-1, -1},
		},
		{
			name: "Should not return a range2",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: []int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

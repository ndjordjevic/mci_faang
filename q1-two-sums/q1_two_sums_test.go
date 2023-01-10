package q1_two_sums

import (
	"reflect"
	"testing"
)

func Test_twoSumsBruteForce(t *testing.T) {
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
			name: "Best case",
			args: args{
				nums:   []int{1, 3, 7, 9, 2},
				target: 11,
			},
			want: []int{3, 4},
		},
		{
			name: "Nothing found",
			args: args{
				nums:   []int{1, 3, 7, 9, 2},
				target: 25,
			},
			want: nil,
		},
		{
			name: "Nothing found empty array",
			args: args{
				nums:   []int{},
				target: 1,
			},
			want: nil,
		},
		{
			name: "Nothing found single array element",
			args: args{
				nums:   []int{5},
				target: 5,
			},
			want: nil,
		},
		{
			name: "Target found two elems",
			args: args{
				nums:   []int{1, 6},
				target: 7,
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumsBruteForce(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumsBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSums(t *testing.T) {
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
			name: "Best case",
			args: args{
				nums:   []int{1, 3, 7, 9, 2},
				target: 11,
			},
			want: []int{3, 4},
		},
		{
			name: "Nothing found",
			args: args{
				nums:   []int{1, 3, 7, 9, 2},
				target: 25,
			},
			want: nil,
		},
		{
			name: "Nothing found empty array",
			args: args{
				nums:   []int{},
				target: 1,
			},
			want: nil,
		},
		{
			name: "Nothing found single array element",
			args: args{
				nums:   []int{5},
				target: 5,
			},
			want: nil,
		},
		{
			name: "Target found two elems",
			args: args{
				nums:   []int{1, 6},
				target: 7,
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSums(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSums() = %v, want %v", got, tt.want)
			}
		})
	}
}

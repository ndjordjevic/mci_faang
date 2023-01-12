package q2_container_with_most_water

import "testing"

func Test_maxWaterContainerBruteForce(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should return 28",
			args: args{
				[]int{7, 1, 2, 3, 9},
			},
			want: 28,
		},
		{
			name: "Should return 0 empty array",
			args: args{
				[]int{},
			},
			want: 0,
		},
		{
			name: "Should return 0 one el only",
			args: args{
				[]int{7},
			},
			want: 0,
		},
		{
			name: "Should return 32",
			args: args{
				[]int{6, 9, 3, 4, 5, 8},
			},
			want: 32,
		},
		{
			name: "Should return another 32",
			args: args{
				[]int{4, 8, 1, 2, 3, 9},
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWaterContainerBruteForce(tt.args.heights); got != tt.want {
				t.Errorf("maxWaterContainerBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxWaterContainer(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should return 28",
			args: args{
				[]int{7, 1, 2, 3, 9},
			},
			want: 28,
		},
		{
			name: "Should return 0 empty array",
			args: args{
				[]int{},
			},
			want: 0,
		},
		{
			name: "Should return 0 one el only",
			args: args{
				[]int{7},
			},
			want: 0,
		},
		{
			name: "Should return 32",
			args: args{
				[]int{6, 9, 3, 4, 5, 8},
			},
			want: 32,
		},
		{
			name: "Should return another 32",
			args: args{
				[]int{4, 8, 1, 2, 3, 9},
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWaterContainer(tt.args.heights); got != tt.want {
				t.Errorf("maxWaterContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

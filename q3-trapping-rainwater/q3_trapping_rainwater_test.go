package q3_trapping_rainwater

import "testing"

func Test_trappingRainWaterBruteForce(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should return 8",
			args: args{
				heights: []int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2},
			},
			want: 8,
		},
		{
			name: "Should return 0",
			args: args{
				heights: []int{},
			},
			want: 0,
		},
		{
			name: "Should return 0, one el",
			args: args{
				heights: []int{3},
			},
			want: 0,
		},
		{
			name: "Should return 0, no container",
			args: args{
				heights: []int{3, 4, 3},
			},
			want: 0,
		},
		{
			name: "Should return 20",
			args: args{
				heights: []int{5, 0, 3, 0, 0, 0, 2, 3, 4, 2, 1},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trappingRainWaterBruteForce(tt.args.heights); got != tt.want {
				t.Errorf("trappingRainWaterBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trappingRainWaterBrute(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should return 8",
			args: args{
				heights: []int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2},
			},
			want: 8,
		},
		{
			name: "Should return 0",
			args: args{
				heights: []int{},
			},
			want: 0,
		},
		{
			name: "Should return 0, one el",
			args: args{
				heights: []int{3},
			},
			want: 0,
		},
		{
			name: "Should return 0, no container",
			args: args{
				heights: []int{3, 4, 3},
			},
			want: 0,
		},
		{
			name: "Should return 20",
			args: args{
				heights: []int{5, 0, 3, 0, 0, 0, 2, 3, 4, 2, 1},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trappingRainWaterBrute(tt.args.heights); got != tt.want {
				t.Errorf("trappingRainWaterBrute() = %v, want %v", got, tt.want)
			}
		})
	}
}

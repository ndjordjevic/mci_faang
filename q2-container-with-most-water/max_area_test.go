package q2_container_with_most_water

import "testing"

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
			name: "Should return correct result of 32",
			args: args{
				[]int{4, 8, 1, 2, 3, 9},
			},
			want: 32,
		},
		{
			name: "Should return 0",
			args: args{
				[]int{4},
			},
			want: 0,
		},
		{
			name: "Should return 0 empty array",
			args: args{
				[]int{},
			},
			want: 0,
		},
		{
			name: "Should return 0 empty array",
			args: args{
				[]int{2, 4},
			},
			want: 2,
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

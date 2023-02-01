package q5_longest_substring

import "testing"

func Test_longestSubstringBruteForce(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should return 3",
			args: args{"abccabb"},
			want: 3,
		},
		{
			name: "Should return 4",
			args: args{"abcbdca"},
			want: 4,
		},
		{
			name: "Should return 1",
			args: args{"cccccc"},
			want: 1,
		},
		{
			name: "Should return 0",
			args: args{""},
			want: 0,
		},
		{
			name: "Should return 4",
			args: args{"abcbda"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubstringBruteForce(tt.args.s); got != tt.want {
				t.Errorf("longestSubstringBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should return 3",
			args: args{"abccabb"},
			want: 3,
		},
		{
			name: "Should return 4",
			args: args{"abcbdca"},
			want: 4,
		},
		{
			name: "Should return 1",
			args: args{"cccccc"},
			want: 1,
		},
		{
			name: "Should return 0",
			args: args{""},
			want: 0,
		},
		{
			name: "Should return 4",
			args: args{"abcbda"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubstring(tt.args.s); got != tt.want {
				t.Errorf("longestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

package q10_stacks_valid_parentheses

import "testing"

func Test_isValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should be valid",
			args: args{"{([])}"},
			want: true,
		},
		{
			name: "Should not be valid",
			args: args{"{([]"},
			want: false,
		},
		{
			name: "Should not be valid2",
			args: args{"{[(])}"},
			want: false,
		},
		{
			name: "Should be valid2",
			args: args{"{[]()}"},
			want: true,
		},
		{
			name: "Should not be valid3",
			args: args{"]"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("isValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidParenthesesRuneVer(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should be valid",
			args: args{"{([])}"},
			want: true,
		},
		{
			name: "Should not be valid",
			args: args{"{([]"},
			want: false,
		},
		{
			name: "Should not be valid2",
			args: args{"{[(])}"},
			want: false,
		},
		{
			name: "Should be valid2",
			args: args{"{[]()}"},
			want: true,
		},
		{
			name: "Should not be valid3",
			args: args{"]"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidParenthesesRuneVer(tt.args.s); got != tt.want {
				t.Errorf("isValidParenthesesRuneVer() = %v, want %v", got, tt.want)
			}
		})
	}
}

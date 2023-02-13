package q6_valid_almost_valid_palindrome

import "testing"

func Test_clearString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should return valid cleared string",
			args: args{"A man, a plan, a canal: Panama"},
			want: "amanaplanacanalpanama",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clearStringAndLowercase(tt.args.s); got != tt.want {
				t.Errorf("clearString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidPalindromeTwoPointersFromOutside(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid",
			args: args{"aabaa"},
			want: true,
		},
		{
			name: "Valid2",
			args: args{"aabbaa"},
			want: true,
		},
		{
			name: "Invalid",
			args: args{"abc"},
			want: false,
		},
		{
			name: "Valid3",
			args: args{"a"},
			want: true,
		},
		{
			name: "Valid4",
			args: args{""},
			want: true,
		},
		{
			name: "Valid5",
			args: args{"A man, a plan, a canal: Panama"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPalindromeTwoPointersFromOutside(tt.args.s); got != tt.want {
				t.Errorf("isValidPalindromeTwoPointersFromOutside() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidPalindromeFromCenter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "Valid",
			args: args{"aabaa"},
			want: true,
		},
		{
			name: "Valid2",
			args: args{"aabbaa"},
			want: true,
		},
		{
			name: "Invalid",
			args: args{"abc"},
			want: false,
		},
		{
			name: "Valid3",
			args: args{"a"},
			want: true,
		},
		{
			name: "Valid4",
			args: args{""},
			want: true,
		},
		{
			name: "Valid5",
			args: args{"A man, a plan, a canal: Panama"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPalindromeFromCenter(tt.args.s); got != tt.want {
				t.Errorf("isValidPalindromeFromCenter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidPalindromeCompareAgainstReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "Valid",
			args: args{"aabaa"},
			want: true,
		},
		{
			name: "Valid2",
			args: args{"aabbaa"},
			want: true,
		},
		{
			name: "Invalid",
			args: args{"abc"},
			want: false,
		},
		{
			name: "Valid3",
			args: args{"a"},
			want: true,
		},
		{
			name: "Valid4",
			args: args{""},
			want: true,
		},
		{
			name: "Valid5",
			args: args{"A man, a plan, a canal: Panama"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPalindromeCompareAgainstReverse(tt.args.s); got != tt.want {
				t.Errorf("isValidPalindromeCompareAgainstReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_almostPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "Valid",
			args: args{"raceacar"},
			want: true,
		},
		{
			name: "Valid2",
			args: args{"abccdba"},
			want: true,
		},
		{
			name: "Invalid",
			args: args{"abcdefdba"},
			want: false,
		},
		{
			name: "Valid3",
			args: args{"a"},
			want: true,
		},
		{
			name: "Valid4",
			args: args{""},
			want: true,
		},
		{
			name: "Valid5",
			args: args{"ab"},
			want: true,
		},
		{
			name: "Valid6",
			args: args{"aca"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := almostPalindrome(tt.args.s); got != tt.want {
				t.Errorf("almostPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

package q4_backspace_compare

import "testing"

func Test_typedOutStringsBruteForce(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Best case",
			args: args{
				s: "ab#z",
				t: "az#z",
			},
			want: true,
		},
		{
			name: "Should return false",
			args: args{
				s: "abc#d",
				t: "acc#c",
			},
			want: false,
		},
		{
			name: "Should return true, both strings empty",
			args: args{
				s: "x#y#z#",
				t: "a#",
			},
			want: true,
		},
		{
			name: "Multiple #",
			args: args{
				s: "a###b",
				t: "b",
			},
			want: true,
		},
		{
			name: "Case sensitive, not equal",
			args: args{
				s: "Ab#z",
				t: "ab#z",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompareBruteForce(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompareBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_typedOutStrings(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Best case",
			args: args{
				s: "ab#z",
				t: "az#z",
			},
			want: true,
		},
		{
			name: "Should return false",
			args: args{
				s: "abc#d",
				t: "acc#c",
			},
			want: false,
		},
		{
			name: "Should return true, both strings empty",
			args: args{
				s: "x#y#z#",
				t: "a#",
			},
			want: true,
		},
		{
			name: "Multiple #",
			args: args{
				s: "a###b",
				t: "b",
			},
			want: true,
		},
		{
			name: "Case sensitive, not equal",
			args: args{
				s: "Ab#z",
				t: "ab#z",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: "nzp#o#g",
				t: "b#nzp#o#g",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cleanBS(t *testing.T) {
	type args struct {
		s string
		p int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Return 3",
			args: args{
				s: "ab#z",
				p: 3,
			},
			want: 3,
		},
		{
			name: "Return 0",
			args: args{
				s: "ab#z",
				p: 2,
			},
			want: 0,
		},
		{
			name: "Return 0",
			args: args{
				s: "ab#z",
				p: 0,
			},
			want: 0,
		},
		{
			name: "Return -1",
			args: args{
				s: "###",
				p: 2,
			},
			want: -1,
		},
		{
			name: "Return 0",
			args: args{
				s: "abc##",
				p: 4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanBS(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("cleanBS() = %v, want %v", got, tt.want)
			}
		})
	}
}

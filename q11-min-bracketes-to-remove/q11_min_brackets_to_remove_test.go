package q11_min_bracketes_to_remove

import "testing"

func Test_minRemoveToMakeValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Valid1",
			args: args{
				s: "a)bc(d)",
			},
			want: "abc(d)",
		},
		{
			name: "Valid2",
			args: args{
				s: "(ab(c)d",
			},
			want: "ab(c)d",
		},
		{
			name: "Valid3",
			args: args{
				s: "))((",
			},
			want: "",
		},
		{
			name: "Valid4",
			args: args{
				s: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRemoveToMakeValid(tt.args.s); got != tt.want {
				t.Errorf("minRemoveToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minRemoveToMakeValidLeetCode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Valid1",
			args: args{
				s: "a)bc(d)",
			},
			want: "abc(d)",
		},
		{
			name: "Valid2",
			args: args{
				s: "(ab(c)d",
			},
			want: "ab(c)d",
		},
		{
			name: "Valid3",
			args: args{
				s: "))((",
			},
			want: "",
		},
		{
			name: "Valid4",
			args: args{
				s: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRemoveToMakeValidLeetCode(tt.args.s); got != tt.want {
				t.Errorf("minRemoveToMakeValidLeetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

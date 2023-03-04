package q7_reverse_linked_list

import (
	"reflect"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	type args struct {
		head *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "Should reverse",
			args: args{
				&node{
					value: 1,
					next: &node{
						value: 2,
						next: &node{
							value: 3,
							next: &node{
								value: 4,
								next: &node{
									value: 5,
									next:  nil,
								},
							},
						},
					},
				},
			},
			want: &node{
				value: 5,
				next: &node{
					value: 4,
					next: &node{
						value: 3,
						next: &node{
							value: 2,
							next: &node{
								value: 1,
								next:  nil,
							},
						},
					},
				},
			},
		},
		{
			name: "Only one element",
			args: args{
				&node{3, nil},
			},
			want: &node{3, nil},
		},
		{
			name: "nil linkedliost",
			args: args{
				nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLinkedList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

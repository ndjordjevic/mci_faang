package q7_reverse_linked_list

import (
	"reflect"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Should reverse",
			args: args{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			name: "Only one element",
			args: args{
				&ListNode{3, nil},
			},
			want: &ListNode{3, nil},
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

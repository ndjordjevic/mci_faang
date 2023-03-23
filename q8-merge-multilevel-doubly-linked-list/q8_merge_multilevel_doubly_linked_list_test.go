package q8_merge_multilevel_doubly_linked_list

import (
	"testing"
)

type List struct {
	head   *Node
	curr   *Node
	levels []*Node
}

func (l *List) Insert(val int) {
	n := &Node{
		Val:   val,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}

	if l.curr == nil {
		l.head = n
		l.curr = n
		return
	}

	n.Prev = l.curr
	l.curr.Next = n
	l.curr = n
}

func (l *List) InsertChild(val int) {
	n := &Node{
		Val:   val,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	l.curr.Child = n
	l.levels = append(l.levels, l.curr)
	l.curr = n
}

func (l *List) LevelUp() {
	last := l.levels[len(l.levels)-1]
	l.curr = last
	l.levels = l.levels[:len(l.levels)-1]
}

func Test_flatten(t *testing.T) {
	l := &List{}
	l.Insert(1)
	l.Insert(2)
	l.InsertChild(7)
	l.Insert(8)
	l.InsertChild(10)
	l.Insert(11)
	l.LevelUp()
	l.Insert(9)
	l.LevelUp()
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)
	l.InsertChild(12)
	l.Insert(13)
	l.LevelUp()
	l.Insert(6)

	lw := &List{}
	lw.Insert(1)
	lw.Insert(2)
	lw.Insert(7)
	lw.Insert(8)
	lw.Insert(10)
	lw.Insert(11)
	lw.Insert(9)
	lw.Insert(3)
	lw.Insert(4)
	lw.Insert(5)
	lw.Insert(12)
	lw.Insert(13)
	lw.Insert(6)

	fhead := flatten(l.head)
	if !compareLL(fhead, lw.head) {
		t.Errorf("compareLL() returned false")
	}
}

func compareLL(head1, head2 *Node) bool {
	if head1 == nil && head2 == nil {
		return true
	}

	if head1 == nil && head2 != nil || head1 != nil && head2 == nil {
		return false
	}

	for head1 != nil || head2 != nil {
		if head1 == nil && head2 != nil || head1 != nil && head2 == nil || head1.Val != head2.Val {
			return false
		}

		head1 = head1.Next
		head2 = head2.Next
	}

	return true
}

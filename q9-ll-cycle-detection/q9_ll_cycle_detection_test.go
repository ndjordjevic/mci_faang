package q9_ll_cycle_detection

import (
	"reflect"
	"testing"
)

type CycledLL struct {
	head *ListNode
	cn   *ListNode
}

func (cll *CycledLL) Insert(v int, saveCN bool) {
	n := &ListNode{
		Val:  v,
		Next: nil,
	}
	if cll.head == nil {
		cll.head = n
		return
	}

	t := cll.head
	for t.Next != nil {
		t = t.Next
	}
	t.Next = n

	if saveCN {
		cll.cn = n
	}
}

func (cll *CycledLL) InsertToCycle(v int) {
	n := &ListNode{
		Val:  v,
		Next: cll.cn,
	}

	t := cll.head
	for t.Next != nil {
		t = t.Next
	}
	t.Next = n
}

func Test_detectCycleBruteForce_cycleExists(t *testing.T) {
	cll := CycledLL{}
	cll.Insert(1, false)
	cll.Insert(2, false)
	cll.Insert(3, true)
	cll.Insert(4, false)
	cll.Insert(5, false)
	cll.Insert(6, false)
	cll.Insert(7, false)
	cll.InsertToCycle(8)

	want := cll.cn

	if got := detectCycleBruteForce(cll.head); !reflect.DeepEqual(got, want) {
		t.Errorf("detectCycleBruteForce() = %v, want %v", got, want)
	}
}

func Test_detectCycleBruteForce_cycleNotExists(t *testing.T) {
	cll := CycledLL{}
	cll.Insert(1, false)
	cll.Insert(2, false)
	cll.Insert(3, false)
	cll.Insert(4, false)
	cll.Insert(5, false)
	cll.Insert(6, false)
	cll.Insert(7, false)

	want := cll.cn

	if got := detectCycleBruteForce(cll.head); !reflect.DeepEqual(got, want) {
		t.Errorf("detectCycleBruteForce() = %v, want %v", got, want)
	}
}

func Test_detectCycleFloydTortoiseHair_cycleExists(t *testing.T) {
	cll := CycledLL{}
	cll.Insert(1, false)
	cll.Insert(2, false)
	cll.Insert(3, true)
	cll.Insert(4, false)
	cll.Insert(5, false)
	cll.Insert(6, false)
	cll.Insert(7, false)
	cll.InsertToCycle(8)

	want := cll.cn

	if got := detectCycleFloydTortoiseHair(cll.head); !reflect.DeepEqual(got, want) {
		t.Errorf("detectCycleBruteForce() = %v, want %v", got, want)
	}
}

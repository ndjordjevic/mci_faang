package q12_implement_queue_with_stacks

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack{}
	s.Push(3)
	s.Push(5)
	s.Push(7)

	got := s.Pop()
	want := 7
	if got != want {
		t.Errorf("s.Pop()=%v, but want: %v", got, want)
	}

	gotIsEmpty := s.IsEmpty()
	if gotIsEmpty != false {
		t.Errorf("s.IsEmpty()=%v, but want: %v", gotIsEmpty, false)
	}

	gotPeek := s.Peek()
	wantPeek := 5
	if gotPeek != wantPeek {
		t.Errorf("s.Peek()=%v, but want: %v", gotPeek, wantPeek)
	}
}

func TestMyQueue(t *testing.T) {
	q := Constructor()
	q.Push(3)
	q.Push(5)
	q.Push(7)

	gotPopped := q.Pop()
	wantPopped := 3
	if gotPopped != wantPopped {
		t.Errorf("q.Pop()=%v, but want: %v", gotPopped, wantPopped)
	}

	gotPeek := q.Peek()
	wantPeek := 5
	if gotPeek != wantPeek {
		t.Errorf("q.Peek()=%v, but want: %v", gotPeek, wantPeek)
	}

	gotIsEmpty := q.Empty()
	if gotIsEmpty != false {
		t.Errorf("q.Empty()=%v, but want: %v", gotIsEmpty, false)
	}
}

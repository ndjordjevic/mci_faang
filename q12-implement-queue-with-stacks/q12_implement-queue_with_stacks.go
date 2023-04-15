package q12_implement_queue_with_stacks

// leetcode #232
// Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

// Queue operations should be implemented with stacks (one or more stacks)
// We need two stacks, In and Out
// Enqueue (Push): will push the value to the In stack
// Dequeue (Pop): will check If Out queue has any value inside, and if Not it will While loop and Pop the value from the In queue and Push it to the Out queue.
// Then it will Pop the value from the Out queue
// Peek: Will do the similar to the Dequeue but will not Pop it, it will get the last element from the Out queue
// Empty: will return if both length of In and Out queues are == 0

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(i int) {
	*s = append(*s, i) // Simply append the new value to the end of the stack
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element
	}
}

func (s *Stack) Peek() int {
	index := len(*s) - 1 // Get the index of the top most element.
	element := (*s)[index]
	return element
}

type MyQueue struct { // Space: O(n)
	in  Stack
	out Stack
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) { // Time: O(1)
	q.in.Push(x)
}

func (q *MyQueue) Pop() int { // Time: O(n)
	if q.out.IsEmpty() {
		for !q.in.IsEmpty() {
			q.out.Push(q.in.Pop())
		}
	}
	return q.out.Pop()
}

func (q *MyQueue) Peek() int { // Time: O(n)
	if q.out.IsEmpty() {
		for !q.in.IsEmpty() {
			q.out.Push(q.in.Pop())
		}
	}
	return q.out.Peek()
}

func (q *MyQueue) Empty() bool { // Time: O(1)
	return q.in.IsEmpty() && q.out.IsEmpty()
}

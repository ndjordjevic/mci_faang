package q10_stacks_valid_parentheses

// leetcode #20
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
//
// 1. Constraints: Empty string is valid, string consists of parentheses only '()[]{}'
// 2. Test cases:
// "": true
// "{([])}": true
// "{([]": false
// "{[(])}": false
// "{[]()}": true

// Create a map of valid parentheses pairs, key left, value right parentheses
// If len of s is 0 return true
// Iterate over s characters and check from the map if the char is a left bracket
// If yes push to the stack
// Else, we know it must be a right bracket, and we pop out the latest value from the stack, and get the correct right bracket from the map, key left bracket
// then we compare the current bracket and the correct, and if they are != return false
// close the loop
// return if the stack is empty (return true), if it's not return false (there are still some brackets left in it, without matching right ones)
func isValidParentheses(s string) bool { // Time: O(n), Space: O(n)
	var stack []string
	m := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	if len(s) == 0 {
		return false
	}

	for _, b := range s {
		if _, ok := m[string(b)]; ok {
			stack = append(stack, string(b))
		} else {
			if len(stack) == 0 {
				return false
			}
			leftB := stack[len(stack)-1] // get the last element from stack
			stack = stack[:len(stack)-1] // pop the last element from stack
			correctB := m[leftB]
			if string(b) != correctB {
				return false
			}
		}
	}

	return len(stack) == 0
}

func isValidParenthesesRuneVer(s string) bool { // Time: O(n), Space: O(n)
	var stack []rune
	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	if len(s) == 0 {
		return false
	}

	for _, b := range s {
		if _, ok := m[(b)]; ok {
			stack = append(stack, b)
		} else {
			if len(stack) == 0 {
				return false
			}
			leftB := stack[len(stack)-1] // get the last element from stack
			stack = stack[:len(stack)-1] // pop the last element from stack
			correctB := m[leftB]
			if b != correctB {
				return false
			}
		}
	}

	return len(stack) == 0
}

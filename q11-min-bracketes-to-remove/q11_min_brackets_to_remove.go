package q11_min_bracketes_to_remove

import "strings"

// leetcode #1249
// Given a string s of '(' , ')' and lowercase English characters.
// Your task is to remove the minimum number of parentheses ( '(' or ')', in any positions ) so that the resulting parentheses string is valid and return any valid string.
// Formally, a parentheses string is valid if and only if:
// It is the empty string, contains only lowercase characters, or
// It can be written as AB (A concatenated with B), where A and B are valid strings, or
// It can be written as (A), where A is a valid string.
//
// Constraints: return valid string, no spaces in 's'. It contains only lowercase chars and two mentioned brackets
// also string is valid if there are no brackets at all, empty string is valid too
//
// Test cases:
// "a)bc(d)": "abc(d)"
// "(ab(c)d": "ab(c)d" or "(abc)d"
// "))((": ""
// "": ""

// Declare a result -> a valid string, which should be an array for easier manipulation. Declare a stack too
// Iterate over an array fori loop
// If a char is a left bracket push an index of that bracket to the stack
// If a char is a right bracket AND stack has values inside of it (it doesn't matter how many) then pop out and move on
// If a char is a right bracket set that char to an empty string
// End the Loop
// Then we need to do something with brackets that are left in the stack
// Iterate while the stack has some brackets left
// For each iteration we need to pop the index from the stack and set the char of that index in the array to an empty char
// Then after the loop is done we just convert the array to a string and return it
func minRemoveToMakeValid(s string) string { // Time: O(n) , Space: O(n)
	if len(s) == 0 {
		return ""
	}

	res := strings.Split(s, "")
	var stack []int

	for i := 0; i < len(res); i++ {
		if res[i] == "(" {
			stack = append(stack, i)
		} else if res[i] == ")" && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if res[i] == ")" {
			res[i] = ""
		}
	}

	for len(stack) > 0 {
		index := stack[len(stack)-1]
		res[index] = ""
		stack = stack[:len(stack)-1]
	}

	return strings.Join(res, "")
}

func minRemoveToMakeValidLeetCode(s string) string {
	var stack []int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			} else {
				s = s[:i] + s[i+1:]
				i--
			}
		}
	}

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		s = s[:pop] + s[pop+1:]
	}
	return s
}

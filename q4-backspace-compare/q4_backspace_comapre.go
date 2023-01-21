package q4_backspace_compare

// leetcode #844

// Given two strings S and T, return if they equal when they both typed out.
// Any "#" that appears in the string counts as a backspace
// 1. Constraints:
//		Two #'s next to each other: delete the two values before the first #
//		What happens to # where there's no character to remove: deletes nothing then
//		Are two empty strings equal to each other: Yes
// 		Does case sensitivity matters: Yes
// 2. Test cases:
// S:"ab#z" T:"az#z" true
// S:"abc#d" T:"acc#c" false
// S:"x#y#z#" T:"a#" true
// S:"a###b" T:"b" true
// S:"Ab#z" T:"ab#z" false

// Type out both strings and compare the results
// S:"ab#z" T:"az#z" true
func backspaceCompareBruteForce(s, t string) bool { // Time: O(n+m), Space: O(n+m)
	var r1, r2 string

	for _, c := range s {
		if c != '#' {
			r1 += string(c)
		} else {
			if len(r1) > 0 {
				r1 = r1[:len(r1)-1]
			}
		}
	}

	for _, c := range t {
		if c != '#' {
			r2 += string(c)
		} else {
			if len(r2) > 0 {
				r2 = r2[:len(r2)-1]
			}
		}
	}

	return len(r1) == len(r2) && r1 == r2
}

// Two pointers, each one on each string, starting from the last char in both strings
// while p1 >= 0 || p2 >= 0
// If the current char at p1 or p2 is # backspace till find the char to compare
// If p1 and p2 are > 0 and chars don't match return false
// If one of the p1 or p2 is -1 and the other is not return false
// If no false return is reached return true at the end
// S:"ab#z" T:"az#z" true
// S:"x#y#z#" T:"a#" true
// S:"a###b" T:"b" true
// S:"bxj##tw" T:"bxj###tw" false
func backspaceCompare(s, t string) bool { // Time: O(n+m), Space: O(1)
	p1, p2 := len(s)-1, len(t)-1 // p1=-1, p2=1

	for p1 >= 0 || p2 >= 0 {
		if p1 >= 0 && s[p1] == '#' {
			// clean s #
			p1 = backspace(s, p1)
		}

		if p2 >= 0 && t[p2] == '#' {
			// clean t #
			p2 = backspace(t, p2)
		}

		if p1 >= 0 && p2 >= 0 && s[p1] != t[p2] {
			return false
		}

		if (p1 >= 0) != (p2 >= 0) {
			return false
		}

		p1--
		p2--
	}

	return true
}

// s="ab#z" p=3
func backspace(s string, p int) int {
	bc := 2
	for p >= 0 && bc > 0 {
		p--
		bc--
		if p >= 0 && s[p] == '#' {
			bc += 2
		}
	}

	return p
}

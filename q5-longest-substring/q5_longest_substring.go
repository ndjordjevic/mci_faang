package q5_longest_substring

// leetcode #3
// Given a string s, find the length of the longest substring without repeating characters
//
// 1. Constraints:
// contiguous string, all lowercase
//
// 2. Test cases:
// "abccabb", 3
// "abcbdca", 4
// "cccccc", 1
// "", 0
// "abcbda", 4

// Generate every possible substring and compute the longest
//
// "abccabb", 3
func longestSubstringBruteForce(s string) int { // Time: O(n^2), Space: O(n)
	if len(s) <= 1 {
		return len(s)
	}

	longest := 0

	for l := 0; l < len(s); l++ {
		seenChars := make(map[uint8]struct{})
		currLength := 0
		for r := l; r < len(s); r++ {
			if _, ok := seenChars[s[r]]; ok {
				break
			}
			currLength++
			seenChars[s[r]] = struct{}{}
			if currLength > longest {
				longest = currLength
			}
		}
	}

	return longest
}

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

// Sliding window with two pointers
// still have a map, key: char, value: index
// initialize left and largest to 0
// iterate right from 0 to len
// check if current char at right exist in a map, and it's index is >=:
// yes: update left to be one more than the right (move it forward to be one in front of the right)
// update a map, key char at right, value current index
// if right - left + 1 > longest that's a new longest
//
// "abccabb", 3
func longestSubstring(s string) int { // Time:O(n), Space: O(n)
	if len(s) <= 1 {
		return len(s)
	}

	seenChars := make(map[uint8]int) // a:4, b:6, c:3
	left := 0                        // 6
	longest := 0                     // 3

	for right := 0; right < len(s); right++ { // 6
		currentChar := s[right]                          // 'b'
		previouslySeenChar, ok := seenChars[currentChar] // previouslySeenChar=5
		if ok && previouslySeenChar >= left {
			left = previouslySeenChar + 1
		}

		seenChars[currentChar] = right
		if (right - left + 1) > longest {
			longest = right - left + 1
		}
	}

	return longest
}

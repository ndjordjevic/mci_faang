package q6_valid_almost_valid_palindrome

import (
	"math"
	"regexp"
	"strings"
)

// leetcode #125
// Given a string, determine if it's a palindrome, considering alphanumeric chars and ignoring case sensitivity
// 1. Constraints are clear
// 2. Test cases:
// "aabaa" true
// "aabbaa" true
// "abc" false
// "a" true
// "" true
// "A man, a plan, a canal: Panama" true

// Two pointers from outside
// Initialize left/right pointers at start and end of string respectively
// Loop through string chars while comparing them and the move pointers closer to the center
// "A man, plan, a canal: Panama" true
func isValidPalindromeTwoPointersFromOutside(s string) bool {
	s = clearStringAndLowercase(s)
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// Two pointers from the center
// Initialize left/right pointers to point to the middle index of the string
// We need to floor the value from dividing length by 2 to get the index of the center
// If string is even move left back by 1
// Loop through the string and compare chars while expanding pointers outwards
// "A man, plan, a canal: Panama" true
func isValidPalindromeFromCenter(s string) bool {
	s = clearStringAndLowercase(s)
	l := int(math.Floor(float64(len(s) / 2)))
	r := l

	if len(s)%2 == 0 {
		l--
	}

	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			return false
		}
		l--
		r++
	}

	return true
}

// Compare against the reversed string
// "A man, plan, a canal: Panama" true
func isValidPalindromeCompareAgainstReverse(s string) bool {
	s = clearStringAndLowercase(s)

	rev := reverseString(s)

	return s == rev
}

// leetcode #680
// Given a string determine if it's almost a palindrome, and if it becomes a palindrome by removing one char
// 1. Constraints
// Only consider alphanumeric chars and ignore case sensitivity
// 2. Test cases
// "raceacar" true
// "abccdba" true
// "abcdefdba" false
// "" true
// "a" true
// "ab" true
// "aca" true
//
// Initialize left and right pointers, left to 0, right to length - 1
// Write validSubpalindrome func with an original string, left and right params that will check if the subpalindrome is valid
// while left < right check if left != right and
// yes: check if moved left or moved right by one is a valid subpalindrome, if yes it's almost a palindrome
func almostPalindrome(s string) bool { // Time: O(n), Space: O(1)
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return validSubpalindrome(s, left+1, right) || validSubpalindrome(s, left, right-1)
		}
		left++
		right--
	}

	return true
}

func validSubpalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func clearStringAndLowercase(s string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	return strings.ToLower(nonAlphanumericRegex.ReplaceAllString(s, ""))
}

// The most efficient way to reverse a string
func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

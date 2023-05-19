package q14_start_end_of_target

// initialize two pointers, l=0, r=len(array)-1
// iterate while l <= r
// for every iteration calc mid that would be (l+r)1/2
// define foundValue as array[mid]
// if foundValue == target return mid
// if foundValue < target move l to mid+1
// else move r to mid-1 and end the loop
// return -1 if target couldn't be found
func binarySearch(array []int, l, r, target int) int { // Time: O(logn), Space: O(1)
	for l <= r {
		mid := (l + r) / 2
		foundValue := array[mid]
		if foundValue == target {
			return mid
		} else if foundValue < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// leetcode #34
// If you see a sorted order probably you gonna use a binary search
// Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
// If target is not found in the array, return [-1, -1].
// You must write an algorithm with O(log n) runtime complexity.
// Constraints:
// If target is not found in the array, return [-1, -1]
//
// Test cases:
// [1,3,3,5,5,5,8,9], t=5 -> [3,5]
// [1,2,3,4,5,6], t=4 -> [3,3]
// [1,2,3,4,5], t=9 -> [-1,-1]
// [], t=9 -> [-1,-1]

// we should find the very firstPosition of the target  (not the beginning of the range, any position with this target)
// if it's -1 return [-1,-1]
// if it's not then we need to binarySearch on the left and on the right side from the firstPosition and continue BS until it gives us -1
// for each side (left and right) we need to track endPosition (will be overwritten with -1 at some point)  and another var called temp (keeps the last successful found target value)
// we have to do the same thing on the left side as well and need to keep track of the startPosition which would be init to firstPosition so as the endPosition
// while loop till startPosition!=-1 to search left side and end the loop
// then do the same for the right: while loop till endPosition!=-1 to search right side and end the loop
func searchRange(nums []int, target int) []int { // Space: O(1), Time: O(logn)
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	firstPosition := binarySearch(nums, 0, len(nums)-1, target)
	if firstPosition == -1 {
		return []int{-1, -1}
	}

	startPosition, endPosition := firstPosition, firstPosition
	var temp1, temp2 int

	for startPosition != -1 {
		temp1 = startPosition
		// search left side
		startPosition = binarySearch(nums, 0, startPosition-1, target)
	}
	startPosition = temp1

	for endPosition != -1 {
		temp2 = endPosition
		// search right side
		endPosition = binarySearch(nums, endPosition+1, len(nums)-1, target)
	}
	endPosition = temp2

	return []int{startPosition, endPosition}
}

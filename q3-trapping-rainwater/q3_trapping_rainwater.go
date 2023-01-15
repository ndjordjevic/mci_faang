package q3_trapping_rainwater

import "math"

// leetcode #42

// Given n non-negative integers representing an elevation map where the width of each bar is 1,
// compute how much water it is able to trap after raining.
//
// 1. Constraints:
//		left and right sides don't count as walls
//		all ints are positive
// 2. Test cases
// 	[0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2], 8
// 	[], 0
// 	[3], 0
//	[3, 4, 3], 0
//	[5, 0, 3, 0, 0, 0, 2, 3, 4, 2, 1], 20
// 3. Solution without a code and formula
// currentWater = min(maxL, maxR) - currentHeight

// Brute force:
// currentWater = min(maxL, maxR) - currentHeight
// Iterate over an array (p)
// 	calc maxLeft and maxRight from the p using two while loops
// 	calc currentWater using formula above
//  if currentWater > 0 add to totalWater
//
/*
[0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2], 8
                               p
*/
func trappingRainWaterBruteForce(heights []int) int { // Time: O(n^2), Space: O(1)
	var totalWater int // totalWater=8

	for p := 0; p < len(heights); p++ { // p=10, heights[p]=2
		leftP, rightP, maxLeft, maxRight := p, p, 0, 0 // leftP=10, rightP=10, maxLeft=3, maxRight=0

		for leftP >= 0 {
			maxLeft = int(math.Max(float64(maxLeft), float64(heights[leftP])))
			leftP--
		}

		for rightP < len(heights) {
			maxRight = int(math.Max(float64(maxRight), float64(heights[rightP])))
			rightP++
		}

		currentWater := int(math.Min(float64(maxLeft), float64(maxRight)) - float64(heights[p]))
		if currentWater > 0 {
			totalWater += currentWater
		}
	}

	return totalWater
}

// Optimized:
// Two pointers solution
// Steps:
// 0. While leftP < rightP
// 1. Identify th pointer with lesser or equal value, which side we will operate
// 2. Is this pointer value greater than or equal to the max on that side
//		yes: update the max on that side
//		no: get the water for that pointer and add to the total
// 3. Move the pointer inwards
// 4. Repeat this for the other pointer
/*
[0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2], 8
       l                       r
*/
func trappingRainWaterBrute(heights []int) int { // Time: O(n), Space: O(1)
	left := 0                 // left=2
	right := len(heights) - 1 // right=10
	leftMax := 0              // leftMax=1
	rightMax := 0             // rightMax=0
	total := 0                // total=1

	for left < right {
		if heights[left] <= heights[right] {
			// left side
			if heights[left] >= leftMax {
				leftMax = heights[left]
			} else {
				total += leftMax - heights[left]
			}
			left++
		} else {
			// right side
			if heights[right] >= rightMax {
				rightMax = heights[right]
			} else {
				total += rightMax - heights[right]
			}
			right--
		}
	}

	return total
}

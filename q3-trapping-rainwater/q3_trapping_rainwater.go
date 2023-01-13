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

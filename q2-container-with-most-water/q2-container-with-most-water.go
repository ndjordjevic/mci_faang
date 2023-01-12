package q2_container_with_most_water

import "math"

// leetcode #11

// You are given an array of positive ints where each int represents the height of a vertical line on a chart
// Find two lines which together with x-axis forms a container that would hold the grates amount of water.
// Return the area that would hold
//
// 1. Constraints:
//		thickness of the line doesn't affect the area
//		the sides can't be used to form a container
//		lines inside a container don't affect the area, you can ignore them
// 2. Test cases
//		[7, 1, 2, 3, 9], 28
// 		[], 0
//		[7], 0
//		[6, 9, 3, 4, 5, 8], 32
//		[4, 8, 1, 2, 3, 9], 32
// 3. Solution without code
//		max question
//		area = len * width; area = min(a, b) * (bi - ai)

// Brute force: nested loops
//
//			maxArea = 0
//			iterate over array (p1=0)
//	     		iterate over array (p2=p1+1)
//					check if min(heights[p1], heights[p2]) * (p2 - p1) > maxArea
//						yes: set new maxArea
//						no: continue
func maxWaterContainerBruteForce(heights []int) int { // Time: O(n^2), Space: O(1)
	maxArea := 0                           // maxArea=32
	for p1 := 0; p1 < len(heights); p1++ { // p1=1
		for p2 := p1 + 1; p2 < len(heights); p2++ { // p2=5
			height := int(math.Min(float64(heights[p1]), float64(heights[p2]))) // height=8
			width := p2 - p1                                                    // width=4
			area := height * width                                              // area=32
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

// Optimized:
// two pointers, shift/move smaller number (value) to see if we can get a bigger one next and a bigger area
// move pointers until they meet
/*
 max = 32
 0  1  2  3  4  5
[6, 9, 3, 4, 5, 8]
    a  b
area = min(a, b) * (bi - ai)
	        3          1 = 3
*/
func maxWaterContainer(heights []int) int { // Time: O(n), Space: O(1)
	p1, p2, maxArea := 0, len(heights)-1, 0 // p1 = 1, p2 = 4, maxArea = 32

	for p1 < p2 {
		height := int(math.Min(float64(heights[p1]), float64(heights[p2]))) // height = 8
		width := p2 - p1                                                    // width = 4
		area := height * width                                              // area = 32
		maxArea = int(math.Max(float64(maxArea), float64(area)))

		if heights[p1] <= heights[p2] {
			p1++
		} else {
			p2--
		}
	}

	return maxArea
}

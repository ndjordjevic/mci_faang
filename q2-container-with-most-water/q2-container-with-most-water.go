package q2_container_with_most_water

import "math"

// leetcode #11
func maxWaterContainer(heights []int) int {
	p1, p2, maxArea := 0, len(heights)-1, 0

	for p1 < p2 {
		height := int(math.Min(float64(heights[p1]), float64(heights[p2])))
		width := p2 - p1
		area := height * width
		maxArea = int(math.Max(float64(maxArea), float64(area)))

		if heights[p1] <= heights[p2] {
			p1++
		} else {
			p2++
		}
	}

	return maxArea
}

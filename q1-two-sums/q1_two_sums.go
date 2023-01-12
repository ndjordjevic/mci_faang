package q1_two_sums

// leetcode #1

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
//  1. Write and understand the problem and verify the constraints:
//     All nums positive
//     No duplicates in the array
//     No they may not be always a solution, return nil
//     You may assume that each input would have exactly one solution, and you may not use the same element twice.
//  2. Write func signature and test-cases (use unit-tests)
//  3. Figure out solution without code, write logical steps
//		once it's found return
//  4. Write brute force solution
//  5. Double-check for errors
//     Check for typos, closing our loops etc...
//  6. Test with test cases manually walking through the code execution. Test edge cases too!
//  7. Time and space complexity
//  8. Optimize solution and repeat testing and calc time and space capacity

// Brute force: nested loops
//
//	iterate over array (p1=0)
//	calc numberToFind := target - nums[p1]
//	iterate over array (p2=p1+1)
//	if numberToFind == nums[p2]
//		yes: return [p1, p2]
//		no: continue
//
// nums = [1, 3, 7, 9, 2], target = 11
// nums = [1, 3, 7, 9, 2], target = 25
// nums = [], target = 11
// nums = [1], target = 11
func twoSumsBruteForce(nums []int, target int) []int { // Time: O(n^2), Space: O(1)
	for p1 := 0; p1 < len(nums); p1++ { // p1 = 0
		numberToFind := target - nums[p1]        // numberToFind = 0
		for p2 := p1 + 1; p2 < len(nums); p2++ { // p2 = 4
			if numberToFind == nums[p2] {
				return []int{p1, p2}
			}
		}
	}
	return nil
}

// Optimized:
//
//	     iterate over an array (p as a pointer)
//			check if there is a [target-num] item in a map
//				yes: return idx, p
//				no:  store in a map: num as a key, p as a value (numsMap[num] = p), continue looping
//
//		nums = [1, 3, 7, 9, 2], target = 11
//		nums = [1, 3, 7, 9, 2], target = 25
//		nums = [], target = 11
//		nums = [1], target = 11
func twoSums(nums []int, target int) []int { // Time: O(n), Space: O(n)
	numsMap := make(map[int]int) // {1:0, 3:1, 7:2, 9:3}
	for p, num := range nums {   // p=4, num=2
		if idx, ok := numsMap[target-num]; ok {
			return []int{idx, p} // [3, 4]
		}
		numsMap[num] = p
	}
	return nil
}

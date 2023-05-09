package q13_kth_largest_element_array

// leetcode #215
// Given an integer array nums and an integer k, return the kth largest element in the array.
// Note that it is the kth largest element in the sorted order, not the kth distinct element.
// You must solve it in O(n) time complexity.
//
// Constraints: k should be in the array boundaries
//
// Test cases:
// nums=[5,3,1,6,4,2], k=2: [1,2,3,4,5,6], 5
// nums=[2,3,1,2,4,2], k=4: [1,2,2,2,3,4], 2
// nums=[3], k=1: [3], 3

// Use quicksort to sort an array and then find the kth largest element
// Quicksort: sort array in place, uses recursion.
// Set a pivot number (last el in the array). We should try to find the final resting place for the pivot when the array is sorted
// All the elements on the left when pivot is in the right spot are less than pivot, on the right are greater than pivot, and they don't have to be sorted
// Then we have to sort the left and the right portion from the pivot using the same quicksort algorithm sorting one element at the time
// How to sort the pivot: instantiate two vars i (partitionIndex) and j on the first el on the left.
// i points where the final resting place for the pivot should be. j will scan through the array and compare the el at j with pivot
// if the el at j is < than the pivot we are going to swap el at j with el at i and move i over, j too. if it's >= don't do anything, move j over
// once the j hits pivot swap pivot with el at i at el at pivot gets to his final place... all els at left are <, and all els at right are >
// then we do the same for the els at left and with els at right from the 'Set a pivot number'
func quickSort(array []int, left, right int) { // Time: O(nlogn), Space: O(logn) - cause of the recursive calls
	if left < right {
		partitionIndex := partition(array, left, right)
		quickSort(array, left, partitionIndex-1)
		quickSort(array, partitionIndex+1, right)
	}
}

func partition(array []int, left, right int) int {
	pivotElement := array[right]
	partitionIndex := left

	for j := left; j < right; j++ {
		if array[j] < pivotElement {
			array[partitionIndex], array[j] = array[j], array[partitionIndex]
			partitionIndex++
		}
	}

	array[partitionIndex], array[right] = array[right], array[partitionIndex]

	return partitionIndex
}

func findKthLargest(array []int, k int) int {
	indexToFind := len(array) - k
	quickSort(array, 0, len(array)-1)

	return array[indexToFind]
}

func findKthLargestQuickSelect(array []int, k int) int {
	indexToFind := len(array) - k
	return quickSelect(array, 0, len(array)-1, indexToFind)
}

// QuickSelect: similar to quicksort
func quickSelect(array []int, left, right int, idxToFind int) int { // Time: O(n)-O(n^2), Space: O(1)
	partitionIndex := partition(array, left, right)
	if partitionIndex == idxToFind {
		return array[partitionIndex]
	}
	if idxToFind < partitionIndex {
		return quickSelect(array, left, partitionIndex-1, idxToFind)
	}

	return quickSelect(array, partitionIndex+1, right, idxToFind)
}

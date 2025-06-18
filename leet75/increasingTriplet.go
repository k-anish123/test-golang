package leet75
/**
Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

 

Example 1:

Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.
Example 2:

Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.
Example 3:

Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
*/
// This is a brute force solution with O(n^3) time complexity.
func increasingTriplet(nums []int) bool {
    
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] < nums[j] && nums[j] < nums[k] {
					return true
				}
			}
		}
	}
	return false
}


// A more efficient solution would be to use two variables to track the smallest and second smallest elements seen so far, and check if a third element is greater than the second smallest.
func IncreasingTriplet(nums []int) bool {	
	first, second := 1<<31-1, 1<<31-1 // Initialize to max int value
	for _, num := range nums {
		if num <= first {
			first = num // Update first smallest
		} else if num <= second {
			second = num // Update second smallest
		} else {
			return true // Found a third element greater than both
		}
	}
	return false // No increasing triplet found
}

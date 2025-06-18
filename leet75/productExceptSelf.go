package leet75

/**
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.



Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
*/
func ProductExceptSelf(nums []int) []int {
	zeroCount := 0
	product := 1
	positiveProduct := 1
	for _, num := range nums {
		if num == 0 {
			zeroCount++
			continue
		}
		product *= num
		if num > 0 {
			positiveProduct *= num
		}
	}
	result := make([]int, len(nums))
	if zeroCount > 1 {
		// If there are more than one zero, all products will be zero
		return result
	}
	for i, num := range nums {
		 if zeroCount == 1 {
			if num == 0 {
				result[i] = product
			} else {
				result[i] = 0
			}
		} else {
			result[i] = product / num
		}
	}

	return result

    
}

// func productExceptSelf(nums []int) []int {
//     res := make([]int, len(nums))

//     pref := 1
//     for i, v := range(nums) {
//         res[i] = pref
//         pref *= v
//     }

//     suf := 1
//     for i := len(nums) - 1; i >= 0; i-- {
//         res[i] *= suf
//         suf *= nums[i]
//     }

//     return res
// }
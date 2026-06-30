func findMin(nums []int) int {
	/*
	Find Minimum in Rotated Sorted Array

	two pointers, lower and upper
	lower should always be less than upper
	since the array was sorted, but rotated, and find the point where lower is greater than upper
	
	
	*/

	// we are searching for where nums[i] < nums[i+1]
	// nums[i+1] will be the minimum. if not found, then nums[0] is the minimum

	lower := 0
	upper := len(nums) - 1

	if nums[lower] <= nums[upper] {
		return nums[lower]
	}
	

	for upper - lower > 1 {
		mid := lower + (upper - lower) / 2
		// if mid is greater than lower and upper, then it's part of the first half 
		if nums[mid] < nums[lower] {
			upper = mid
		} else {
			lower = mid
		}
	}

	return nums[upper]
}

// [3,4,5,6,1,2]
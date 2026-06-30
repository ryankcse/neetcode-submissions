func search(nums []int, target int) int {
	
	// find where the deflection cut starts
	cut := 0
	// start with two pointers
	lower := 0
	upper := len(nums) - 1
	
	for lower < upper {
		cut = lower + (upper - lower)/2 
		if nums[cut] > nums[upper] { // if cut is greater than upper, then move lower
			lower = cut + 1
		} else {
			upper = cut
		}
	}

	cut = lower

	if target < nums[0] {
		lower = cut
		upper = len(nums) - 1
	} else if target > nums[len(nums) - 1] {
		lower = 0
		upper = cut - 1
	} else {
		lower = 0
		upper = len(nums) - 1
	}

	fmt.Println(lower, upper)

	mid := lower + (upper - lower)/2

	// binary search in sub array
	for lower < upper {
		mid = lower + (upper - lower)/2
		if nums[mid] < target {
			lower = mid + 1
		} else if nums[mid] > target {
			upper = mid - 1
		} else {
			return mid
		}
	}
	// lower == upper
	if nums[lower] == target {
		return lower
	} else {
		return -1
	}

}

func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	for left <= right {
		i := left + (right - left)/ 2
		if target < nums[i] {
			right = i - 1
		} else if target > nums[i] {
			left = i + 1
		} else {
			if nums[i] == target {
				return i
			} else {
				return -1
			}
		}
	}

	return -1
}

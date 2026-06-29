func twoSum(numbers []int, target int) []int {
	var answer []int

	for firstIdx := 0; firstIdx < len(numbers) - 1; firstIdx++ {
		for secondIdx := firstIdx + 1; secondIdx < len(numbers); secondIdx++ {
			if numbers[firstIdx] + numbers[secondIdx] == target {
				answer = []int{firstIdx + 1, secondIdx + 1}
			}
		}
	}

	return answer
}



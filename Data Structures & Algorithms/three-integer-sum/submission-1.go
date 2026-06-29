func threeSum(nums []int) [][]int {
	/*
	brute force approach: check all combinations of 3 indices
	- n^3 time complexity

	second approach: store the counts of each integer in a map.
	use 2 pointers, a first and second pointer, to point to an index in the array
	the 3rd value will be calculated by first - third, and we can check if the value exists in the map.
	time complexity: n^2
	*/
	var results [][]int
	resultsMap := make(map[string]bool)

	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}
	fmt.Println("counter: ", counter)
	fmt.Println("counter[0] ", counter[0] )


	for second := 1; second < len(nums) - 1; second ++ {
		for first := 0; first < second; first ++ {
			// remove the numbers that first and second are pointing to
			counter[nums[first]]--
			counter[nums[second]]--

			// Check if -(nums[first] + nums[second]) is present in the map
			thirdValue := -(nums[first] + nums[second])
			fmt.Println("counter[0] ", counter[0] )

			if counter[thirdValue] > 0 {
				// we found a triple whose sum = 0
				// let's sort the triplet to standardize it

				triplet := []int{nums[first], nums[second], thirdValue}
				fmt.Println("triplet pre sort: ", triplet)
				fmt.Println("thirdValue", thirdValue)
				fmt.Println("counter[thirdValue] ", counter[thirdValue] )
				sort.Ints(triplet)
				
				key := fmt.Sprintf("%d,%d,%d", triplet[0], triplet[1], triplet[2])
				fmt.Println(key, " ", second, nums[second])
				if !resultsMap[key] {
					resultsMap[key] = true
					// add triplet to results only if the triplet doesn't already exist
					results = append(results, triplet)
				}
			}

			// add them back in
			counter[nums[first]]++
			counter[nums[second]]++
		}
	}
	return results
}

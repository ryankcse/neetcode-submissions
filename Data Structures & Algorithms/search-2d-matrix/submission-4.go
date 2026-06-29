func searchMatrix(matrix [][]int, target int) bool {
	// idea: binary search - even though it's a 2d array of m x n, can we treat it as a 1 D array?
	// e.g. if the 2d index is 2, 3, the 1d index would be 
	// it depends on the width of the first array, lets say the dimensions are 3 x 4
	// then the (2,3) would turn into 2(3) + 3 = 9 
	// but our expected is 11... why is there a discrepancy?
	// 
	// example: (2, 0) should be 8
	// 2 * 4 + 0 = 8
	
	// 3 * 4 + 0 = 12 which is out of bounds
	// 1 * 4 + 0 = 4
	// m * 4 + n = 1-d index

	// let's start the binary search 
	/*
	binary search algorithm steps:
	calculate left and right pointers
	left := 0
	right := len(matrix) * len(matrix[0])
	calculate the middle index
	mid = (left + right) / 2
	 
	while left >= right

		// convert mid back to 2d
		// assume x, y => matrix[x][y]
		// x := mid / len(matrix) // this would give a number from [0, len(matrix)]
		// e.g. mid = 7, x,y => 1,3
		// len(matrix) = 4
		// 7/4 = 1
		// 7 % 4 = 3 - 1 = 3

		// e.g. mid = 3
		// 3 / 4 = 0
		// 3 % 4 = 3

		// e.g. mid = 4
		// 4 / 4 = 1
		// 4 % 4 = 0

		x := mid / len(matrix)
		y := mid % len(matrix)

	 	if matrix[x][y] > target:
			// target is to the left of matrix[x][y]
			// set right to mid - 1
			right = mid - 1
			// set new mid
			mid = (left + right) / 2
		else if matrix[x][y] < target:
			left = mid + 1
			mid = (left + right) / 2
		else
			return true
	return false
*/
	left := 0
	right := len(matrix) * len(matrix[0]) - 1
	mid := (left + right) / 2
	fmt.Println(mid)
	for left <= right {
		x := mid / len(matrix[0])
		y := mid % len(matrix[0])
		fmt.Println(mid, x, y)
		if matrix[x][y] > target {
			// target is to the left of matrix[x][y]
			// set right to mid - 1
			right = mid - 1
			// set new mid
			mid = (left + right) / 2
		} else if matrix[x][y] < target {
			left = mid + 1
			mid = (left + right) / 2
		} else {
			return true
		}
		// fmt.Println(mid)
	}
	return false


}

func dailyTemperatures(temperatures []int) []int {
	// for each day (i) of temperatures
	// return the number of days after i before a warmer day appears
	// if no days in the future are warmer, then the result will be 0

	// given day i, how would you know many days would pass before a warmer temp is found?

	// approach 1: brute force
	// starting from i, you could iterate over the remaining elements to find the result
	// time: O(n^2)
	// space: O(1)

	// approach 2: 
	// not sure atm, but want to try to achieve this in O(n)
	/* 
	i  |  temp 	| result 
	0  |  30	| 1
	1  |  38	| 0
	2  |  30	| 1
	3  |  36	| 0



	*/ 

	// The one hint i have is that we can use a stack (since this problem is in the stack category)
	// hint #2: consider a reverse approach
	// given array [2, 1, 1, 3], consider the element with [3]. 
	// results = [3, 2, 1, 0]

	// lets look at another example:
	// [1, 2, 1, 3]
	// if we start at 3, 
	// then we can calculate the results for the rest of the array as such
	// results = [1, 2, 1, 0]
	// hint #3 says to maintain indices in a decreasing order, popping indices where values are smaller than the current element
	// wow i didnt think to store indices in  

	// approach #2
	// initialize stack
	// iterate through the array
	// at day i, compare temp[i] with the top of the stack (peek)
	// if stack is empty or less than or equal to peek()
	// push i
	// else (temp[i] is greater than peek())
	// pop from the stack until stack is empty
	// set result[i] = i - popped index

	// i = 0
	// temp[i] = 30
	// stack = []
	// since len(stack) == 0, push 0 onto stack
	// stack = [0]
	
	// i = 1
	// temp[i] = 38
	// stack = [0]
	// 38 is greater than 30
	// while len(stack) > 0
		// len(stack) == 1
		// popped = stack[0] = 0
		// results[0] = i - popped = 1

		// len(stack) = 0

	results := make([]int, len(temperatures))
	stack := []int{}
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack) - 1]] {

			results[stack[len(stack) - 1]] = i - stack[len(stack) - 1]
			// pop stack
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}
	return results
}

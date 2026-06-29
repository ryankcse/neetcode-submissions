func largestRectangleArea(heights []int) int {
	/*
		we're given an array of heights where each height corresponds to 
		the height of a rectangle with a width of 1. 
		We want to find the largest rectangle that can be formed from the 
		bars 

		examples:

			[1, 5] : These two rectangles have increasing heights. 
				the two rectangles that can be formed from these two bars are:
				1 * 5 = 5 // the bar at i=1
				or 1 * 2 = 2 // the width from both bars with the height of 1
				at i=0, stack is empty so lets push i=0 on the stack
				at i=1, the stack contains i=0, compare the height to the top of the stack (i=0). since i=1 is greater, we can push i=1 onto the stack


			[5, 2] :
				5 * 1 = 5 -- note: when i=1 was reached, i=0 could not make a rectangle with i=1 because heights[1] was less than heights[0]
				2 * 2 = 4

				at i=0, stack is empty so push i=0
				at i=1, stack has i=0, but heights[1] is less than heights[0], so i=1 is a right boundary for the rectangle at i=0
				calculate the area, determine if it's greater than the maxArea, then pop i=0, and push i=1. however, since heights[1] is less than heights[0], we could extend the rectangle made from 
				i=1 to i=0. we could push i=0 and set heights[0] to heights[1]. 

			[5, 5]:
				5 * 2 = 10 -- clearly, the largest rectangle would be the two bars put together
				at i=0, stack is empty so lets push i=0 on the stack
				at i=1, the stack contains i=0, compare the height to the top of the stack (i=0). since i=1 is equal to i=0, we can push i=1 onto the stack

			[7,1,7,2,2,4]:
				i = 0:
					1. check stack: stack is empty
					2. push 0 to stack
				i = 1:
					1. check stack: stack has 0
					2. heights[0] is greater than heights[1] // indicates right boundary
					3. when a right boundary is reached, we need to calculate all the areas for all of the indices on the stack whose heights are greater than the current index


		algorithm:
			we can use a stack to keep track of the indices of monotonically increasing heights of rectangles

			maxArea := 0
			stack := make([]int, len(heights))
			for index, height := range heights 
				if stack is empty
					then push the index to the stack
				else compare the top of stack to height
					while stack is not empty and heights[stack.pop] > height
						top = stack.pop
						stack = stack[:len(stack)-1]
						area = (index - top) * (heights[top])
						maxArea = max(maxArea, area)
					stack.push(index) // right boundary becomes the new starting point
			
			// after this loop, there may still be elements in the stack
			
			while stack is not empty
				i = len(stack)
				maxArea = max(maxArea, stack[i] * len(heights) - i)
				// right boundary = len(stack)
				// left boundary = 0
				i--
	*/
	heights = append(heights, 0)
	maxArea := 0
	stack := []int{}
	for index, height := range heights {
		for len(stack) > 0 && heights[stack[len(stack) - 1]] > height {
			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1] // pop

			width := index
			if len(stack) > 0 {
				width = index - stack[len(stack) - 1] - 1
			}
			maxArea = max(maxArea, width * heights[top]) // here, the left boundary might not be at index
			/*
			
			// Calculate width correctly based on the NEW top of the stack				
			*/
		}
		stack = append(stack, index)
	}
	return maxArea
}

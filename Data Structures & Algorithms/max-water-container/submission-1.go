func maxArea(heights []int) int {

    // We want to find the maximum amount of water
    // amount of water is determined by the area of the box given by the x and y distance
    // x distance = | i1 - i2 |
    // y distance = min(height[i2], height[i2])
    // area = x_distance * y_distance
    // Find greatest area


    // Honestly this is a type of problem that i would use a computer to solve

    // how would I approach this problem?

    // Consider the inputs: we know the input is an array of integers
    // height cannot be 0
    // We know this is a 2 pointers problem, so we have our 2 pointers i1 and i2

    // Are there any insights about the smaller index and the width of the water area?
    // We can start on the outer indices first, then start working our way towards the middle by incrementing i1/decrementing i2

    // First iteration:
    // i1 := 0
    // i2 := len(heights) - 1   // 7


    // x_distance := int(math.Abs(float64(i1 - i2)))
    // y_distance := math.Min(float64(heights[i1]), float64(heights[i2]))
    // calculate area
    // area := x_distance * y_distance
    // area = 7 = 7 * 1

    // The question is: How do we determine which pointer to update?
    // Scenario 1: we increment i1
    // Why: We should increment i1 if... area ends up being larger than if we decrement i2
    // i1_prime = 1
    // keep i2
    // x_distance = 6
    // y_distance = 6
    // area_if_i1_incremented = 36

    // Scenario 2: we decrement i2
    // i1 = 0
    // i2_prime = 7-1 = 6
    // x_distance = 6
    // y_distance = 3
    // area_if_i2_decremented = 18

    // So in this case, we'd go with scenario 1 since the area is larger than if we went with scenario 2

    // What is our base case? In other words, how do we know when to stop? Can we stop when x_distance < 1? Yes I think that can work

    // i1 := 0
    // i2 := len(heights) - 1

    // x_distance := int(math.Abs(float64(i1 - i2)))
    // y_distance := int(math.Min(float64(heights[i1]), float64(heights[i2])))

    // // calculate area
    // area := x_distance * y_distance
    // fmt.Println("x_distance:", x_distance)

    // for x_distance > 0 {
    //     fmt.Println("Hello")
    //     // calculate area if i1 is incremented
    //     i1_prime := i1 + 1
    //     // keep i2 value
        
    //     x_distance_i1_inc := int(math.Abs(float64(i1_prime - i2)))
    //     y_distance_i1_inc := int(math.Min(float64(heights[i1_prime]), float64(heights[i2])))
    //     area_i1_inc := x_distance_i1_inc * y_distance_i1_inc
        
    //     // calculate area if i2 is decremented
    //     i2_prime := i2 - 1
    //     // keep i1 value
        
    //     x_distance_i2_dec := int(math.Abs(float64(i2_prime - i1)))
    //     y_distance_i2_dec := int(math.Min(float64(heights[i1]), float64(heights[i2_prime])))

    //     area_i2_dec := x_distance_i2_dec * y_distance_i2_dec
        
    //     // check which area is greater
    //     if area_i1_inc >= area_i2_dec {
    //         // If incrementing i1 results in a greater area, increment i1
    //         i1 += 1
    //         if area_i1_inc > area {
    //             area = area_i1_inc
    //         }
    //     } else {
    //         i2 -= 1
    //         if area_i2_dec > area {
    //             area = area_i2_dec
    //         }
    //     }
    //     x_distance = int(math.Abs(float64(i1 - i2)))
    // }

    // Okay, I received some feedback and another approach is to move the pointer that's pointing to the lower bar. 
    // The reason is because increasing the lower bar would increase the area, but if you move the pointer pointing at the higher bar, the area can only decrease

    // First two pointers
    i1 := 0
    i2 := len(heights) - 1
    
    area := 0
    x := int(math.Abs(float64(i1 - i2)))
    y := int(math.Min(float64(heights[i1]), float64(heights[i2])))
    area = x * y
    
    // While i1 is before i2
    for i1 < i2 {
        

        if heights[i1] <= heights[i2] {
            i1 += 1
        } else {
            i2 -= 1
        }
        // calculate the area
        x_prime := int(math.Abs(float64(i1 - i2)))
        y_prime := int(math.Min(float64(heights[i1]), float64(heights[i2])))
        area_prime := x_prime * y_prime

        if area_prime > area {
            area = area_prime
        }
    }


// fmt.Println(area)
return area
}

func trap(height []int) int {

    // Input: height = [0,2,0,3,1,0,1,3,2,1]
    // There's actually 3 rectangles here
    // and then you subtract the squares that are "walls"
    // First approach (n^2):
    // 1. for each y value, i'm going to calculate how many squares can have water 
    // example: given x, x can only hold water if height[x-1] and height[x+1] are both greater than height[x]
    // iterate y from 0 to the maximum height
    // iterate x from 1 to len(height) - 2 

    // how do we know what the max height is?
    // at each y iteration, keep a value of the current y. if the current y is greater than all elements, then we can stop looping
    // y := 0
    // y_maxed_out := false

    // for y_maxed_out == false {

    // }

    // First approach is not going to work because this approach fails for consecutive water blocks

    // Second approach: 
    // try to aim for a single pass
    // use 2 pointers (i1 and i2)
    // 1. set both pointers to 0
    // 2. increment both pointers until a local maximum is found
    // 2a. local maximum is defined by height[i1 + 1] < height[i1]
    // 3. Set i1 to the local maximum and i2 to (i1 + 1), increment i2 until another local maximum is found
    // 4. once i2 is at the next local maximum, find the area of water between the two local maximums
    // 4a. the area of water will be the x BETWEEN (noninclusive, e.g. i1=3 and i2=7 will be 3) times the height of the smaller local maximum, minus the heights in between the two local maximums
    // 5. set i1 to i2 and repeat steps 3-5
    // note: add all areas found

//     i1, i2 := 0, 0
//     totalArea := 0
//     for i2 < len(height) - 1 {
//         // check if 0 is a local maximum
//         if height[i1 + 1] < height[i1] {
//             // // start i2 at the next index
//             // i2 += 1
//             // i1 is a local maximum, increment i2 until it is also a local maximum
//             for height[i2 + 1] > height [i2] {
//                 i2 += 1
//             }
//             // now i2 is also a local maximum

//             // find the area between i1 and i2
//             minHeight := int(math.Min(float64(height[i1]), float64(height[i2])))
//             width := int(math.Abs(float64(i1 - (i2 - 1))))   // we need i2 - 1 because it is non inclusive of i2 index itself
//             // Calculate the height * width
//             area := minHeight * width
//             // We need to subtract the squares that are acting as walls in our water area
//             for i := i1 + 1; i < i2; i++ {
//                 if height[i] >= minHeight { // clamp the squares to be subtracted to minHeight
//                     area -= minHeight
//                 } else {
//                     area -= height[i]
//                 }
//             }
//             totalArea += area

//             // set i1 to i2's local maximum
//             i1 = i2
//         } else {
//             // if not a local max, then increment both pointers
//             i1 += 1
//             i2 += 1
//         }

//     }
//     return totalArea
// }


    // 2nd approach is not working because the local maximum appraoch can detect too small of local maximums which make it not optimal to finding the entire area

    // approach 3:
    // 2 pointers
    // start both at 0
    // for each pointer, we iterate through the array and find the max value that's less than or equal to the current index's height
    // if we find a height that's greater than or equal to the current height, stop iteration
    // keep track of filler blocks inside the area
    // calculate the area and subtract the filler blocks

    totalArea := 0
    for i1 := 0; i1 < len(height) - 1; {
        area := 0
        i2 := i1 + 1
        largestEndWallPosition := i2
        // this loop is to find the wall that would trap water
        for ; i2 < len(height); i2++ {
            // wall tall enough found to trap water
            if height[i2] >= height[i1] {
                largestEndWallPosition = i2
                break
            }
            if height[i2] > height[largestEndWallPosition] {
                largestEndWallPosition = i2
            }
        }
        fmt.Println("index", i1, "highest wall pos", largestEndWallPosition)
        // highest wall that would trap water has been found. calculate area now
        width := int(math.Abs(float64(i1 - largestEndWallPosition))) - 1
        minHeight := int(math.Min(float64(height[i1]), float64(height[largestEndWallPosition])))
        // calculate fillers
        filler := 0
        for i := i1 + 1; i < largestEndWallPosition; i++ {
            if height[i] > minHeight {
                filler += minHeight
            } else {
                filler += height[i]
            }
        }
        area = (width * minHeight) - filler
        fmt.Println("index", i1, "area", area)
        totalArea += area


        // we can skip everything between i1 and largestEndWallPosition if width > 0
        if width > 0 {
            i1 = largestEndWallPosition
        } else {
            i1++
        }
    }
    return totalArea


}

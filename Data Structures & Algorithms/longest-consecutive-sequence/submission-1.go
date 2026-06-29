func longestConsecutive(nums []int) int {
    numsCount := make(map[int]int)
    for _, num := range nums {
        numsCount[num]++
    }

    retVal := 0

    for _, num := range nums {
        currLongestConsecutive := 1
        if numsCount[num - 1] < 1 && numsCount[num + 1] >= 1 { 
            // num can be considered a starting point
            for numsCount[num + currLongestConsecutive] >= 1 {
                currLongestConsecutive++
            }
        }
        if currLongestConsecutive > retVal {
            retVal = currLongestConsecutive
        }
    }

    return retVal

}

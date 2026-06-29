func maxProfit(prices []int) int {
    // Approach 1
    // at any index, keep track of the maximum value of the right side
    // loop through array backwards, and track maximums in a separate array
    // maxFuturePrice = [10,7,7,7,7,1]
    // then go back and loop through prices
    // calculate profit by maxFuturePrice[i] - prices[i]
    // return max profit

    maxProfit := 0
    // load maxFuturePrices array
    maxFuturePrice := make([]int,len(prices))
    maxFuturePrice[len(prices) - 1] = prices[len(prices) - 1] // set first value


    for i := len(prices) - 2; i >= 0; i-- {
        maxFuturePrice[i] = int(math.Max(float64(maxFuturePrice[i + 1]), float64(prices[i])))
    }

    for i := 0; i < len(prices); i++ {
        if maxFuturePrice[i] - prices[i] > maxProfit {
            maxProfit = maxFuturePrice[i] - prices[i]
        }
    }

    return maxProfit
}

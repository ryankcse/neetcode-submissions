import "slices"
func minEatingSpeed(piles []int, h int) int {
	/*
	k = eatingspeed (bananas/hour)
	piles[i] = bananas
	h = total time in hours

	we want to find smallest k that allows koko to eat all bananas in h hours

	the upper bound of k is the size of the largest pile 
	how do we find the lower bound (minimum)

	for any pile, it takes ceil(piles[i]/k) hours to finish piles[i] bananas 
	if we add all the bananas in piles, and divide by h, then we should get the avg k required to eat all bananas
	total_bananas = sum(piles)
	avg_k = total_bananas/h
	*/

	// search from k = [1, max(piles)]
	lower := 1
	upper := slices.Max(piles)

	// we want to minimize k. If we get a k such that the total time is larger than h, then we can't use this k.
	for lower < upper {

		k := lower + (upper-lower)/2
		// calculate time required
		total_time := 0
		for _, pile := range piles {
			total_time += (pile + k - 1) / k // integer ceiling arithmetic
		}
		if total_time > h {
			lower = k + 1
		} else {
			upper = k
		}
	}

	return lower
}

import (
	"slices"
)

func carFleet(target int, position []int, speed []int) int {
	// given target, position, and speed,
	// the time it takes for a car [i] to reach target is
	// y = mx + b
	// target = speed * time + position
	// time = (target - position) / speed
	// We want to calculate the times for each car to arrive at the target. 
	// we want to sort positions array in descending order


	indices := make([]int, len(position))
	for i := range indices {
		indices[i] = i
	}

	// sort the indices array
	slices.SortFunc(indices, func(a int, b int) int {
		return position[b] - position[a] // descending order
	})

	sortedPosition := make([]int, len(position))
	sortedSpeed := make([]int, len(speed))

	// rebuild sorted speed and position slices
	for new_i, old_i := range indices {
		sortedPosition[new_i] = position[old_i]
		sortedSpeed[new_i] = speed[old_i]
	}

	// calculate finish times, push finish times onto a stack
	finishTimes := make([]float64, len(position))
	for i := range sortedPosition {
		finishTimes[i] = float64(target - sortedPosition[i]) / float64(sortedSpeed[i])
	}

	fmt.Println(sortedPosition, sortedSpeed, finishTimes)


	fleets := 1
	bottleneck := finishTimes[0] // first car is the bottleneck
	// time to add up fleets
	for i := range finishTimes {
		if finishTimes[i] > bottleneck {
			bottleneck = finishTimes[i]
			fleets++
		}
	}
	return fleets
}

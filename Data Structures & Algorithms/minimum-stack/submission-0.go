type MinStack struct {
    values []int
    mins []int
}

func Constructor() MinStack {
    newMinStack := MinStack{
        values: make([]int, 0),
        mins: make([]int, 0),
    }
    return newMinStack
}

func (this *MinStack) Push(val int) {
    /*
    when we push an element, how should we keep track of min? 
    We can keep track of min as a single integer, but then problems arise when we pop that min
    // when we pop a min, how do we recalculate the min?
    // I'm thinking we keep a sorted array of values, but how do we make the insertion O(1)
    // We can't use the value as an index to an array because the value is too large


    */
    // when pushing, push the minimum between val and the top of prefix
    this.values = append(this.values, val)

    if len(this.mins) == 0 {
        this.mins = append(this.mins, val)
    } else {
        currentMin := int(math.Min(float64(val), float64(this.mins[len(this.mins) - 1])))

        this.mins = append(this.mins, currentMin)
    }

    // fmt.Println(this.values)
}

func (this *MinStack) Pop() {
    this.values = this.values[:len(this.values) - 1]
    this.mins = this.mins[:len(this.mins) - 1]
}

func (this *MinStack) Top() int {
    return this.values[len(this.values) - 1]
}

func (this *MinStack) GetMin() int {
    return this.mins[len(this.mins) - 1]
}

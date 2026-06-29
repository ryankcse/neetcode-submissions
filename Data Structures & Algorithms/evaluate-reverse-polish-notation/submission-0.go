func evalRPN(tokens []string) int {
    // if token[i] is number, push onto stack
    // if token[i] is operator, pop operand 1 and operand 2 from stack
    // perform the operation
    // push result of operation onto the stack
    // keep processing until all tokens have been processed
    stack := []int{}

    for i := 0; i < len(tokens); i++ {
        // fmt.Println(stack)
        if val, err := strconv.Atoi(tokens[i]); err == nil {
            // tokens[i] is a number
            // push onto stack

            stack = append(stack, val)
        } else {
            if tokens[i] == "+" {
                // pop top two operands
                o2 := stack[len(stack) - 1]
                o1 := stack[len(stack) - 2]
                stack = stack[:len(stack) - 1] // shrink stack by one
                // fmt.Println("Inside +:", stack)
                stack[len(stack) - 1] = o1 + o2 // replace o1 with the result of the operation
            } else if tokens[i] == "-" {
                // pop top two operands
                o2 := stack[len(stack) - 1]
                o1 := stack[len(stack) - 2]
                stack = stack[:len(stack) - 1] // shrink stack by one
                stack[len(stack) - 1] = o1 - o2 // replace o1 with the result of the operation                
            } else if tokens[i] == "*" {
                // pop top two operands
                o2 := stack[len(stack) - 1]
                o1 := stack[len(stack) - 2]
                stack = stack[:len(stack) - 1] // shrink stack by one
                stack[len(stack) - 1] = o1 * o2 // replace o1 with the result of the operation     
            } else if tokens[i] == "/" {
                // pop top two operands
                o2 := stack[len(stack) - 1]
                o1 := stack[len(stack) - 2]
                stack = stack[:len(stack) - 1] // shrink stack by one
                stack[len(stack) - 1] = o1 / o2 // replace o1 with the result of the operation     
            }
        }
    }
    // the result should be stored in stack[0]
    return stack[0]
}

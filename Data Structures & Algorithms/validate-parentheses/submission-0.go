func isValid(s string) bool {
    myStack := []byte{}
	for i := 0; i < len(s); i++ {
		// if s[i] is an open parentheses, then push it onto the stack
		// if s[i] is a closed parentheses, then pop the stack and make sure it corresponds to s[i]
		switch s[i] {
			case '(':
				myStack = append(myStack, s[i])
			case '{':
				myStack = append(myStack, s[i])
			case '[':
				myStack = append(myStack, s[i])
			case ')':
				if len(myStack) == 0 {
					return false
				}
				// pop element from stack
				popped := myStack[len(myStack) - 1]
				if popped != '(' {
					return false
				}
				myStack = myStack[:len(myStack) - 1]
			case '}':
				if len(myStack) == 0 {
					return false
				}
				// pop element from stack
				popped := myStack[len(myStack) - 1]
				if popped != '{' {
					return false
				}
				myStack = myStack[:len(myStack) - 1]
			case ']':
				if len(myStack) == 0 {
					return false
				}
				// pop element from stack
				popped := myStack[len(myStack) - 1]
				if popped != '[' {
					return false
				}
				myStack = myStack[:len(myStack) - 1]
		}
	}
	// return true if len(mystack) is 0
	return len(myStack) == 0
}

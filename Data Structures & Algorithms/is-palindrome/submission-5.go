func isPalindrome(s string) bool {
	re, _ := regexp.Compile("[^a-zA-Z0-9]+")
	cleanS := strings.ToLower(re.ReplaceAllString(s, ""))
	start := 0
	end := len(cleanS) - 1

	// loop
	// if cleanS[start]
	for start < end {
		// fmt.Println(cleanS)
		// fmt.Println(cleanS[start])//, cleanS[end])
		// fmt.Println(cleanS[end])
		if cleanS[start] == cleanS[end] {
			
			start++
			end--
		} else {
			return false
		}
	}
	return true
}

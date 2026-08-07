func lengthOfLongestSubstring(s string) int {
	max := 0

	chars := strings.Split(s, "")
	store := map[string]int{} //char to index 

	start := 0

	for i := 0; i < len(chars); i++ {
		if index, found := store[chars[i]]; found && index >= start {
			fmt.Println("reset at index: ", i, index, chars[i])
			start = index + 1
		} 

		store[chars[i]] = i

		if i - start + 1 > max {
			max = i - start + 1
		}
	}	

	return max
}
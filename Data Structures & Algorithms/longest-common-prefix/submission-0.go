func longestCommonPrefix(strs []string) string {
    result := ""

	for i := 0; i < 200; i++ {
		cur := ""

		for _, str := range strs {
			if i >= len(str) { return result }

			if cur == "" { 
				cur = string(str[i])
				continue 
			}

			if string(str[i]) != cur { return result }
		}

		result += cur 
	}

	return result
}
func longestCommonPrefix(strs []string) string {
    result := ""

	for i := 0; i < 200; i++ {
		cur := ""

		for _, str := range strs {
			switch {
				case i >= len(str): return result  
				case cur == "": cur = string(str[i])
				case string(str[i]) != cur: return result
				default: continue 
			}
		}

		result += cur 
	}

	return result
}
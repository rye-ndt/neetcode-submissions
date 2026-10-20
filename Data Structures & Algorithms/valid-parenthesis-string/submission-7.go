func checkValidString(s string) bool {	
	hi, lo := 0, 0

	for _, r := range s {
		switch r {
			case '(': lo++; hi++
			case ')': lo--; hi--
			default: lo--; hi++
		}
		
		if hi < 0 { return false }
		lo = max(lo, 0)
	}

	return lo == 0
}
func checkValidString(s string) bool {	
	left, right, star := 0, 0, 0 

	for _, r := range s {
		switch r {
			case '(': left++
			case '*': star++ 
			default:
				switch {
					case left > 0: left--
					case star > 0: star--
					default: return false
				}
		}
	}

	star = 0

	for i := len(s)-1; i >= 0; i-- {
		switch s[i] {
			case ')': right++
			case '*': star++ 
			default:
				switch {
					case right > 0: right--
					case star > 0: star--
					default: return false
				}
		}
	}

	return true 
}
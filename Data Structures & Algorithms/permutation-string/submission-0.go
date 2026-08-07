func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false 
	}

	need := [26]int{}

	for i := 0; i < len(s1); i++ {
		cur := s1[i] - 'a'
		
		need[cur]++
	}

	has := [26]int{}

	l := 0
	r := 0

	for r < len(s2) {
		curIndex := s2[r] - 'a'
		has[curIndex]++

		width := r-l+1

		if width > len(s1) {
			lIndex := s2[l] - 'a'
			has[lIndex]--
			l++
		}

		newWidth := r-l+1

		if newWidth == len(s1) && has == need {
			return true
		}

		r++
	}

	return false
}
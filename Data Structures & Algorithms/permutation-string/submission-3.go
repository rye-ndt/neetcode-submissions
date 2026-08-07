import "slices"

func sort2(v string) string {
	s := []byte(v)
	slices.Sort(s)
	return string(s)
}

func checkInclusion(s1 string, s2 string) bool {
	sortedS1 := sort2(s1)
	l, r := 0, 0

	for r < len(s2) {
		size := r - l + 1

		if size < len(s1) {
			r++ 
			continue
		}

		if sort2(s2[l:r+1]) == sortedS1 {
			return true
		}

		l++
		r++
	}

	return false
}
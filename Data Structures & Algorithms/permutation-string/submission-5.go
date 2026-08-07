import "slices"

func sortStr(v string) string {
	s := []byte(v)
	slices.Sort(s)
	return string(s)
}

func checkInclusion(s1 string, s2 string) bool {
	for r := len(s1)-1; r < len(s2); r++ {
		if sortStr(s2[r-len(s1)+1:r+1]) == sortStr(s1) {
			return true
		}
	}

	return false
}
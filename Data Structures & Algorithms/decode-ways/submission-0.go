// either take the last one, 
// or merge it with the current one, 
// if the last one is less than 2 (26 chars)
// and the current char is less than 6

func numDecodings(s string) int {
    note := make([]int, len(s)+1)

	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	note[0] = 1
	note[1] = 1

	for i := 2; i <= len(s); i++ {
		cur := s[i-1]
		last := s[i-2]

		if cur != '0' {
			note[i] += note[i-1]
		}

		if last == '1' || (last == '2' && cur <= '6') {
			note[i] += note[i-2]
		}
	}

	return note[len(s)]
}

func numDecodings(s string) int {
	if s[0] == '0' { return 0 }

	walk := make([]int, len(s)+1)
	walk[0], walk[1] = 1, 1

	for i := 2; i <= len(s); i++ {
		cur := 0

		twoDigit := num(s[i-2]) * 10 + num(s[i-1])

		if num(s[i-1]) > 0 { cur += walk[i-1] }
		if twoDigit >= 10 && twoDigit <= 26 { cur += walk[i-2] }

		walk[i] = cur
	}

	return walk[len(walk)-1]
}

func num(r byte) int {
	return int(r - '0')
}
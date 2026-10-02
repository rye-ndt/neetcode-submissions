func numDecodings(s string) int {
	if s[0] == '0' { return 0 }

	prev1, prev2 := 1, 1

	for i := 2; i <= len(s); i++ {
		cur := 0
		one := s[i-1]
		two := int(s[i-2]-'0')*10 + int(s[i-1]-'0')

		if one != '0' { cur += prev1 }
		if 10 <= two && two <= 26 { cur += prev2 }

		prev2 = prev1
		prev1 = cur
	}

	return prev1
}
func countSubstrings(s string) int {
    counter := 0

	for i := 0; i < len(s); i++ {
		counter++
		l, r := i, i 

		for r+1 < len(s) && s[r+1] == s[i] {
			counter++
			r+=1
		}

		for l > 0 && r + 1 < len(s) && s[l-1] == s[r+1] {
			counter++
			l--
			r++
		}
	}

	return counter 
}

 //aababbabca
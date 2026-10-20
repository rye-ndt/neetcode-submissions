func partitionLabels(s string) []int {
    lastPos := map[rune]int{}
	for i, r := range s { lastPos[r] = i }

	subs := []int{}

	for i := 0; i < len(s); i++ {
		lastPosCur := lastPos[rune(s[i])]

		for j := i; j < lastPosCur; j++ {
			lastPosCur = max(lastPosCur, lastPos[rune(s[j])])
		} 

		subs = append(subs, lastPosCur - i + 1)
		i = lastPosCur
	}

	return subs
}
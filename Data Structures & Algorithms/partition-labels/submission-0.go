func partitionLabels(s string) []int {
    lastIndex := map[string]int{}
	chars := strings.Split(s, "")

	for i, c := range chars {
		lastIndex[c] = i
	}

	result := []int{}
	i := 0

	for i < len(chars) {
		cur := chars[i]
		lastIndexCur := lastIndex[cur]

		for j := i; j <= lastIndexCur; j++ {
			lastIndexJ := lastIndex[chars[j]]

			if lastIndexJ > lastIndexCur {
				lastIndexCur = lastIndexJ
			}
		}

		result = append(result, lastIndexCur - i + 1)

		i = lastIndexCur+1
	}

	return result
}

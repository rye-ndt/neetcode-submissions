type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	result := ""

	for _, s := range strs {
		result += ("#" + fmt.Sprint(len(s)) + "#" + s)
	}

	return result
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}

	i := 0

	for i < len(encoded) {
		counterStr := ""

		j := i+1

		for j < len(encoded) && encoded[j] != '#' {
			counterStr += string(encoded[j])
			j++
		}

		count, _ := strconv.Atoi(counterStr)

		charStart := j+1

		result = append(result, encoded[charStart:charStart + count])

		i = charStart + count
	}

	return result
}

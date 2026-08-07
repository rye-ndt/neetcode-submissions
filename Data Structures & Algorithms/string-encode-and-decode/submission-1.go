type Solution struct{
}

const splitter = "randomstring"
const empty = "empty"

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return empty
	}

	result := ""

	for index, str := range strs {
		next := str

		if index < len(strs) - 1 {
			next += splitter
		}

		result += next
	}

	return result
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == empty {
		return []string{}
	}

	return strings.Split(encoded, splitter)
}

var nums = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	result := []string{}
	
	var backtrack func(path string, index int) 
	backtrack = func(path string, index int) {
		if len(path) == len(digits) {
			if path != "" { result = append(result, path) }
			return
		}

		for _, s := range nums[digits[index]] {
			backtrack(path + s, index+1)
		}
	}

	backtrack("", 0)

	return result
}

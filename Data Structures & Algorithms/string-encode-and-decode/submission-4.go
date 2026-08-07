// idea: encode the string length along with the strings
// so the decoder knows where to break
// 5#hello5#world

type Solution struct{
}

func (s *Solution) Encode(strs []string) string {
	lenToStr := ""

	for _, str := range strs {
		ele := strconv.Itoa(len(str)) + "#" + str
		lenToStr += ele
	}

	fmt.Println("lenToStr: ", lenToStr)

	return lenToStr
}

func (s *Solution) Decode(encoded string) []string {
	i := 0

	result := []string{}

	for i < len(encoded) {
		j := i //to trace the 2 digit numbers

		for encoded[j] != '#' {
			j++
		}

		num, err := strconv.Atoi(encoded[i:j])
		if err != nil {
			continue
		}

		fmt.Println("num: ", num)

		start := j+1
		end := start+num

		ele := encoded[start:end]

		fmt.Println("ele: ", ele)

		result = append(result, ele)
		i = end
	}

	return result
}

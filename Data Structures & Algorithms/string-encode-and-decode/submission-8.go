type Solution struct{

}

func (s *Solution) Encode(strs []string) string {
	result := ""

	for _, s := range strs {
		fmt.Println("string(len(s)): ", (len(s)))
		encoded := "#" + fmt.Sprint(len(s)) + "#" + s

		result += encoded
	}
	
	fmt.Println("result: ", result)

	return result
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}

	for i := 0; i < len(encoded); i++ {
		cur := encoded[i]

		fmt.Println("cur: ", i, string(cur))

		if cur != '#' {
			break
		}

		counterStr := ""

		j := i+1

		for j < len(encoded) && encoded[j] != '#' {
			counterStr += string(encoded[j])
			j++
		}

		fmt.Println("counter is: ", counterStr)

		count, _ := strconv.Atoi(counterStr)

		fmt.Println("extraction: ", encoded[j+1:j+1+count])

		result = append(result, encoded[j+1:j+1+count])
		i = j+count
	}

	return result
}

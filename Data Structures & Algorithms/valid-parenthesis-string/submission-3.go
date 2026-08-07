func checkValidString(s string) bool {	
	chars := strings.Split(s, "")

	stack := 0
	for i := 0; i < len(chars); i++ {
		if chars[i] == ")" {
			stack-- 

			if stack < 0 {
				return false
			}

			continue
		}

		stack++
	}

	stack = 0
	for i := len(chars)-1; i >= 0; i-- {
		if chars[i] == "(" {
			stack--

			if stack < 0 {
				return false
			}

			continue
		}

		stack++
	}

	return true 
}

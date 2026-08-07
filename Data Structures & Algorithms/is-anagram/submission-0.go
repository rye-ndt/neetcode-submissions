func isAnagram(s string, t string) bool {
    if len(t) != len(s) {
        return false
    }

    counter := map[byte]int{}

    //detect if there is a missing character first 
    for i := 0; i < len(t); i++ {
        curS := s[i]
        curT := t[i]

        counter[curS]++
        counter[curT]--
    }

    for _, counter := range counter {
        if counter != 0 {
            return false
        }
    }

    return true
}

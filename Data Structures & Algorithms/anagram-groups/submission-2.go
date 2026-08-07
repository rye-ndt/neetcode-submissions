//each str will maintain a signature, then compare those signatures

func sortStr(s string) string {
    runes := []rune(s)

    sort.Slice(runes, func(i, j int) bool {
        return runes[i] < runes[j]
    })

    return string(runes)
}

func groupAnagrams(strs []string) [][]string {
    mapper := map[string][]string{}

    for _, str := range strs {
        sig := sortStr(str)

        if _, found := mapper[sig]; !found {
            mapper[sig] = []string{}
        } 

        mapper[sig] = append(mapper[sig], str)
    }

    result := [][]string{}

    for _, val := range mapper {
        result = append(result, val)
    }

    return result
}

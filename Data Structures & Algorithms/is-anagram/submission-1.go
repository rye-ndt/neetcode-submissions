func sortStr(s string) string {
    runes := []rune(s)

    sort.Slice(runes, func(i, j int) bool {
        return runes[i] < runes[j]
    })

    return string(runes)
}

func isAnagram(s string, t string) bool {
    return sortStr(s) == sortStr(t)    
}

import "slices"

func groupAnagrams(strs []string) [][]string {
    store := map[string][]string{}

    for _, str := range strs {
        clone := str

        r := []rune(str)
        slices.Sort(r)

        store[string(r)] = append(store[string(r)], clone)
    }

    result := [][]string{}

    for _, s := range store {
        result = append(result, s)
    }

    return result
}

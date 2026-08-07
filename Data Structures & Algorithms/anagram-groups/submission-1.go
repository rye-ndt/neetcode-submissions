//each str will maintain a signature, then compare those signatures

func strSignature(str string) string {
    store := []byte(str)

    sort.Slice(store, func(i, j int) bool {
        return store[i] < store[j]
    })

    return string(store)
}


func groupAnagrams(strs []string) [][]string {
    sigMap := map[string][]string{}

    for _, str := range strs {
        sig := strSignature(str)

        sigMap[sig] = append(sigMap[sig], str)
    }

    resp := [][]string{}

    for _, val := range sigMap {
        resp = append(resp, val)
    }

    return resp
}

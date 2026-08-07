func isAnagram(s string, t string) bool {
    s1 := strings.Split(s, "")
    t1 := strings.Split(t, "")

    sort.Strings(s1)
    sort.Strings(t1)

    return strings.Join(s1, "") == strings.Join(t1, "")
}

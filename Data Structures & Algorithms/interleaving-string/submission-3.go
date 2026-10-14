func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1+s2) != len(s3) { return false }

	cache := map[[3]int]bool{}

    var dfs func(i, j, k int) bool 
	dfs = func(i, j, k int) bool {
		if k == len(s3) { return true }

		key := [3]int{i, j, k}
		if v, ok := cache[key]; ok { return v }

		if i < len(s1) && s1[i] == s3[k] {
			cache[key] = cache[key] || dfs(i+1, j, k+1)
		}

		if j < len(s2) && s2[j] == s3[k] {
			cache[key] = cache[key] || dfs(i, j+1, k+1)
		}


		return cache[key]
	}

	return dfs(0, 0, 0)
}

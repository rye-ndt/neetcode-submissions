func hasDuplicate(nums []int) bool {
    store := map[int]bool{}

    for _, num := range nums {
        if _, found := store[num]; found == true {
            return true
        }

        store[num] = true
    }    

    return false
}

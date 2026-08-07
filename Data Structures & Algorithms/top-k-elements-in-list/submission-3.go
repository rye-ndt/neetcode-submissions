type count struct {
    num int
    count int
}

func topKFrequent(nums []int, k int) []int {
    counter := map[int]count{}

    for _, n := range nums {
        increase := 1
        if cur, found := counter[n]; found {
            increase += cur.count
        }

        counter[n] = count{
            num: n, 
            count: increase,
        }
    }

    store := []count{}

    for _, c := range counter {
        store = append(store, c)
    }

    fmt.Println("store: ", store)

    sort.Slice(store, func(i, j int) bool {
        return store[i].count > store[j].count
    })

    result := []int{}

    for i := 0; len(result) < k; i++ {
        result = append(result, store[i].num)
    }

    return result
}

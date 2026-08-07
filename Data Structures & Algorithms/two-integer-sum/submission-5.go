func twoSum(nums []int, target int) []int {
   store := map[int]int{}

   for i, n := range nums {
        store[n] = i
   }

   for i, n := range nums {
        find := target - n
        index, found := store[find]
        if !found || index <= i {
            continue
        }

        return []int{i, index}
   }

   return []int{}
}

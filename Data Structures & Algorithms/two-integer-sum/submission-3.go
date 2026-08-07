// must use hash map, since you can only use 2 pointers if the array is sorted 
// else, it will fail test case nums=[-1,-2,-3,-4,-5] target=-8
func twoSum(nums []int, target int) []int {
   numToIndex := map[int]int{}

   for i, val := range nums {
        if missingIndex, found := numToIndex[target-val]; found {
            return []int{missingIndex, i}
        }

        numToIndex[val] = i
   }

   return []int{}
}

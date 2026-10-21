import "slices"

func insert(list [][]int, item []int) [][]int {
	l := 0
	for l < len(list) && list[l][1] < item[0] { // skip intervals before item
		l++
	}
	
	r := l
	for r < len(list) && list[r][0] <= item[1] { // merge all overlaps
		item = []int{min(item[0], list[r][0]), max(item[1], list[r][1])}
		r++
	}
	
	return slices.Concat(list[:l], [][]int{item}, list[r:])
}
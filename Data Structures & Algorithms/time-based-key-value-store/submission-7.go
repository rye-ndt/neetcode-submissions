type TimeMap struct {
	data map[string]map[int]string
	added map[string][]int
}

func Constructor() TimeMap {
	return TimeMap{
		data: map[string]map[int]string{},
		added: map[string][]int{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, found := this.data[key]; !found {
		this.data[key] = map[int]string{}
		this.added[key] = []int{}
	}

	this.data[key][timestamp] = value
	this.added[key] = append(this.added[key], timestamp)

	sort.Ints(this.added[key])
}

func (this *TimeMap) Get(key string, timestamp int) string {

	keyAdded, found := this.added[key]
	if !found{
		return ""
	}

	// 2
	l, r := 0, len(keyAdded)-1

	biggest := -1

	for l <= r {
		m := (l + r) / 2



		if timestamp > keyAdded[m] {
			biggest = max(biggest, keyAdded[m])
			l = m + 1
		} else if timestamp < keyAdded[m] {
			if timestamp < keyAdded[l] { 
				break
			}
			biggest = max(biggest, keyAdded[l])
			r = m - 1
		} else {
			biggest = keyAdded[m]
			break
		}
	}

	if biggest == -1 {
		return ""
	}

	return this.data[key][biggest]
}

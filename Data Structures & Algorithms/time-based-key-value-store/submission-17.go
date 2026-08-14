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
}

func (this *TimeMap) Get(key string, timestamp int) string {
	keyAdded, found := this.added[key]
	if !found { return "" }

	l, r, latest := 0, len(keyAdded)-1, -1

	for l <= r {
		m := (l + r) / 2
		if timestamp >= keyAdded[m] {
			latest = max(latest, keyAdded[m])
			l = m + 1
		} else {
			r = m - 1
		}
	}

	if latest == -1 {
		return ""
	}

	return this.data[key][latest]
}

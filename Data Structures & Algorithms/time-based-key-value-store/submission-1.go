type Data struct {
	Time int
	Val string
}

type TimeMap struct {
	Mapper map[string][]Data // name to data 
}

func Constructor() TimeMap {
	return TimeMap{
		Mapper: map[string][]Data{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.Mapper[key] = append(this.Mapper[key], Data{
		Time: timestamp,
		Val: value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	// binary search 
	vals, found := this.Mapper[key]
	if !found {
		return ""
	}

	l, r := 0, len(vals) - 1
	result := ""

	for l <= r {
		m := l + (r - l) / 2
		middleTime := vals[m].Time

		if timestamp >= middleTime {
			// middle might be the candidate 
			result = vals[m].Val
			l = m + 1
			continue 
		}

		if middleTime > timestamp {
			r = m - 1
			continue
		}
	}

	return result
}

type value1 struct {
	val []string
	ts []int
	capacity int
	i int
}
type TimeMap struct {
	time map[string]*value1
}

func Constructor() TimeMap {
	return TimeMap {
		time: make(map[string]*value1),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, exists := this.time[key]; !exists {
		this.time[key] = &value1{
			capacity: 1,
			val: make([]string, 1),
			ts: make([]int, 1),
			i: 0,
		}
	}
	this.time[key].val[this.time[key].i] = value
	this.time[key].ts[this.time[key].i] = timestamp
	this.time[key].i++
	if this.time[key].i == this.time[key].capacity {
		temp := make([]string, this.time[key].capacity)
		copy(temp, this.time[key].val[:this.time[key].capacity])
		this.time[key].val = make([]string, this.time[key].capacity*2)
		copy(this.time[key].val[:this.time[key].capacity], temp)
		temp1 := make([]int, this.time[key].capacity)
		copy(temp1, this.time[key].ts[:this.time[key].capacity])
		this.time[key].ts = make([]int, this.time[key].capacity*2)
		copy(this.time[key].ts[:this.time[key].capacity], temp1)
		this.time[key].capacity *= 2
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, exists := this.time[key]; exists {
		if this.time[key].ts[this.time[key].i -1] <= timestamp {
			return this.time[key].val[this.time[key].i -1]
		}
		r := this.time[key].i - 1
		l := 0

		for l <= r {
			m := (l + r) / 2
			if timestamp < this.time[key].ts[m] {
				r = m-1
			} else if timestamp > this.time[key].ts[m] {
				if m+1 == this.time[key].i || this.time[key].ts[m+1] > timestamp {
					return this.time[key].val[m]
				}
				l = m+1
			} else {
				return this.time[key].val[m]
			}
		}
	}
	return ""
}

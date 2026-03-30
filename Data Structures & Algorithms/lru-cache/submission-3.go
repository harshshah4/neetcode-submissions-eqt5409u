type LRUCache struct {
    capacity int
	cache []int
	maps map[int]int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		capacity: capacity,
		cache: []int{},
		maps: make(map[int]int),
	}
}

func (this *LRUCache) Get(key int) int {
	fmt.Println("getting")
	val, exists := this.maps[key]
	if !exists {
		return -1
	}
	this.removeFromSlice(key)
	this.cache = append([]int{key}, this.cache...)
	print(this)
	return val
}

func (this *LRUCache) Put(key int, value int) {
	fmt.Println("putting")
	if _, exists := this.maps[key]; exists {
		this.maps[key] = value
		this.Get(key)
		return
	}
    if len(this.cache) == this.capacity {
		valRemove := this.cache[this.capacity-1]
		this.cache = this.cache[:this.capacity-1]
		delete(this.maps, valRemove)
	}
	this.cache = append([]int{key}, this.cache...)
	this.maps[key] = value
	print(this)
}

func print(this *LRUCache){
	fmt.Printf("cache %v maps %v\n", this.cache, this.maps)
}

func (this *LRUCache) removeFromSlice(k int) {
	temp := []int{}
	for _, val := range this.cache {
		if val != k {
			temp = append(temp, val)
		}
	}
	fmt.Printf("temp %v \n", temp)
	this.cache = temp
	fmt.Println("removing")
	print(this)
	return
}
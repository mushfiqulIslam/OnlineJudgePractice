import "math/rand"

type RandomizedSet struct {
	valueMap map[int]bool
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		valueMap: map[int]bool{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	exists := this.valueMap[val]
	if exists {
		return false
	}

	this.valueMap[val] = true
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	exists := this.valueMap[val]
	if exists {
		delete(this.valueMap, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	var keys []int
	for key := range this.valueMap {
		keys = append(keys, key)
	}
	return keys[rand.Intn(len(keys))]
}


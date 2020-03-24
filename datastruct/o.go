package datastruct

type AllOne struct {
	store map[string]int
	maxAndMin MaxAndMin
}


/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		store: make(map[string]int),
		maxAndMin:MaxAndMin{},
	}
}


/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string)  {
	if v,ok:=this.store[key];ok{
		v++
		if v > this.maxAndMin.maxValue {
			this.maxAndMin.maxValue = v
			this.maxAndMin.maxKey = key
		}
		return
	}
	this.store[key] = 1
}


/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string)  {
	if v,ok:=this.store[key];ok{
		if 1 == v {
			delete(this.store,key)
			return
		}
		v--
	}
}


/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {

}


/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {

}


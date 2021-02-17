package lru_cache

type Solution interface {
	Lookup(key int) int
	Insert(key, value int)
	Erase(key int) bool
}

type LRUCache struct {
	// TODO - Add your code here
}

func NewLRUCache(capacity int) Solution {
	// TODO - Add your code here
	return &LRUCache{}
}

func (q *LRUCache) Lookup(key int) int {
	// TODO - Add your code here
	return 0
}

func (q *LRUCache) Insert(key, value int) {
	// TODO - Add your code here
}

func (q *LRUCache) Erase(key int) bool {
	// TODO - Add your code here
	return true
}

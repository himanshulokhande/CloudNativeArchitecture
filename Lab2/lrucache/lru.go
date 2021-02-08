package lrucache

import "errors"

type Cacher interface {
	Get(interface{}) (interface{}, error)
	Put(interface{}, interface{}) error
}

type lruCache struct {
	size      int
	remaining int
	cache     map[string]string
	queue     []string
}

func NewCache(size int) Cacher {
	return &lruCache{size: size, remaining: size, cache: make(map[string]string), queue: make([]string, size)}
}

func (lru *lruCache) Get(key interface{}) (interface{}, error) {
	//check if key is present in map
	if _, pres := lru.cache[key.(string)]; !pres {
		return key, errors.New("Not present in cache")
	}

	//update queue
	lru.queue = append(lru.queue, key.(string))

	//return value
	return lru.cache[key.(string)], nil
}

func (lru *lruCache) Put(key, val interface{}) error {
	//check if cache has empty space
	if lru.remaining == 0 {
		//delete least recently used entry
		delete(lru.cache, lru.queue[0])

		//update size
		lru.remaining++

		//remove entry from queue
		lru.qDel(lru.queue[0])
	}

	//put value in cache
	lru.cache[key.(string)] = val.(string)

	//update queue
	lru.queue = append(lru.queue, key.(string))

	//update cache remaining size
	lru.remaining--
	return nil
}

// Delete element from queue
func (lru *lruCache) qDel(ele string) {
	for i := 0; i < len(lru.queue); i++ {
		if lru.queue[i] == ele {
			oldlen := len(lru.queue)
			copy(lru.queue[i:], lru.queue[i+1:])
			lru.queue = lru.queue[:oldlen-1]
			break
		}
	}
}

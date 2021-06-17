package main

/*
Least recently used (LRU)
Discards the least recently used items first.
This algorithm requires keeping track of what was used when, which is expensive if one wants to make
sure the algorithm always discards the least recently used item.
*/
import (
	"container/list"
	"fmt"
)

type CacheItem struct {
	CacheKey   string
	CacheValue interface{}
}

type LRU struct {
	capacity int
	items    map[string]*list.Element
	queue    *list.List // in out case linked list will be a queue
}

func NewLru(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}

/*

When saving an item in the cache, we initialize a new Item structure and add it to the head of the c.queue.PushFront (item) queue.
The *Element returned by the queue is added to the hash table, where the key is the record identifier, and the value is a reference
to the element in queue.

*/
func (c *LRU) Set(key string, value interface{}) {
	if element, exists := c.items[key]; exists == true {
		//Before adding to the queue, we check if there is already such a key in the hash table, and if there is,
		//then we replace the value with a new one and move to the beginning of the queue c.queue.MoveToFront (element).
		c.queue.MoveToFront(element)
		element.Value.(*CacheItem).CacheValue = value
		return
	}

	// If the number of elements in the queue is equal to the maximum number of capacity,
	// then it's time to discard the last element from the queue and the key from the hash table by calling the purge () function.
	if c.queue.Len() == c.capacity {
		c.purge()
	}

	item := &CacheItem{
		CacheKey:   key,
		CacheValue: value,
	}

	element := c.queue.PushFront(item)
	c.items[item.CacheKey] = element
}

func (c *LRU) purge() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*CacheItem)
		delete(c.items, item.CacheKey)
	}
}

func (c *LRU) Get(key string) interface{} {
	element, ok := c.items[key]
	if !ok {
		return nil
	}

	c.queue.MoveToFront(element)
	return element.Value.(*CacheItem).CacheValue
}

func main() {
	cache := NewLru(2)
	cache.Set("key1", "data")
	fmt.Println(cache.Get("key1"))
	// lets add 2 more and key1 should be deleted
	cache.Set("key2", "data2")
	cache.Set("key3", "data3")
	fmt.Println(cache.Get("key1"))

}

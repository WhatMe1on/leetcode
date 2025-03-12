package main_test

import (
	"fmt"
	"testing"
)

func Test00042(t *testing.T) {
	trap := func(height []int) int {
		length := len(height)
		leftMax := make(map[int]int, len(height))
		rightMax := make(map[int]int, len(height))
		leftM := 0
		rightM := 0
		for i := range height {
			if leftM > height[i] {
			} else {
				leftM = height[i]
			}
			leftMax[i] = leftM

			if rightM > height[length-1-i] {
			} else {
				rightM = height[length-1-i]
			}
			rightMax[length-1-i] = rightM
		}

		s := 0
		for i := range height {
			if leftMax[i] > rightMax[i] {
				s += rightMax[i] - height[i]
			} else {
				s += leftMax[i] - height[i]
			}
		}
		return s
	}

	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println(trap(input))
}

// https://leetcode.cn/problems/lru-cache/description/
func Test00146_1(t *testing.T) {
	lru := Constructor(3)
	lru.Put(1, 1) // 缓存是 {1=1}
	lru.Put(2, 2) // 缓存是 {1=1, 2=2}
	lru.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lru.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Println(lru.Get(4), "   // 返回 4")
	fmt.Println(lru.Get(3), "   // 返回 3")
	fmt.Println(lru.Get(2), "   // 返回 2")
	fmt.Println(lru.Get(1), "   // 返回 -1")
	lru.Put(5, 5) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Println(lru.Get(1), "   // 返回 -1")
	fmt.Println(lru.Get(2), "   // 返回 2")
	fmt.Println(lru.Get(3), "   // 返回 3")
	fmt.Println(lru.Get(4), "   // 返回 -1")
	fmt.Println(lru.Get(5), "   // 返回 5")

	// [3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]

}

func Test00146_2(t *testing.T) {

	lru := Constructor(2)
	lru.Put(2, 1) // 缓存是 {1=1}
	lru.Put(3, 2) // 缓存是 {1=1, 2=2}
	fmt.Println(lru.Get(3), "   // 返回 2")
	fmt.Println(lru.Get(2), "   // 返回 1")
	lru.Put(4, 3) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Println(lru.Get(2), "   // 返回 1")
	fmt.Println(lru.Get(3), "   // 返回 -1")
	fmt.Println(lru.Get(4), "   // 返回 3")
	//[[2],[2,1],[3,2],[3],[2],[4,3],[2],[3],[4]]
}

type LRUCache struct {
	store   map[int]*Pair
	f       *Pair
	counter int
	cap     int
}

type Pair struct {
	prev  *Pair
	next  *Pair
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		store: map[int]*Pair{},
		cap:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.store[key]; ok {
		defer this.refresh(key)
		key := this.store[key].value
		return key
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.store[key]; ok {
		this.store[key].value = value
		this.refresh(key)
	} else if this.f == nil {
		pair := Pair{
			key:   key,
			value: value,
		}
		pair.next = &pair
		pair.prev = &pair
		this.store[key] = &pair
		this.f = &pair
		this.counter++
	} else if this.counter < this.cap {
		pair := Pair{
			prev:  this.f,
			next:  this.f.next,
			key:   key,
			value: value,
		}
		this.insertNode(&pair, this.f)
		this.f = &pair
		this.counter++
		this.store[key] = &pair
	} else {
		pair := Pair{
			prev:  this.f,
			key:   key,
			value: value,
		}
		deleteNode := this.f.next
		this.removeNode(deleteNode)
		delete(this.store, deleteNode.key)

		this.insertNode(&pair, this.f)
		this.f = &pair
		this.store[key] = &pair
	}
}

// deleteNode := this.f.next
// this.f.next.next.prev = &pair
// this.f.next = &pair
// delete(this.store, deleteNode.key)

func (this *LRUCache) refresh(key int) {
	pair := this.store[key]
	if this.f.key == pair.key {
		return
	}
	this.removeNode(pair)
	this.insertNode(pair, this.f)
	this.f = pair
}

func (this *LRUCache) removeNode(node *Pair) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) insertNode(node *Pair, prevNode *Pair) {
	prevNode.next.prev = node
	node.next = prevNode.next

	prevNode.next = node
	node.prev = prevNode
}

// LeetCode : LRU Cache
// https://leetcode.com/problems/lru-cache/
// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
//
// LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
// int get(int key) Return the value of the key if the key exists, otherwise return -1.
// void put(int key, int value) Update the value of the key if the key exists.
// Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity
// from this operation, evict the least recently used key.
//
// submission : faster than 50%

package main

import "fmt"

type Node struct {
	Key      int
	Value    int
	Previous *Node
	Next     *Node
}

type LRUCache struct {
	Capacity int
	Length   int
	Table    map[int]*Node
	Head     *Node
	Tail     *Node
}

func addHead(key int, val int, head *Node, tail *Node) (*Node, *Node) {
	node := &Node{Key: key, Value: val}
	if head == nil {
		head = node
		tail = node
	} else {
		head.Previous = node
		node.Next = head
		head = node
	}
	return head, tail
}

func cutTail(tail *Node) *Node {
	if tail == nil {
		return nil
	}
	if tail.Previous != nil {
		tail.Previous.Next = nil
	}
	return tail.Previous
}

func removeNode(node *Node, head *Node, tail *Node) (*Node, *Node) {
	if node.Previous == nil && node.Next != nil {
		// remove the head
		node.Next.Previous = nil
		head = node.Next
	} else if node.Previous != nil && node.Next == nil {
		// remove the tail
		node.Previous.Next = nil
		tail = node.Previous
	} else if node.Previous != nil && node.Next != nil {
		// remove node in middle
		pre := node.Previous
		next := node.Next
		pre.Next = next
		next.Previous = pre
	}
	node = nil
	return head, tail
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{Capacity: capacity, Length: 0, Table: map[int]*Node{}}
	return cache
}

func (this *LRUCache) Get(key int) int {
	// key exists
	if node, ok := this.Table[key]; ok {
		value := node.Value
		// remove node from linked list
		this.Head, this.Tail = removeNode(node, this.Head, this.Tail)

		// add node to head
		this.Head, this.Tail = addHead(key, value, this.Head, this.Tail)
		this.Table[key] = this.Head
		return value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// key exists
	if node, ok := this.Table[key]; ok {
		// update value
		this.Head, this.Tail = removeNode(node, this.Head, this.Tail)
		this.Head, this.Tail = addHead(key, value, this.Head, this.Tail)
		this.Table[key] = this.Head
		return
	}

	// key doesn't exist
	this.Length++
	if this.Length > this.Capacity {
		delete(this.Table, this.Tail.Key)
		this.Tail = cutTail(this.Tail)
		this.Length--
	}
	this.Head, this.Tail = addHead(key, value, this.Head, this.Tail)
	this.Table[key] = this.Head
}

func main() {
	// cache := Constructor(2)
	// cache.Put(1, 1) 	   // {1: 1}
	// cache.Put(2, 2)     // {1:1, 2:2}
	// fmt.Println(cache.Get(1)) // return 1
	// cache.Put(3, 3) 	   // {1:1, 3:3}
	// fmt.Println(cache.Get(2)) // return -1
	// cache.Put(4, 4)     // { 3:3, 4:4 }
	// fmt.Println(cache.Get(1)) // return -1
	// fmt.Println(cache.Get(3)) // return 3
	// fmt.Println(cache.Get(4)) // return 4

	cache := Constructor(3)
	cache.Put(1, 1) 	   // 1
	cache.Put(2, 2)     // 2 1
	cache.Put(3, 3) 	   // 3 2 1
	cache.Put(4, 4)     // 4 3 2
	fmt.Println("head: tail:")
	fmt.Println(cache.Head, cache.Tail)
	fmt.Println(cache.Get(4)) // return 4
	fmt.Println(cache.Get(3)) // return 3
	fmt.Println(cache.Get(2)) // return 2, 2 3 4
	fmt.Println(cache.Get(1)) // return -1
	cache.Put(5, 5)     // 5 2 3
	fmt.Println(cache.Get(1)) // return -1
	fmt.Println(cache.Get(2)) // return 2
	fmt.Println(cache.Get(3)) // return 3
	fmt.Println(cache.Get(4)) // return -1
	fmt.Println(cache.Get(5)) // return 5

}

package leetcode

type LRUCache struct {
	head, tail *Node
	data       map[int]*Node
	capacity   int
	size       int
}

type Node struct {
	key, value int
	prev, next *Node
}

func initNode(key, value int) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	// 构造一个dummy head和dummy tail
	cache := LRUCache{
		head:     initNode(0, 0),
		tail:     initNode(0, 0),
		data:     make(map[int]*Node),
		capacity: capacity,
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.data[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.data[key]
	if ok {
		node.value = value
		this.moveToHead(node)
		return
	}
	node = initNode(key, value)
	this.data[key] = node
	this.addToHead(node)
	this.size++
	if this.size > this.capacity {
		removed := this.removeTail()
		delete(this.data, removed.key)
		this.size--
	}
}

func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) removeTail() *Node {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

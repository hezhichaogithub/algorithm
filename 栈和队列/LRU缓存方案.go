// 题目: 实现LRU缓存方案。

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	cache    map[int]*Node
	size     int
	capacity int
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{0, 0, nil, nil}
	tail := &Node{0, 0, nil, nil}
	head.next = tail
	tail.prev = head

	return LRUCache{
		make(map[int]*Node),
		0,
		capacity,
		head,
		tail,
	}
}

func (this *LRUCache) addNode(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) popTail() *Node {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func (this *LRUCache) Get(key int) int {
	node := this.cache[key]
	if node == nil {
		return -1
	}

	this.moveToHead(node)

	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node := this.cache[key]
	if node == nil {
		node = &Node{
			key,
			value,
			nil,
			nil,
		}
		this.cache[key] = node
		this.addNode(node)
		this.size = this.size + 1
		if this.size > this.capacity {
			node = this.popTail()
			delete(this.cache, node.key)
			this.size = this.size - 1
		}

	} else {
		node.val = value
		this.moveToHead(node)
	}
}
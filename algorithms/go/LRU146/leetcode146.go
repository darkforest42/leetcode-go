package LRU146

type doublyLinkedNode struct {
	key, val int
	prev, next *doublyLinkedNode
}

type LRUCache struct {
	len, cap int
	head, tail *doublyLinkedNode
	cache map[int]*doublyLinkedNode
}


func Constructor(capacity int) LRUCache {
	return LRUCache{cap:capacity, cache: make(map[int]*doublyLinkedNode, capacity)}
}

func (this *LRUCache) removeLast()  {
	if this.tail.prev != nil {
		this.tail.prev.next = nil
	}else{
		this.head = nil
	}
	this.tail = this.tail.prev
}

func (this *LRUCache) insertToFirst(node *doublyLinkedNode) {
	if this.tail == nil {
		this.tail = node
	}else{
		node.next = this.head
		this.head.prev = node
	}
	this.head = node
}

func (this *LRUCache) moveToFirst(node *doublyLinkedNode){
	switch node {
	case this.head:
		return
	case this.tail:
		this.removeLast()
	default:
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	this.insertToFirst(node)
}


func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToFirst(node)
		return node.val
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	node, ok := this.cache[key]
	if ok {
		node.val = value
		this.moveToFirst(node)
	}else{
		if this.cap == this.len{
			delete(this.cache, this.tail.key)
			this.removeLast()
		}else{
			this.len++
		}
		node = &doublyLinkedNode{key,value,nil,nil}
		this.cache[key] = node
		this.insertToFirst(node)
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
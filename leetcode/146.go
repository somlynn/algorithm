package leetcode

/*
	LRU缓存机制
*/

/*
	1、通过双向链表+哈希表实现，双向链表用来维护LRU队列，淘汰最后一个元素。哈希表存储key与node的映射。
	2、put已经存在的元素,更新node的值，然后把node移动链表头部，
	3、put不存在的元素，先判断size是否等于capacity，如果等于则移除尾部元素，再把新节点插入头部。size++
	4、get元素之后，如果元素存在，把node移到链表头部.
*/

type Node struct {
	Key, Val  int
	Pre, next *Node
}

type LRUCache struct {
	Capacity   int
	Size       int
	Cache      map[int]*Node
	Head, Tail *Node
}

func Constructor2(capacity int) LRUCache {
	cache := LRUCache{
		capacity,
		0,
		make(map[int]*Node),
		&Node{},
		&Node{},
	}
	cache.Head.next = cache.Tail
	cache.Tail.Pre = cache.Head
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		this.moveToHead(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if node, ok := this.Cache[key]; ok {
		if node.Val != value {
			node.Val = value
		}
		this.moveToHead(node)
	} else {
		if this.Size == this.Capacity {
			this.deleteTail()
		}
		newNode := &Node{Key: key, Val: value}
		this.add(newNode)
		this.Cache[key] = newNode
	}
}

func (this *LRUCache) add(node *Node) {
	next := this.Head.next
	node.Pre = this.Head
	node.next = next
	this.Head.next = node
	next.Pre = node
	this.Size++
}
func (this *LRUCache) remove(node *Node) {
	pre := node.Pre
	next := node.next
	pre.next = next
	next.Pre = pre
	node.next = nil
	node.Pre = nil
	this.Size--
}

func (this *LRUCache) moveToHead(node *Node) {
	this.remove(node)
	this.add(node)
}

func (this *LRUCache) deleteTail() {
	pre := this.Tail.Pre
	this.remove(pre)
	delete(this.Cache, pre.Key)
}

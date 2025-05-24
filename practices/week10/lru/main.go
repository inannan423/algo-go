package main

import (
	"fmt"
)

// 用一个双向链表存储使用频次，链表头部是最新使用的，链表尾部是长久未使用的
// 链表节点
type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	capacity int
	hash     map[int]*Node // 用于快速查找，存链表节点！！
	head     *Node         // 用于记录频次链表头部
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		hash:     make(map[int]*Node, capacity),
		head:     nil,
		tail:     nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.hash[key]; ok {
		// 将 node 移到前面
		this.moveFront(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.hash[key]
	if ok {
		// 存在，更新
		node.Val = value
		this.moveFront(node)
	} else {
		// 不存在
		newNode := &Node{
			Val: value,
			Key: key,
		}
		this.hash[key] = newNode

		// 如果超过容量，删除最久未使用的节点（尾部节点）
		if len(this.hash) > this.capacity {
			delete(this.hash, this.tail.Key)

			// 更新尾部指针
			if this.tail.Prev != nil {
				this.tail = this.tail.Prev
				this.tail.Next = nil
			} else {
				// 只有一个节点的情况
				this.head = nil
				this.tail = nil
			}
		}

		// 将新节点添加到链表头部
		this.moveFront(newNode)
	}
}

func (this *LRUCache) moveFront(node *Node) {
	// 如果是空链表
	if this.head == nil {
		this.head = node
		this.tail = node
		node.Prev = nil
		node.Next = nil
		return
	}

	// 如果节点已经在链表中，先断开它的连接
	if node.Prev != nil || node.Next != nil || node == this.tail || node == this.head {
		// 如果是尾节点
		if node == this.tail {
			this.tail = node.Prev
			if this.tail != nil {
				this.tail.Next = nil
			}
		} else if node == this.head {
			// 如果是头节点，不需要移动
			return
		} else {
			// 中间节点
			if node.Prev != nil {
				node.Prev.Next = node.Next
			}
			if node.Next != nil {
				node.Next.Prev = node.Prev
			}
		}
	}

	// 将节点放到链表头部
	node.Next = this.head
	node.Prev = nil
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node

	// 如果链表只有这一个节点，它也是尾节点
	if this.tail == nil {
		this.tail = node
	}
}

func main() {
	cache := Constructor(2) // 创建容量为 2 的 LRU 缓存

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 输出 1

	cache.Put(3, 3)           // 淘汰 key=2
	fmt.Println(cache.Get(2)) // 输出 -1（未找到）

	cache.Put(4, 4)           // 淘汰 key=1
	fmt.Println(cache.Get(1)) // 输出 -1（未找到）
	fmt.Println(cache.Get(3)) // 输出 3
	fmt.Println(cache.Get(4)) // 输出 4
}

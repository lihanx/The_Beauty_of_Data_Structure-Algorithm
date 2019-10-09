package _06_linked_list

type DNode struct {
	Value interface{}
	Prev *DNode
	Next *DNode
}

type DoubleLinkedList struct {
	Head *DNode
	Tail *DNode
	Length int
}

// 创建新的双端节点
func NewDNode(value interface{}) *DNode {
	dnode := DNode{value, nil, nil}
	return &dnode
}

// 创建新的双端链表
func NewDoubleLinkedList() *DoubleLinkedList {
	defaultHead := NewDNode(nil)
	newDoubleLinkedList := DoubleLinkedList{defaultHead, nil, 0}
	return &newDoubleLinkedList
}

// 在指定节点前插入
func (this *DoubleLinkedList) InsertBefore(node *DNode, value interface{}) int {
	if node == nil || value == nil {
		return -1
	}

	newNode := NewDNode(value)
	prevNode := node.Prev

	newNode.Next = node
	newNode.Prev = prevNode

	prevNode.Next = newNode
	node.Prev = newNode

	// 如果给定节点是原头结点, 更新Head
	if prevNode.Value == nil {
		this.Head = newNode
	}

	this.Length++
	return this.Length
}

// 删除指定节点
func (this *DoubleLinkedList) DeleteByNode(node *DNode) int {
	if node == nil {
		return -1
	}
	prevNode := node.Prev
	nextNode := node.Next
	prevNode.Next = nextNode
	if nextNode != nil {
		nextNode.Prev = prevNode
	} else {
		if prevNode.Value != nil {
			this.Tail = prevNode			
		}
	}
	node.Prev = nil
	node.Next = nil
	this.Length--
	return this.Length
}

// 向尾部添加节点
func (this *DoubleLinkedList) Append(value interface{}) int {
	if value == nil {
		return -1
	}
	node := NewDNode(value)
	if this.Tail == nil {
		this.Head.Next = node
		node.Prev = this.Head
		this.Tail = node
	} else {
		node.Prev = this.Tail
		this.Tail.Next = node
		this.Tail = node
	}
	this.Length++
	return this.Length
}
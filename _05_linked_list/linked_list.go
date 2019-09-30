package _05_linked_list

type Node struct {
	Value interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
	Length int
}

// 创建节点
func NewNode(value interface{}) *Node {
	node := Node{value, nil}
	return &node
}

// 创建带头链表
func NewLinkedLIst() *LinkedList {
	newLinkedList := LinkedList{nil, 0}
	node := NewNode(nil)
	newLinkedList.Head = node
	return &newLinkedList
}

// 插入节点
func (this *LinkedList) Insert(idx int, value interface{}) int {
	if idx < 0 || idx >= this.Length || value == nil {
		return -1
	}
	current := this.Head
	for i := 0; i < idx -1; i++ {
		current = current.Next
	}
	node := NewNode(value)
	node.Next = current.Next
	current.Next = node
	this.Length++
	return idx
}

// 根据值删除节点
func (this *LinkedList) Delete(value interface{}) *Node {
	if value == nil || this.Length == 0 {
		return nil
	}
	before := this.Head
	for before.Next != nil {
		current := before.Next
		if current.Value == value {
			before.Next = current.Next
			current.Next = nil
			this.Length--
			return current
		}
		before = current
	}
	return nil
}

// 在指定节点之前插入节点
func (this *LinkedList) InsertBefore(posValue, newValue interface{}) int {
	if posValue == nil || newValue == nil || this.Length == 0 {
		return -1
	}
	before := this.Head
	idx := 0
	for before.Next != nil {
		current := before.Next
		if current.Value == posValue {
			node := NewNode(newValue)
			node.Next = current
			before.Next = node
			this.Length++
			return idx
		}
		before = current
		idx++
	}
	return -1
}

// 根据index查找节点
func (this *LinkedList) Find(pos int) *Node {
	if pos < 0 || pos >= this.Length {
		return nil
	}
	node := this.Head
	for i := -1; i < pos; i++ {
		node = node.Next
	}
	return node
}

// 查找指定节点的index
func (this *LinkedList) Index(value interface{}) int {
	if value == nil || this.Length == 0 {
		return -1
	}
	node := this.Head
	pos := 0
	for node.Next != nil {
		node = node.Next
		if node.Value == value {
			return pos
		}
		pos++
	}
	return -1
}

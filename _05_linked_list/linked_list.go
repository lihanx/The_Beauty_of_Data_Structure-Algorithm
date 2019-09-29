package _05_linked_list

type Node struct {
	Value interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
	Length int
}

func NewNode(value interface{}) *Node {
	node := Node{value, nil}
	return &node
}

func NewLinkedLIst() *LinkedList {
	newLinkedList := LinkedList{nil, 0}
	node := NewNode(nil)
	newLinkedList.Head = node
	return &newLinkedList
}

func (this *LinkedList) Insert(pos int, value interface{}) {}
func (this *LinkedList) Delete(value interface{}) {}
func (this *LinkedList) InsertBefore(posValue, newValue interface{}) {}

func (this *LinkedList) DeleteBefore(posValue interface{}) {}

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

func (this *LinkedList) Index(value interface{}) int {
	if value == nil {
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

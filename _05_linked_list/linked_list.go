package _05_linked_list

type Node struct {
	Value interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
	Length int
}

func NewNode() *Node {
	node := Node{nil, nil}
	return &node
}

func NewLinkedLIst() *LinkedList {
	newLinkedList := LinkedList{nil, 0}
	node := NewNode()
	newLinkedList.Head = node
	return &newLinkedList
}

func (this *LinkedList) Insert(pos int, value interface{}) {}
func (this *LinkedList) Delete(value interface{}) {}
func (this *LinkedList) InsertBefore(posValue, newValue interface{}) {}
func (this *LinkedList) DeleteBefore(posValue interface{}) {}
func (this *LinkedList) Find(pos int) {}
func (this *LinkedList) Index(value interface{}) {}

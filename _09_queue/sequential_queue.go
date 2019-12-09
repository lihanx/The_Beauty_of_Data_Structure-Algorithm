package _09_queue

// 利用取模运算的周期性
// 实现基于数组的循环队列

type ArrayQueue struct {
	Items     []interface{}
	MaxLength int
	Head      int
	Tail      int
}

func NewArrayQueue(maxLength int) *ArrayQueue {
	items := make([]interface{}, maxLength, maxLength)
	aq := ArrayQueue{
		Items:     items,
		MaxLength: 0,
		Head:      0,
		Tail:      0,
	}
	return &aq
}

// 出队
func (this *ArrayQueue) dequeue() interface{} {
	if this.Head == this.Tail {
		return nil
	}
	item := this.Items[this.Head]
	this.Head = (this.Head + 1) % this.MaxLength
	return item
}

// 入队
// 为避免频繁数据搬移，根据队列长度固定的特点
// 使用取模的周期运算可以很好的实现无需数据搬移的空间循环利用
func (this *ArrayQueue) enqueue(value interface{}) bool {
	if (this.Tail+1)%this.MaxLength == this.Head {
		return false
	}
	this.Items[this.Tail] = value
	this.Tail = (this.Tail + 1) % this.MaxLength
	return true
}

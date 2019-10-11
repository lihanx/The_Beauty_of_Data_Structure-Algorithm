package _08_stack

// 顺序栈
type SequentialStack struct {
	Items []interface{}
	Length int
	MaxLength int
}

// 创建顺序栈
func NewSequentialStack(maxLength int) *SequentialStack {
	Items := make([]interface{}, maxLength, maxLength)
	seqStack := SequentialStack{
		Items:     Items,
		Length:    maxLength,
		MaxLength: maxLength,
	}
	return &seqStack
}

// 出栈
func (this *SequentialStack) pop() interface{} {
	if this.Length == 0 {
		return nil
	}
	value := this.Items[this.Length-1]
	this.Length--
	return value
}

// 压栈
func (this *SequentialStack) push(value interface{}) int {
	if this.Length == this.MaxLength {
		return -1
	}
	this.Items[this.Length] = value
	this.Length++
	return this.Length
}
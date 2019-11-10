// 题目: 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
// push(x) -- 将元素 x 推入栈中。
// pop() -- 删除栈顶的元素。
// top() -- 获取栈顶元素。
// getMin() -- 检索栈中的最小元素。

type MinStack struct {
	items []int
	mins  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.items = append(this.items, x)
	if (len(this.mins) == 0) || (x <= this.mins[len(this.mins)-1]) {
		this.mins = append(this.mins, x)
	}
}

func (this *MinStack) Pop() {
	length := len(this.items)
	if length == 0 {
		return
	}

	var item int
	item = this.items[length-1]
	if length == 1 {
		this.items = make([]int, 0)
	} else {
		this.items = this.items[:length-1]
	}

	if len(this.mins) > 0 && item == this.mins[len(this.mins)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	length := len(this.items)
	if length == 0 {
		return 0
	}

	return this.items[length-1]
}

func (this *MinStack) GetMin() int {
	if len(this.mins) == 0 {
		return 0
	}

	return this.mins[len(this.mins)-1]
}
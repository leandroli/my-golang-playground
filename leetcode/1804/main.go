package main

import "container/list"

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.Len() == 0 {
		if this.stack1.Len() == 0 {
			return -1
		}
		for this.stack1.Len() > 1 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
		return this.stack1.Remove(this.stack1.Back()).(int)
	} else {
		return this.stack2.Remove(this.stack2.Back()).(int)
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

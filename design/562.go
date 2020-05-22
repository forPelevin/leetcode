package design

import "math"

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/98/design/562/
//Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//
//push(x) -- Push element x onto stack.
//pop() -- Removes the element on top of the stack.
//top() -- Get the top element.
//getMin() -- Retrieve the minimum element in the stack.
//
//
//Example 1:
//
//Input
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//Output
//[null,null,null,null,-3,null,0,-2]
//
//Explanation
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin(); // return -3
//minStack.pop();
//minStack.top();    // return 0
//minStack.getMin(); // return -2
//
//
//Constraints:
//
//Methods pop, top and getMin operations will always be called on non-empty stacks.
type MinStack struct {
	items []item
}

type item struct {
	val int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		items: make([]item, 0, 100),
	}
}

func (ms *MinStack) Push(x int) {
	min := x
	if len(ms.items) > 0 && ms.topItem().min < x {
		min = ms.topItem().min
	}
	ms.items = append(ms.items, item{
		val: x,
		min: min,
	})
}

func (ms *MinStack) Pop() {
	ms.items = ms.items[:len(ms.items)-1]
}

func (ms *MinStack) Top() int {
	return ms.topItem().val
}

func (ms *MinStack) GetMin() int {
	return ms.topItem().min
}

func (ms *MinStack) topItem() item {
	return ms.items[len(ms.items)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

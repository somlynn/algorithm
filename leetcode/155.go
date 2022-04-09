package leetcode

// 赋值栈
// 当一个元素要入栈时，我们取当前辅助栈的栈顶存储的最小值，与当前元素比较得出最小值，将这个最小值插入辅助栈中；
// 当一个元素要出栈时，我们把辅助栈的栈顶元素也一并弹出；
// 在任意一个时刻，栈内元素的最小值就存储在辅助栈的栈顶元素中。
//

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor3() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: []int{},
	}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	if len(s.minStack) == 0 {
		s.minStack = append(s.minStack, val)
		return
	}
	top := s.minStack[len(s.minStack)-1]
	s.minStack = append(s.minStack, Min(val, top))
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}

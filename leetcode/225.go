package leetcode

// 使用一个队列实现栈：思路就是把新加入之前的元素都移到后面去，保证头部是最新的
// 3
// 3 4 -> 4 3
// 4 3 5 -> 5 4 3

type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor8() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {
	s.queue2 = append(s.queue2, x)
	for len(s.queue1) > 0 {
		s.queue2 = append(s.queue2, s.queue1[0])
		s.queue1 = s.queue1[1:]
	}
	s.queue1, s.queue2 = s.queue2, s.queue1
}

func (s *MyStack) Pop() int {
	if len(s.queue1) > 0 {
		e := s.queue1[0]
		s.queue1 = s.queue1[1:]
		return e
	}
	return -1
}

func (s *MyStack) Top() int {
	if len(s.queue1) > 0 {
		return s.queue1[0]
	}
	return -1
}

func (s *MyStack) Empty() bool {
	return len(s.queue1) == 0
}

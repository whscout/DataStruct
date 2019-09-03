package stack

type (
	Stack struct {
		top *node
		length int
	}

	node struct {
		value interface{}
		prev *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

// get Stack len
func (s *Stack) Len() int {
	return s.length
}

// get Stack peek item
func (s *Stack) Peek() interface{} {
	return s.top.value
}

// pop (out stack)
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	node := s.top
	s.top = node.prev
	s.length--
	return node.value
}

// push (in stack)
func (s *Stack) Push(value interface{}) {
	node := &node{
		value: value,
		prev:  s.top,
	}
	s.top = node
	s.length--
}

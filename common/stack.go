package common

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*node
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}
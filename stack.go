package goutline

//Stack - stack datastructure for nodes
type Stack struct {
	Stack []*Node
}

//Len - return length of the stack
func (s *Stack) Len() int {
	return len(s.Stack)
}

//Push - put an item on the end of the stack
func (s *Stack) Push(n *Node) {
	s.Stack = append(s.Stack, n)
	return
}

//Pop - removes and returns last item on stack
func (s *Stack) Pop() *Node {
	n, o := s.Stack[s.Len()-1], s.Stack[:s.Len()-1]
	s.Stack = o
	return n
}

//Reverse - returns a new stack with the current one reversed
func (s *Stack) Reverse() *Stack {
	newStack := *s
	for i := len(newStack.Stack)/2 - 1; i >= 0; i-- {
		opp := len(newStack.Stack) - 1 - i
		newStack.Stack[i], newStack.Stack[opp] = newStack.Stack[opp], newStack.Stack[i]
	}
	return &newStack
}

//IndexOfLevel - searches a stack backwards for the first appearance of a level
// and returns its index
func (s *Stack) IndexOfLevel(level int) int {
	stack := s.Reverse()
	index := -1
	for i, node := range stack.Stack {
		if node.Level == level {
			index = stack.Len() - i
		}
	}
	return index
}

//PopToLevel - pops all the items off up to the level passed.
func (s *Stack) PopToLevel(level int) *Stack {
	newStack := Stack{}

	index := s.IndexOfLevel(level)
	newStack.Stack, s.Stack = s.Stack[index:], s.Stack[:index]
	return &newStack
}

func (s *Stack) BuildTree() *Node {
	var root *Node
	var prevNode *Node
	for _, node := range s.Stack {
		if root == nil {
			root = node
		} else {
			prevNode.Children = append(prevNode.Children, node)
		}
		prevNode = node
	}
	return root
}

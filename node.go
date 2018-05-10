package goutline

//Node -
type Node struct {
	Content  string
	Parent   *Node
	Children []*Node
	NodeType string
	Level    int
}

//AddChild - add a child node to the node.
func (n *Node) AddChild(child *Node) error {
	n.Children = append(n.Children, child)
	child.Parent = n
	return nil
}

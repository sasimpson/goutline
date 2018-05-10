package goutline

import "errors"

type Tree struct {
	Root         *Node
	CurrentNode  *Node
	PreviousNode *Node
	Size         int
	Depth        int
}

//Initialize - get our node tree started right.
func (t *Tree) Initialize() error {
	t.Root = &Node{}
	t.CurrentNode = t.Root
	return nil
}

//AddNode - adds a node to a tree
func (t *Tree) AddNode(node *Node) error {
	if t.Root == nil {
		return errors.New("no valid tree, please initialize")
	}

	//node is a child of current node
	if node.Level > t.CurrentNode.Level {
		err := t.CurrentNode.AddChild(node)
		if err != nil {
			return err
		}
		t.PreviousNode = t.CurrentNode
		t.CurrentNode = node
		t.Size++
		if node.Level > t.Depth {
			t.Depth = node.Level
		}
		return nil
	}

	//node is a peer of current node
	if node.Level == t.CurrentNode.Level {
		if t.CurrentNode.Parent == nil {
			return errors.New("no parent for node")
		}
		t.CurrentNode.Parent.AddChild(node)
		t.PreviousNode = t.CurrentNode
		t.CurrentNode = node
		t.Size++
		if node.Level > t.Depth {
			t.Depth = node.Level
		}
		return nil
	}
	//node is a peer of the parent of the current node
	if node.Level < t.CurrentNode.Level {
		if t.CurrentNode.Parent == nil {
			return errors.New("no parent for node")
		}
		c := t.CurrentNode
		p := c.Parent
		for p != nil {
			if node.Level == p.Level {
				p.Parent.AddChild(node)
				t.PreviousNode = t.CurrentNode
				t.CurrentNode = node
				t.Size++
				if node.Level > t.Depth {
					t.Depth = node.Level
				}
				return nil
			}
			p = p.Parent
		}
		return errors.New("No peer found in tree")
	}
	return errors.New("node not added, conditions failed")
}

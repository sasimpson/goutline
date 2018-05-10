package goutline

import (
	"errors"
	"regexp"
)

//ParseLine - parses a line of data into a valid node
func ParseLine(line string) (*Node, error) {
	acceptableNodeRE := regexp.MustCompile("^[.*]+ .*$")
	getNodeDataRE := regexp.MustCompile("^([*|.]+) (.*)")
	isStarRE := regexp.MustCompile("[*]+")
	isDotRE := regexp.MustCompile("[.]+")

	if ok := acceptableNodeRE.MatchString(line); !ok {
		return nil, errors.New("no matches ")
	}
	if ok := getNodeDataRE.MatchString(line); ok {
		node := &Node{}
		matches := getNodeDataRE.FindStringSubmatch(line)
		if isStarRE.MatchString(matches[1]) {
			node.NodeType = "star"
		}
		if isDotRE.MatchString(matches[1]) {
			node.NodeType = "dot"
		}
		node.Content = matches[2]
		node.Level = len(matches[1])
		return node, nil
	}
	return nil, errors.New("invalid node line")
}

// //OrganizeNodes - take the list of nodes, organize those
// // into a tree based on order and level
// func OrganizeNodes(nodes []*Node) (*Node, error) {
// 	//each node will have a level.  each node will be put on a
// 	// stack until the level is less than or equal to the previous level
// 	//if the previous level is > current then we must flush the stack
// 	//to build out the tree up to that point.
// 	stack := Stack{}
// 	//root node to hold our tree
// 	root := &Node{Level: 0, Content: "root"}
// 	//prev node to track the last node processed.
// 	prevNode := root

// 	for i, node := range nodes {
// 		switch {
// 		case node.Level > prevNode.Level:
// 			//current level is below the last one, so we can add to the stack
// 			stack.Push(node)
// 			prevNode = node
// 		case node.Level <= prevNode.Level:
// 			//current level is above or equal to the previous, flush stack
// 			//get index of matching level off stack
// 			poppedStack := stack.PopToLevel(node.Level)
// 			nodeTree := poppedStack.BuildTree()

// 			prevNode = node
// 		}
// 		fmt.Printf("%d: %#v\n", i, node)

// 	}
// 	return root, nil
// }

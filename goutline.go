package goutline

import (
	"errors"
	"fmt"
	"regexp"
)

type Node struct {
	Content  string
	Children []*Node
	NodeType string
	Level    int
}

//parses a line of data into a valid node
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

func OrganizeNodes(nodes []*Node) (*Node, error) {
	root := &Node{Level: 0, Content: "root"}

	for i, node := range nodes {
		switch {
		case i == 0:
			root.Children = append(root.Children, node)
			// case i>0:
			// 	root.Children = append()
		}
		fmt.Printf("%d: %#v\n", i, node)
	}
	return root, nil
}

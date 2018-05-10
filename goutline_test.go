package goutline

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	testCases := []struct {
		desc        string
		input       string
		output      *Node
		returnError error
	}{
		{
			desc:   "star valid",
			input:  "* this is a valid node",
			output: &Node{Content: "this is a valid node", NodeType: "star", Level: 1},
		},
		{
			desc:   "dot valid",
			input:  ". this is a valid node",
			output: &Node{Content: "this is a valid node", NodeType: "dot", Level: 1},
		},
		{
			desc:   "star depth valid",
			input:  "** this is a valid node",
			output: &Node{Content: "this is a valid node", NodeType: "star", Level: 2},
		},
		{
			desc:        "invalid node",
			input:       "this is an invalid node",
			output:      nil,
			returnError: errors.New("no matches"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			node, err := ParseLine(tC.input)
			if tC.returnError == nil {
				assert.Nil(t, err)
				assert.EqualValues(t, tC.output, node)
			} else {
				assert.Error(t, tC.returnError, err)
			}

		})
	}
}

// func TestOrganizeNodes(t *testing.T) {
// 	testCases := []struct {
// 		desc        string
// 		input       []*Node
// 		returnError error
// 		rootSize    int
// 	}{
// 		{
// 			desc: "flat tree",
// 			input: []*Node{
// 				&Node{Content: "first node", NodeType: "star", Level: 1},
// 				&Node{Content: "another first node", NodeType: "star", Level: 1},
// 				&Node{Content: "yet another first node", NodeType: "star", Level: 1},
// 			},
// 			rootSize: 3,
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			root, err := OrganizeNodes(tC.input)
// 			fmt.Printf("%#v\n", root)
// 			if tC.returnError == nil {
// 				assert.Nil(t, err)
// 				assert.Equal(t, tC.rootSize, len(root.Children))
// 			} else {
// 				assert.Error(t, tC.returnError, err)
// 			}
// 		})
// 	}
// }

func TestNodeStack(t *testing.T) {
	testCases := []struct {
		desc  string
		nodes []*Node
	}{
		{
			desc: "one node",
			nodes: []*Node{
				&Node{Level: 1, Content: "A"},
			},
		},
		{
			desc: "two nodes",
			nodes: []*Node{
				&Node{Level: 1, Content: "A"},
				&Node{Level: 1, Content: "B"},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			testStack := Stack{}
			//test push results in right length of stack
			for _, node := range tC.nodes {
				testStack.Push(node)
			}
			assert.Equal(t, len(tC.nodes), testStack.Len())
			//test stack order
			for i, node := range tC.nodes {
				assert.EqualValues(t, node, testStack.Stack[i])
			}
			//test pop order
			counter := 0
			for i := testStack.Len(); i > 0; i-- {
				node := testStack.Pop()
				assert.EqualValues(t, tC.nodes[i-1], node)
				counter++
			}
			assert.Equal(t, len(tC.nodes), counter)
		})
	}
}

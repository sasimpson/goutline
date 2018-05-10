package goutline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNode(t *testing.T) {
	testCases := []struct {
		desc        string
		nodes       []*Node
		size        int
		depth       int
		returnError error
	}{
		{
			desc: "flat",
			nodes: []*Node{
				&Node{Level: 1, Content: "A"},
				&Node{Level: 1, Content: "B"},
			},
			size:  2,
			depth: 1,
		},
		{
			desc: "5 nodes, 3 levels",
			nodes: []*Node{
				&Node{Level: 1, Content: "A"},
				&Node{Level: 2, Content: "B"},
				&Node{Level: 3, Content: "C"},
				&Node{Level: 1, Content: "D"},
				&Node{Level: 2, Content: "E"},
			},
			size:  5,
			depth: 3,
		},
		{
			desc: "5 nodes, 2 levels",
			nodes: []*Node{
				&Node{Level: 1, Content: "A"},
				&Node{Level: 2, Content: "B"},
				&Node{Level: 2, Content: "C"},
				&Node{Level: 1, Content: "D"},
				&Node{Level: 2, Content: "E"},
			},
			size:  5,
			depth: 2,
		},
		{
			desc: "init error",
			nodes: []*Node{
				&Node{Level: 1, Content: "A"},
				&Node{Level: 2, Content: "B"},
				&Node{Level: 2, Content: "C"},
				&Node{Level: 1, Content: "D"},
				&Node{Level: 2, Content: "E"},
			},
			size:  5,
			depth: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var tree Tree
			tree.Initialize()
			for _, node := range tC.nodes {
				err := tree.AddNode(node)
				assert.Nil(t, err)
			}
			assert.Equal(t, tC.size, tree.Size)
			assert.Equal(t, tC.depth, tree.Depth)
		})
	}

	var tree Tree
	err := tree.AddNode(&Node{Level: 1})
	assert.EqualErrorf(t, err, "no valid tree, please initialize", err.Error())
}

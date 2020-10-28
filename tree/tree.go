package tree

import (
	"github.com/alonana/isolationforest/node"
	"github.com/alonana/isolationforest/point"
	"github.com/alonana/isolationforest/points"
)

type Tree struct {
	root node.Node
}

func Build(data *points.Points, maxDepth uint) *Tree {
	t := Tree{}
	t.root = node.Build(data, maxDepth)
	return &t
}

func (t *Tree) ToString() string {
	return t.root.ToString("")
}

func (t *Tree) Score(p *point.Point) float32 {
	return t.root.Score(p, 0)
}

package node

import (
	"fmt"
	"github.com/alonana/isolationforest/point"
	"github.com/alonana/isolationforest/points"
	"github.com/alonana/isolationforest/score"
)

type Node interface {
	ToString(prefix string) string
	Score(p *point.Point, pathLength uint) float32
}

type externalNode struct {
	size uint
}

type internalNode struct {
	splitAttribute uint
	splitValue     float32
	left           Node
	right          Node
}

func Build(data *points.Points, remainingDepth uint) Node {
	if remainingDepth == 0 || data.Size() <= 1 || data.AllIdentical() {
		return &externalNode{
			size: data.Size(),
		}
	}

	splitAttribute := data.RandomAttribute()
	internal := internalNode{
		splitAttribute: splitAttribute,
		splitValue:     data.RandomValue(splitAttribute),
	}

	smaller, bigger := data.Split(internal.splitAttribute, internal.splitValue)
	internal.left = Build(smaller, remainingDepth-1)
	internal.right = Build(bigger, remainingDepth-1)

	return &internal
}

func (n *externalNode) ToString(prefix string) string {
	return fmt.Sprintf("%vsize=%v", prefix, n.size)
}

func (n *internalNode) ToString(prefix string) string {
	self := fmt.Sprintf("%v%v split %v", prefix, n.splitAttribute, n.splitValue)
	left := n.left.ToString(prefix + "+ ")
	right := n.right.ToString(prefix + "\\ ")
	return self + "\n" + left + "\n" + right
}

func (n *externalNode) Score(_ *point.Point, pathLength uint) float32 {
	return float32(pathLength) + score.C(n.size)
}

func (n *internalNode) Score(p *point.Point, pathLength uint) float32 {
	if p.Value(n.splitAttribute) < n.splitValue {
		return n.left.Score(p, pathLength+1)
	}
	return n.right.Score(p, pathLength+1)
}

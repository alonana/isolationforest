/*
See:
https://cs.nju.edu.cn/zhouzh/zhouzh.files/publication/icdm08b.pdf?q=isolation-forest

https://www.codementor.io/@arpitbhayani/isolation-forest-algorithm-for-anomaly-detection-133euqilki

https://en.wikipedia.org/wiki/Isolation_forest

*/

package isolationforest

import (
	"fmt"
	"github.com/alonana/isolationforest/point"
	"github.com/alonana/isolationforest/points"
	"github.com/alonana/isolationforest/score"
	"github.com/alonana/isolationforest/tree"
	"math"
	"strings"
)

type Forest struct {
	trees     []tree.Tree
	cDataSize float32
}

func Create(treesAmount uint, baseline *points.Points) *Forest {
	f := Forest{
		trees: make([]tree.Tree, treesAmount),
	}
	f.cDataSize = score.C(baseline.Size())
	f.learn(baseline)
	return &f
}

func (f *Forest) learn(baseline *points.Points) {
	// we are only interested in data points that have shorter-than-average path lengths
	maxDepth := uint(math.Ceil(math.Log2(float64(baseline.Size()))))

	for i := range f.trees {
		f.trees[i] = *tree.Build(baseline, maxDepth)
	}
}

func (f *Forest) Score(p *point.Point) float32 {
	sumHeights := float32(0)
	for _, t := range f.trees {
		sumHeights += t.Score(p)
	}
	averageHeight := sumHeights / float32(len(f.trees))

	result := math.Pow(2, float64(-averageHeight/f.cDataSize))
	return float32(result)
}

func (f *Forest) ToString() string {
	result := make([]string, 0)
	for i, t := range f.trees {
		result = append(result, fmt.Sprintf("tree %v:", i))
		result = append(result, t.ToString())
	}
	return strings.Join(result, "\n")
}

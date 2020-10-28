package isolationforest

import (
	"fmt"
	"github.com/alonana/isolationforest/point"
	"github.com/alonana/isolationforest/points"
	"math/rand"
	"testing"
)

func Test(t *testing.T) {
	baseline := points.Create()

	for value := 0; value < 1000; value++ {
		x := float32(rand.NormFloat64())
		y := float32(rand.NormFloat64())
		baseline.Add(point.Create(x, y))
	}
	fmt.Printf("%v\n", baseline)

	f := Create(100, baseline)
	//fmt.Printf("%v\n", f.ToString())

	for radix := 0; radix < 10; radix++ {
		fmt.Printf("radix %v score: %v\n", radix, f.Score(point.Create(float32(radix), float32(radix))))
	}
}

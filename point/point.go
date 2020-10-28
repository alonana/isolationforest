package point

import (
	"fmt"
	"math/rand"
)

type Point struct {
	axis []float32
}

func Create(axis ...float32) *Point {
	p := Point{
		axis: make([]float32, len(axis)),
	}
	for i, x := range axis {
		p.axis[i] = x
	}
	return &p
}

func (p *Point) Attributes() uint {
	return uint(len(p.axis))
}

func (p *Point) RandomAttribute() uint {
	attributes := len(p.axis)
	return uint(rand.Intn(attributes))
}

func (p *Point) Value(attribute uint) float32 {
	return p.axis[attribute]
}

func (p *Point) Identical(other Point) bool {
	for i, value := range p.axis {
		if other.axis[i] != value {
			return false
		}
	}
	return true
}
func (p *Point) ToString() string {
	return fmt.Sprintf("%v", p.axis)
}

package points

import (
	"github.com/alonana/isolationforest/point"
	"math/rand"
)

type Points struct {
	points []point.Point
}

func Create() *Points {
	p := Points{
		points: make([]point.Point, 0),
	}
	return &p
}

func (p *Points) Add(point *point.Point) {
	p.points = append(p.points, *point)
}

func (p *Points) Size() uint {
	return uint(len(p.points))
}

func (p *Points) RandomAttribute() uint {
	relevantAttributes := make([]bool, p.points[0].Attributes())
	for i := range relevantAttributes {
		relevantAttributes[i] = p.hasDifferentValues(uint(i))
	}

	randomSize := 0
	for _, relevant := range relevantAttributes {
		if relevant {
			randomSize++
		}
	}

	randomValue := rand.Intn(randomSize)
	for i, relevant := range relevantAttributes {
		if relevant {
			if randomValue == 0 {
				return uint(i)
			}
			randomValue--
		}
	}

	panic("should not get here")
}

func (p *Points) hasDifferentValues(attribute uint) bool {
	var firstValue float32
	for i, dataPoint := range p.points {
		if i == 0 {
			firstValue = dataPoint.Value(attribute)
		} else {
			if firstValue != dataPoint.Value(attribute) {
				return true
			}
		}
	}
	return false
}

func (p *Points) RandomValue(attribute uint) float32 {
	var minUpdated, maxUpdated bool
	var min, max float32
	for _, dataPoint := range p.points {
		value := dataPoint.Value(attribute)
		if !maxUpdated || value > max {
			max = value
			maxUpdated = true
		}
		if !minUpdated || value < min {
			min = value
			minUpdated = true
		}
	}
	result := rand.Float32() * (max - min)
	return min + result
}

func (p *Points) Split(attribute uint, value float32) (*Points, *Points) {
	smaller := Create()
	bigger := Create()
	for _, p := range p.points {
		if p.Value(attribute) < value {
			smaller.Add(&p)
		} else {
			bigger.Add(&p)
		}
	}
	return smaller, bigger
}

func (p *Points) AllIdentical() bool {
	if len(p.points) <= 1 {
		return true
	}

	for i := 0; i < len(p.points)-1; i++ {
		for j := i; j < len(p.points); j++ {
			if !p.points[i].Identical(p.points[j]) {
				return false
			}
		}
	}
	return true
}

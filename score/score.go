package score

import "math"

const eulerConst = float64(0.5772156649)

func C(n uint) float32 {
	if n == 1 {
		return 0
	}
	result := 2 * harmonic(n-1)
	result -= float32((2 * (n - 1)) / n)
	return result
}

func harmonic(i uint) float32 {
	return float32(math.Log(float64(i)) + eulerConst)
}

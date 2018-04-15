package libs

import (
	"math"
)

type Vector []float64
type Matrix []Vector
type Distance float64

func VectorNorm(v Vector) float64 {
	var norm float64
	for _, elem := range v {
		norm += elem * elem
	}
	return math.Sqrt(norm)
}

func CosineDistance(v1, v2 Vector) Distance {
	v1Norm := VectorNorm(v1)
	v2Norm := VectorNorm(v2)
	var sum float64
	for i := 0; i < len(v1); i++ {
		sum += v1[i] * v2[i]
	}
	return Distance(1 - sum/(v1Norm*v2Norm))
}

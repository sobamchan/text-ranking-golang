package libs

import (
	"testing"
)

func TestVectorNorm(t *testing.T) {
	v := Vector{1.0, 1.5}
	r := VectorNorm(v)
	println(r)
}

func TestCosineDistance(t *testing.T) {
	v1 := Vector{1.0, 1.5}
	v2 := Vector{1.0, 1.5}
	r := CosineDistance(v1, v2)
	println(r)
}

package space

import (
	"github.com/kuentra-official/hnsw/mathk"
	"github.com/kuentra-official/hnsw/simd/avx"
)

type avxSpaceImpl struct{}

func (avxSpaceImpl) EuclideanDistance(a, b mathk.Vector) float32 {
	return avx.EuclideanDistance(a, b)
}

func (avxSpaceImpl) ManhattanDistance(a, b mathk.Vector) float32 {
	return avx.ManhattanDistance(a, b)
}

func (avxSpaceImpl) CosineDistance(a, b mathk.Vector) float32 {
	return avx.CosineDistance(a, b)
}

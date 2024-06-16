package space

import (
	"github.com/kuentra-official/hnsw/mathk"
	"github.com/kuentra-official/hnsw/simd/sse"
)

type sseSpaceImpl struct{}

func (sseSpaceImpl) EuclideanDistance(a, b mathk.Vector) float32 {
	return sse.EuclideanDistance(a, b)
}

func (sseSpaceImpl) ManhattanDistance(a, b mathk.Vector) float32 {
	return sse.ManhattanDistance(a, b)
}

func (sseSpaceImpl) CosineDistance(a, b mathk.Vector) float32 {
	return sse.CosineDistance(a, b)
}

package space

import "github.com/kuentra-official/hnsw/mathk"

type nativeSpaceImpl struct{}

func (nativeSpaceImpl) EuclideanDistance(a, b mathk.Vector) float32 {
	var distance float32
	for i := 0; i < len(a); i++ {
		distance += mathk.Square(a[i] - b[i])
	}

	return mathk.Sqrt(distance)
}

func (nativeSpaceImpl) ManhattanDistance(a, b mathk.Vector) float32 {
	var distance float32
	for i := 0; i < len(a); i++ {
		distance += mathk.Abs(a[i] - b[i])
	}

	return distance
}

func (nativeSpaceImpl) CosineDistance(a, b mathk.Vector) float32 {
	var dot float32
	var aNorm float32
	var bNorm float32
	for i := 0; i < len(a); i++ {
		dot += a[i] * b[i]
		aNorm += mathk.Square(a[i])
		bNorm += mathk.Square(b[i])
	}

	return 1.0 - dot/(mathk.Sqrt(aNorm)*mathk.Sqrt(bNorm))
}

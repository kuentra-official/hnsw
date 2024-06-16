package space

import (
	"github.com/klauspost/cpuid"
	"github.com/kuentra-official/hnsw/mathk"
)

type SpaceImpl interface {
	EuclideanDistance(mathk.Vector, mathk.Vector) float32
	ManhattanDistance(mathk.Vector, mathk.Vector) float32
	CosineDistance(mathk.Vector, mathk.Vector) float32
}

type Space interface {
	Distance(mathk.Vector, mathk.Vector) float32
}

type space struct {
	impl SpaceImpl
}

func newSpace() space {
	if cpuid.CPU.AVX() {
		return space{impl: avxSpaceImpl{}}
	}
	if cpuid.CPU.SSE() {
		return space{impl: sseSpaceImpl{}}
	}

	return space{impl: nativeSpaceImpl{}}
}

type Euclidean struct{ space }

type Manhattan struct{ space }

type Cosine struct{ space }

func NewEuclidean() Space {
	return &Euclidean{newSpace()}
}

func (this *Euclidean) Distance(a, b mathk.Vector) float32 {
	return this.impl.EuclideanDistance(a, b)
}

func (this *Euclidean) String() string {
	return "euclidean"
}

func NewManhattan() Space {
	return &Manhattan{newSpace()}
}

func (this *Manhattan) Distance(a, b mathk.Vector) float32 {
	return this.impl.ManhattanDistance(a, b)
}

func (this *Manhattan) String() string {
	return "manhattan"
}

func NewCosine() Space {
	return &Cosine{newSpace()}
}

func (this *Cosine) Distance(a, b mathk.Vector) float32 {
	return mathk.Abs(this.impl.CosineDistance(a, b))
}

func (this *Cosine) String() string {
	return "cosine"
}

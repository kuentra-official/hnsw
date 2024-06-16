package main

import (
	"context"
	"log"
	"sort"

	"github.com/kuentra-official/hnsw"
	"github.com/kuentra-official/hnsw/mathk"
	"github.com/kuentra-official/hnsw/space"
	uuid "github.com/satori/go.uuid"
)

func main() {
	euclidean := space.NewEuclidean()
	index := hnsw.NewHnsw(128, euclidean,
		hnsw.HnswSearchAlgorithm(hnsw.HnswSearchHeuristic))

	ids := make([]uuid.UUID, 100)
	vectors := make([][]float32, 100)
	for idx := range vectors {
		ids[idx] = uuid.NewV4()
		vectors[idx] = mathk.RandomNormalVector(128, mathk.RandomUniform()*10, 1000)
		index.Insert(ids[idx], vectors[idx], nil, index.RandomLevel())
	}

	q := mathk.RandomUniformVector(128)

	result, err := index.Search(context.Background(), q, 100)
	if err != nil {
		log.Fatal(err)
	}
	bruteForceResult := make(hnsw.SearchResult, 100)
	for i, vec := range vectors {
		bruteForceResult[i] = hnsw.SearchResultItem{
			Id:    ids[i],
			Score: euclidean.Distance(vec, q),
		}
	}
	sort.Sort(bruteForceResult)

	for i := 0; i < 10; i++ {
		log.Println(uuid.Equal(result[i].Id, bruteForceResult[i].Id))
	}
}

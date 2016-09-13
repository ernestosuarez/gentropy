package entropy

import (
	"fmt"
	"math"
)

// Computes the entropy of a slice
func Entropy(data []string) float64 {

	m := make(map[string]int)

	length := float64(len(data))

	for _, line := range data {
		m[line] += 1
	}

	var entro float64 = 0.0
	for _, v := range m {
		counts := float64(v)
		entro += -counts / length * math.Log(counts/length)
	}

	return entro
}

func EntropyMatrix(data [][]int) float64 {
	m := make(map[string]int)

	length := float64(len(data))

	for _, line := range data {
		m[fmt.Sprint(line)] += 1
	}

	var entro float64 = 0.0
	for _, v := range m {
		counts := float64(v)
		entro += -counts / length * math.Log(counts/length)
	}

	return entro
}

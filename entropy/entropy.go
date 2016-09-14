package entropy

import (
	"fmt"
	"math"
)

type Sample [][]string

// Len() computes the length of the sequence
func (s Sample) Size() int { return len(s) }

// Nvar() computes the number of variables (columns)
// in the sequence
func (s Sample) Nvar() int { return len(s[0]) }

//Compute the frequency counts for each label
func GetFrequencyCounts(s Sample) map[string]int {

	countsMap := make(map[string]int)

	for _, value := range s {
		countsMap[fmt.Sprint(value)] += 1
	}

	return countsMap
}

// Computes the entropy of a slice
func EntropyML(s Sample) float64 {
	var entro float64
	countsMap := GetFrequencyCounts(s)
	totCounts := s.Size()

	for _, count := range countsMap {
		prob := float64(count) / float64(totCounts)
		if prob != 0.0 {
			entro += -prob * math.Log(prob)
		}
	}

	return entro
}

func EntropyMM(s Sample) float64 {
	totCounts := s.Size()
	lenMap := len(GetFrequencyCounts(s))
	return EntropyML(s) + (float64(lenMap)-1.0)/(2*float64(totCounts))
}

func EntropyChaoShen(s Sample) float64 {
	var entro float64
	totCounts := s.Size()
	countsMap := GetFrequencyCounts(s)

	var numSingletons float64

	for _, value := range countsMap {
		if value == 1 {
			numSingletons++
		}
	}

	for _, count := range countsMap {
		prob := (1.0 - numSingletons/float64(totCounts)) * (float64(count) / float64(totCounts))
		if prob != 0.0 {
			entro += (-prob * math.Log(prob)) / (1.0 - math.Pow((1.0-prob), float64(totCounts)))
		}
	}
	return entro

}


// Computes the mutual information expansion up to order maxOrder
//func EntropyMIE(s Sample, estimator string, maxOrder int) []float64 {
//	Size := s.Size()
//	Nvar := s.Nvar()
//
//	suma := make([]float64, maxOrder)
//	X := make([]int, maxOrder)
//
//	M := make(Sample, Size)
//	for i := range M {
//		M[i] = make([]int, Nvar)
//	}
//
//
//	for order:=1; order <= maxOrder; order++ {
//		for i:=1; i<= order; i++ {
//			X[i] = i
//		}
//
//		for i:=1; i <= Size; i++ {
//			for j:=1; j<= order; j++ {
//				M[i][j] = s[i][X[j]]
//			}
//
//		}
//		suma(order) += EntropyML(M)
//
//
//
//	}
//	return suma
//}
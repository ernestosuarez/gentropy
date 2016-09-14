package entropy

import (
	"fmt"
	"math"
)

//The SequenceI Interface
type SequenceI interface {
	Get(i int) string
	Len() int
}

// Sequence with N rows and 1 column (1 variable sequence)
type Sequence1D []string

// Sequence with N rows and M columns (M variables sequence)
type SequenceND [][]string

func (s Sequence1D) Len() int { return len(s) }
func (s SequenceND) Len() int { return len(s) }

func (s Sequence1D) Get(i int) string { return fmt.Sprint(s[i]) }
func (s SequenceND) Get(i int) string { return fmt.Sprint(s[i]) }

//Compute the frequency counts for each label
func GetFrequencyCounts(seq SequenceI) map[string]int {

	countsMap := make(map[string]int)

	for i := 1; i < seq.Len(); i++ {
		countsMap[fmt.Sprint(seq.Get(i))] += 1
	}

	return countsMap
}

// Computes the entropy of a slice
func EntropyML(seq SequenceI) float64 {
	var entro float64
	countsMap := GetFrequencyCounts(seq)
	totCounts := seq.Len()

	for _, count := range countsMap {
		prob := float64(count) / float64(totCounts)
		if prob != 0.0 {
			entro += -prob * math.Log(prob)
		}
	}

	return entro
}

func EntropyMM(seq SequenceI) float64 {
	totCounts := seq.Len()
	lenMap := len(GetFrequencyCounts(seq))
	return EntropyML(seq) + (float64(lenMap)-1.0)/(2*float64(totCounts))
}

func EntropyChaoShen(seq SequenceI) float64 {
	var entro float64
	totCounts := seq.Len()
	countsMap := GetFrequencyCounts(seq)

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

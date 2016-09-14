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
type SequenceNx1 []string

// Sequence with N rows and M columns (M variables sequence)
type SequenceNxM [][]string

func (s SequenceNx1) Len() int { return len(s) }
func (s SequenceNxM) Len() int { return len(s) }

func (s SequenceNx1) Get(i int) string { return fmt.Sprint(s[i]) }
func (s SequenceNxM) Get(i int) string { return fmt.Sprint(s[i]) }

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

	totCounts := float64(seq.Len())
	for _, c := range countsMap {
		count := float64(c)
		entro += -count / totCounts * math.Log(count/totCounts)
	}

	return entro
}

func EntropyMM(seq SequenceI) float64 {
	totCounts := float64(seq.Len())
	counts := float64(len(GetFrequencyCounts(seq)))
	return EntropyML(seq) + (counts-1.0)/(2*totCounts)

}

//func EntropyChaoShen


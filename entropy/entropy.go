package entropy

import (
	"fmt"
	"math"
	"sync"
)

var (
	//ML is alias for Max. Likelihood method
	ML = MaxLikelihood

	//MM is alias for Miller-Madow method
	MM = MillerMadow

	//CS is alias for the Chao-Shen method
	CS = ChaoShen

	//MIE is awsome
	MIE = MutualInformationExpansion

	mutex sync.Mutex
)

//Sample is contains a sample of multidimensional variables
//Each row would be a single data point.
type Sample [][]string

// Size computes and returns the length of the sample.
func (s Sample) Size() int { return len(s) }

// Nvar computes the number of variables (columns)
// in the sequence.
func (s Sample) Nvar() int { return len(s[0]) }

//GetFrequencyCounts compute the absolute frequency of each
// observed data point.
func GetFrequencyCounts(s Sample) map[string]int {

	countsMap := make(map[string]int)

	for _, value := range s {
		countsMap[fmt.Sprint(value)]++
	}

	return countsMap
}

// MaxLikelihood Computes the Max. Likelihood entropy
// from a sample.
func MaxLikelihood(s Sample) float64 {
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

// MillerMadow Computes the Miller-Madow entropy from a sample.
func MillerMadow(s Sample) float64 {

	totCounts := float64(s.Size())
	lenMap := float64(len(GetFrequencyCounts(s)))

	return MaxLikelihood(s) + (lenMap-1.0)/(2*totCounts)
}

// ChaoShen Computes the Chao-Shen entropy from a sample.
func ChaoShen(s Sample) float64 {
	var entro float64
	totCounts := float64(s.Size())
	countsMap := GetFrequencyCounts(s)

	var numSingletons float64

	for _, value := range countsMap {
		if value == 1 {
			numSingletons++
		}
	}

	for _, count := range countsMap {
		prob := (1.0 - numSingletons/totCounts) * (float64(count) / totCounts)
		if prob != 0.0 {
			entro += (-prob * math.Log(prob)) / (1.0 - math.Pow((1.0-prob), totCounts))
		}
	}
	return entro

}

// MutualInformationExpansion computes the mutual information expansion
// up to order maxOrder.
// See Suarez et al. J. Chem. Theory Comput., 2011, 7 (8), pp 2638–2653
func MutualInformationExpansion(s Sample, maxOrder int) float64 {
	var (
		entro []float64
		wg    sync.WaitGroup
		order int
	)

	Size := s.Size()
	Nvar := s.Nvar()
	suma := make([]float64, maxOrder)


	for order = 1; order <= maxOrder; order++ {

		for comb := range combinations(Nvar, order) {
			combCopy := make([]int, order)
			copy(combCopy, comb)
			wg.Add(1)
			go func(combCopy []int) {
				M := newEmptySample(Size, Nvar)
				for i := 0; i < Size; i++ {
					for j := 0; j < order; j++ {
						M[i][j] = s[i][combCopy[j]]
					}
				}
				entropySubSet := MaxLikelihood(M)
				mutex.Lock()
				suma[order-1] += entropySubSet
				mutex.Unlock()
				wg.Done()
			}(combCopy)


		}
		wg.Wait()

		entro = make([]float64, maxOrder)
		for i := 1; i <= order; i++ {
			for j := 1; j <= i; j++ {
				entro[i-1] += mieCoefficient(Nvar, i, j) * suma[j-1]
			}
		}

		switch order {
		case 1:
			fmt.Println("MIE entropy: ")
			fmt.Printf("Sum of marginal entropies in nats: %.5f\n", entro[order-1])
		case 2:
			fmt.Printf("2nd order correction: %.5f\n", entro[order-1]-entro[order-2])
		case 3:
			fmt.Printf("3rd order correction: %.5f\n", entro[order-1]-entro[order-2])
		default:
			fmt.Printf("%dth order correction: %.5f\n", order, entro[order-1]-entro[order-2])
		}

	}

	return entro[maxOrder-1]
}

// mieCoefficient computes the coefficient needed to compute
// the entropy from its mutual information expansion.
func mieCoefficient(Nvar, maxOrder, k int) float64 {
	var coeff float64
	var binomial float64

	for i := 0; i <= maxOrder-k; i++ {
		binomial = 1.0
		if i > 0 {
			for j := 1; j <= i; j++ {
				binomial *= float64((Nvar - k - j + 1)) / float64(j)
			}
		}
		coeff += math.Pow(-1.0, float64(i)) * binomial
	}
	return coeff

}

// NewEmptySample
func newEmptySample(nRows, nCols int) Sample {
	M := make(Sample, nRows)
	for i := range M {
		M[i] = make([]string, nCols)
	}
	return M
}

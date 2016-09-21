package entropy_test

import (
	"flag"
	"fmt"
	"github.com/ernestosuarez/gentropy/entropy"
	"runtime"
	"testing"
)

func TestEntro(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	pathPtr := flag.String("path", "../data/sequence.txt", "Path to the sequence file")

	flag.Parse()
	path := *pathPtr

	sequenceND := entropy.ReadSequenceND(path)

	fmt.Println("\nML entropy: ", entropy.ML(sequenceND), "\n")

	fmt.Println("MM entropy: ", entropy.MM(sequenceND), "\n")

	fmt.Println("ChaoShen entropy: ", entropy.CS(sequenceND), "\n")

	order := 6

	fmt.Println("MIE entropy of order", order, ": ", entropy.MIE(sequenceND, order), "\n")

	//fmt.Println("Coeff: ", entropy.MieCoefficient(10, 5, 2), "\n")

	//fmt.Println("MIE entorpy n var: ", entropy.EntropyMIE(sequenceND, " ", 3), "\n")
}

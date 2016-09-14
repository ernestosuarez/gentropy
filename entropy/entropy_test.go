package entropy_test

import (
	"flag"
	"fmt"
	"github.com/ernestosuarez/gentropy/entropy"
	"testing"
)

func TestEntro(t *testing.T) {

	pathPtr := flag.String("path", "../data/sequence.txt", "Path to the sequence file")

	flag.Parse()
	path := *pathPtr

	sequenceND := entropy.ReadSequenceND(path)

	fmt.Println("\nML entropy: ", entropy.EntropyML(sequenceND), "\n")

	fmt.Println("MM entropy: ", entropy.EntropyMM(sequenceND), "\n")

	fmt.Println("ChaoShen entropy: ", entropy.EntropyChaoShen(sequenceND), "\n")

	//fmt.Println("MIE entorpy n var: ", entropy.EntropyMIE(sequenceND, " ", 3), "\n")
}

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

	sequence, _ := entropy.ReadSequence1D(path)
	sequenceND := entropy.ReadSequenceND(path)

	fmt.Println("ML entorpy 1 var: ", entropy.EntropyML(sequence))
	fmt.Println("ML entorpy n var: ", entropy.EntropyML(sequenceND), "\n")

	fmt.Println("MM entorpy 1 var: ", entropy.EntropyMM(sequence))
	fmt.Println("MM entorpy n var: ", entropy.EntropyMM(sequenceND), "\n")

	fmt.Println("ChaoShen entorpy 1 var: ", entropy.EntropyChaoShen(sequence))
	fmt.Println("ChaoShen entorpy n var: ", entropy.EntropyChaoShen(sequenceND), "\n")
}

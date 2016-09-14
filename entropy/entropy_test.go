package entropy_test

import (
	"flag"
	"fmt"
	. "github.com/ernestosuarez/gentropy/entropy"
	"log"
	"testing"
)

func TestEntro(t *testing.T) {

	pathPtr := flag.String("path", "../data/sequence.txt", "Path to the sequence file")
	//pathPtr := flag.String("path", "./data/sequence.txt", "Path to the sequence file")

	_ = flag.String("estimator", "mle", "Entropy estimator")

	flag.Parse()
	path := *pathPtr

	sequence, err := ReadSequenceNx1(path)
	if err != nil {
		log.Fatal(err)
	}

	entro := EntropyML(sequence)

	sequenceND := ReadSequenceNxM(path)

	fmt.Println(entro)
	fmt.Println(EntropyML(sequenceND))
	fmt.Println(EntropyMM(sequenceND))
}

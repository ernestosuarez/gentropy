package main

import (
	"flag"
	"fmt"
	ge "github.com/ernestosuarez/gentropy/entropy"
	"log"
)

func main() {

	pathPtr := flag.String("path", "./data/sequence.txt", "Path to the sequence file")

	_ = flag.String("estimator", "mle", "Entropy estimator")

	flag.Parse()
	path := *pathPtr

	sequence, err := ge.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	entropy := ge.Entropy(sequence)

	dataMatrix := ge.ReadMatrixInt(path)

	fmt.Println(entropy)
	fmt.Println(ge.EntropyMatrix(dataMatrix))
}

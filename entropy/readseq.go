package entropy

import (
	"bufio"
	"os"
	"strings"
)

//Error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadSequence1D Reads a whole file into memory
// with N rows and 1 column (Nx1) and returns the
// issequence as []string (alias SequenceNx1)
func ReadSequence1D(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sequence []string //sequence
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sequence = append(sequence, scanner.Text())
	}

	return sequence, scanner.Err()
}

// ReadSequenceND Reads a whole file into memory
// with N rows and M column (NxM) and returns the whole
// matrix (NxM) as [][]string (alias SequenceNXM)
func ReadSequenceND(path string) Sample {
	lines, err := ReadSequence1D(path)
	check(err)

	nRows := len(lines) // number of rows

	//Computing the number of columns
	//from the first, not empty row, later we will check
	//that all the rows have the same size.
	nCols := 0
	for _, line := range lines {
		nFields := len(strings.Fields(line))
		if nFields != 0 {
			nCols = nFields
			break
		}
	}

	if nCols == 0 {
		panic("No data found")
	}

	sequence := make([][]string, nRows)

	for i, line := range lines {
		lineSlice := strings.Fields(line)
		if len(lineSlice) != nCols {
			panic("The number or elemnts should be the same for each row")
		}

		row := make([]string, nCols)

		for j, value := range lineSlice {
			row[j] = value
		}

		sequence[i] = row
	}

	return sequence
}

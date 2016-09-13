package entropy

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadMatrixInt(path string) [][]int {
	lines, err := ReadLines(path)
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

	dataMatrix := make([][]int, nRows)

	for i, line := range lines {
		lineSlice := strings.Fields(line)
		if len(lineSlice) != nCols {
			panic("The number or elemnts should be the same for each row")
		}

		intRow := make([]int, nCols)

		for j, value := range lineSlice {
			intValue, err := strconv.Atoi(value)
			check(err)
			intRow[j] = intValue
		}

		dataMatrix[i] = intRow
	}

	return dataMatrix
}

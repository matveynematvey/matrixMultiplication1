package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	fileName       = "matrix"
	fileNameSize   = "matrixWithSize"
	resultFileName = "resultMatrix"
)

func fileReaderOS() (matrix [][]int) {
	file, err := os.Open(fileNameSize)

	check(err, "ReaderOS fail")
	defer func() { check(file.Close(), "File closing error") }()

	buf := make([]byte, 1)

	file.Read(buf)

	matrixSize, err := strconv.Atoi(string(buf[0]))
	check(err, "size error")

	matrix = make([][]int, matrixSize)
	for ind, _ := range matrix {
		matrix[ind] = make([]int, matrixSize)
	}

	file.Read(buf)
	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			file.Read(buf)
			matrix[i][j], err = strconv.Atoi(string(buf[0]))
			check(err, "Error reading matrix")
			file.Read(buf)
		}
	}

	return
}

func fileReaderBufio() string {
	file, err := os.Open(fileName)

	check(err, "ReaderOS fail")
	defer func() { check(file.Close(), "File closing error") }()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var buf string

	for scanner.Scan() {
		buf += scanner.Text() + " "
	}

	return buf
}

func fileReaderIO() string {
	data, err := ioutil.ReadFile(fileName)

	check(err, "ReaderIO fail")

	return string(data)
}

func parseMatrix(buf string) [][]int {
	sliceString := strings.Fields(buf)
	matrixSize := int(math.Sqrt(float64(len(sliceString))))

	matrix := make([][]int, matrixSize)

	for ind, val := range sliceString {
		valInt, err := strconv.Atoi(val)

		check(err, "Error parsing matrix")

		matrix[ind/matrixSize] = append(matrix[ind/matrixSize], valInt)
	}

	return matrix
}

func multiplyMatrix(matrix *[][]int) {
	factor := initFactor()

	for _, line := range *matrix {
		for ind, _ := range line {
			line[ind] *= factor
		}
	}
}

func initFactor() (factor int) {
	fmt.Print("Enter factor: ")
	fmt.Fscan(os.Stdin, &factor)
	return factor
}

func writeMatrixToFileOS(matrix *[][]int) {
	file, err := os.Create(resultFileName)
	check(err, "Error creating file")

	for _, line := range *matrix {
		file.WriteString(fmt.Sprintln(line))
	}
}

func check(err error, text string) {
	if err != nil {
		fmt.Println(text)
	}
}

func main() {
	matrix := fileReaderOS()
	fmt.Println(matrix)
	multiplyMatrix(&matrix)
	writeMatrixToFileOS(&matrix)
}

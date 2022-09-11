package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	//"os"
)

const (
	fileName = "matrix"
	factor   = 5
)

// func fileReader(name string) string {
// 	f, err := os.Open("D:\\matrixMultiplication\\matrixParams.txt")

// 	if err != nil {
// 		fmt.Println("reader fail", err)
// 	}
// 	defer f.Close()

// 	return string(data)
// }

func fileReaderIO(name string) string {
	data, err := ioutil.ReadFile(name)

	if err != nil {
		fmt.Println("reader fail", err)
	}

	return string(data)
}

func parseMatrix() [][]int {
	buf := fileReaderIO(fileName)
	sliceString := strings.Fields(buf)
	matrixSize := int(math.Sqrt(float64(len(sliceString))))

	matrix := make([][]int, matrixSize)

	for ind, val := range sliceString {
		valInt, err := strconv.Atoi(val)

		if err != nil {
			panic(err)
		}

		matrix[ind/matrixSize] = append(matrix[ind/matrixSize], valInt)
	}

	return matrix
}

func multiplyMatrix(matrix *[][]int) {
	for _, line := range *matrix {
		for ind, _ := range line {
			line[ind] *= factor
		}
	}
}

func main() {
	matrix := parseMatrix()
	fmt.Println(matrix)
	multiplyMatrix(&matrix)
	fmt.Println(matrix)
}

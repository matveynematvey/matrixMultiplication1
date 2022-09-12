package main

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	fileName       = "matrix"
	resultFileName = "resultMatrix"
)

// func fileReaderOS() string {
// 	file, err := os.Open(fileName)

// 	check(err, "ReaderOS fail")
// 	defer func() { check(file.Close(), "File closing error") }()

// 	//buf := make([]byte, 32) // define your buffer size here.
// 	// for {
// 	// 	n, err := file.Read(buf)

// 	// 	if n > 0 {
// 	// 		fmt.Print(buf[:n]) // your read buffer.
// 	// 	}

// 	// 	if err == io.EOF {
// 	// 		break
// 	// 	}
// 	// 	if err != nil {
// 	// 		check(err, "Reading error")
// 	// 		break
// 	// 	}
// 	// }
// 	//scanner := bufio.NewScanner(file)

// 	// for scanner.Scan() {
// 	// 	line := scanner.Text()

// 	// }
// 	// return line
// 	//return scanner.Text()
// }

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
	//factor := initFactor()
	factor := 5

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
	matrix := parseMatrix(fileReaderIO())
	fmt.Println(matrix)
	multiplyMatrix(&matrix)
	writeMatrixToFileOS(&matrix)

}

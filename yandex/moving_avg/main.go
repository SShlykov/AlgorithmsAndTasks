package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func movingAverage(array []int, windowSize int) []float64 {
	if windowSize == 0 {
		return nil
	}
	var res []float64
	var currSum int
	for i := 0; i < windowSize; i++ {
		currSum += array[i]
	}
	res = append(res, float64(currSum)/float64(windowSize))

	for i := 0; i <= len(array)-windowSize-1; i++ {
		currSum -= array[i]
		currSum += array[i+windowSize]
		currAvg := float64(currSum) / float64(windowSize)
		res = append(res, currAvg)
	}

	return res
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	array := readArray(scanner)
	windowSize := readInt(scanner)
	printArray(movingAverage(array, windowSize))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(arr []float64) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.FormatFloat(arr[i], 'f', 8, 64))
		writer.WriteString(" ")
	}
	writer.Flush()
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	SPACE     = 1
	TAB       = 4
	Backspace = -1
)

func main() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	lines, err := getTaskInput(reader)
	if err != nil {
		writeResult(writer, err.Error())
		os.Exit(2)
	}

	var goal int
	for _, countNeededSpaces := range lines {
		goal += getCountOperations(countNeededSpaces)
	}

	writeResult(writer, strconv.Itoa(goal))
}

func getCountOperations(countNeededSpaces int) int {
	if countNeededSpaces == 0 {
		return 0
	}
	divisor := countNeededSpaces / TAB

	viaBackspace := (divisor + 1) + (countNeededSpaces-(divisor+1)*TAB)/Backspace
	viaSpace := (divisor) + (countNeededSpaces-divisor*TAB)/SPACE

	return min(viaBackspace, viaSpace)
}

func getTaskInput(reader *bufio.Reader) ([]int, error) {
	linesNumber, err := readLine(reader)
	if err != nil {
		return []int{}, err
	}

	result := make([]int, linesNumber)
	for i := 0; i < linesNumber; i++ {
		var number int
		number, err = readLine(reader)
		if err != nil {
			return []int{}, err
		}
		result[i] = number
	}

	return result, nil
}

func readLine(reader *bufio.Reader) (int, error) {
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n\r")

	number, err := strconv.Atoi(line)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func writeResult(writer *bufio.Writer, result string) {
	_, _ = writer.WriteString(result)
}

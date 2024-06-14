package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	TAB = 4
)

func main() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	lines, _ := getTaskInput(reader)

	var goal int
	for _, countNeededSpaces := range lines {
		goal += getCountOperations(countNeededSpaces)
	}

	writeResult(writer, strconv.Itoa(goal))
}

func getCountOperations(countNeededSpaces int) int {
	divisor := countNeededSpaces / TAB
	reminder := countNeededSpaces % TAB

	switch reminder {
	case 1, 2:
		return divisor + reminder
	case 3:
		return divisor + 2
	default:
		return divisor
	}
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

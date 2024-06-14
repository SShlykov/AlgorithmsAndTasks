package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	PLUS     = "+"
	MULTIPLY = "x"
)

const (
	EVEN = iota
	ODD
)

func main() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	nums, err := getNumbers(reader)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	res := turnIntoOperations(nums)

	writeResult(writer, res)
}

func turnIntoOperations(numbers []int) string {
	if len(numbers) < 2 {
		return ""
	}
	result := strings.Builder{}
	var prev int
	prevOperation := MULTIPLY // Multiply to 1
	firstOddFound := false

	if numbers[0]%2 == 0 {
		prev = EVEN
	} else {
		prev = ODD
		firstOddFound = true
	}

	for i := 1; i < len(numbers); i++ {
		var curr int
		if numbers[i]%2 == 0 {
			curr = EVEN
		} else {
			curr = ODD
		}

		switch {
		case prev == ODD && prevOperation == MULTIPLY:
			if curr == ODD {
				result.WriteString(MULTIPLY)
				firstOddFound = true
			} else {
				prev = EVEN
				prevOperation = PLUS
				result.WriteString(PLUS)
			}
		case prev == ODD && prevOperation == PLUS:
			if curr == ODD {
				result.WriteString(MULTIPLY)
				prevOperation = MULTIPLY
				firstOddFound = true
			} else {
				prev = EVEN
				result.WriteString(PLUS)
			}
		case prev == EVEN && prevOperation == PLUS:
			if curr == ODD {
				prev = ODD
				prevOperation = MULTIPLY
				result.WriteString(MULTIPLY)

				firstOddFound = true
			} else {
				prevOperation = MULTIPLY
				result.WriteString(MULTIPLY)
			}
		case prev == EVEN && prevOperation == MULTIPLY:
			if curr == ODD {
				prev = ODD
				if firstOddFound {
					prevOperation = MULTIPLY
					result.WriteString(MULTIPLY)
				} else {
					prevOperation = PLUS
					result.WriteString(PLUS)
				}
				firstOddFound = true
			} else {
				result.WriteString(MULTIPLY)
			}
		}
	}

	return result.String()
}

func getNumbers(reader *bufio.Reader) ([]int, error) {
	first, err := readLine(reader)
	if err != nil {
		return nil, err
	}
	var countNumbers int
	countNumbers, err = strconv.Atoi(first)
	if err != nil {
		return nil, err
	}

	second, err := readLine(reader)
	if err != nil {
		return nil, err
	}
	res := make([]int, countNumbers)
	nums := strings.Split(second, " ")

	for i := 0; i < countNumbers; i++ {
		var num int
		num, err = strconv.Atoi(nums[i])
		if err != nil {
			return nil, err
		}
		res[i] = num
	}

	return res, nil
}

func readLine(reader *bufio.Reader) (string, error) {
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n\r")

	return line, nil
}

func writeResult(writer *bufio.Writer, result string) {
	_, _ = writer.WriteString(result)
}

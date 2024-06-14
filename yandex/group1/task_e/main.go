package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	WrongInputLengthError = "wrong input length"
	CantConvertError      = "can't convert input"
)

func findProfit(n, k, d int64) string {
	remainder := n % k
	result := strconv.FormatInt(n, 10)

	for i := int64(0); i < d; i++ {
		var found bool
		for digit := 0; digit <= 9; digit++ {
			if (remainder*10+int64(digit))%k == 0 {
				result += strconv.Itoa(digit)
				remainder = (remainder*10 + int64(digit)) % k
				found = true
				break
			}
		}
		if !found {
			return "-1"
		}
	}

	return result
}
func main() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	lines, err := getTaskInput(reader)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	writeResult(writer, findProfit(lines[0], lines[1], lines[2]))
}

func getTaskInput(reader *bufio.Reader) ([]int64, error) {
	linesNumber, err := readLine(reader)
	if err != nil {
		return []int64{}, err
	}
	items := strings.Split(linesNumber, " ")
	if len(items) != 3 {
		return []int64{}, errors.New(WrongInputLengthError)
	}

	result := make([]int64, 3)
	for i := 0; i < 3; i++ {
		result[i], err = strconv.ParseInt(items[i], 10, 64)
		if err != nil {
			return []int64{}, errors.New(CantConvertError)
		}
	}

	return result, nil
}

func readLine(reader *bufio.Reader) (string, error) {
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n\r")

	return line, nil
}

func writeResult(writer *bufio.Writer, result string) {
	_, _ = writer.WriteString(result)
}

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()

	count, _ := ReadLinesCount(in)

	for i := 0; i < count; i++ {
		a, b := ReadPair(in)
		_, _ = out.WriteString(strconv.Itoa(Sum(a, b)) + "\n")
	}
}

func Sum(a, b int) int {
	return a + b
}

func ReadLinesCount(reader *bufio.Reader) (int, error) {
	line, _ := reader.ReadString('\n')
	return strconv.Atoi(strings.TrimRight(line, "\n\r"))
}

func ReadPair(reader *bufio.Reader) (int, int) {
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n\r")
	parts := strings.Split(line, " ")
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	return a, b
}

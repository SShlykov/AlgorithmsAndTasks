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

	str, _ := ReadInitString(in)
	letters := []rune(str)

	count, _ := ReadLinesCount(in)

	for i := 0; i < count; i++ {
		start, finish, updateLetters := ReadLine(in)

		for j := start; j <= finish; j++ {
			letters[j-1] = updateLetters[j-start]
		}
	}

	_, _ = out.WriteString(string(letters))
}

func ReadLine(reader *bufio.Reader) (int, int, []rune) {
	line, _ := reader.ReadString('\n')
	items := strings.Split(strings.TrimRight(line, "\n\r"), " ")

	start, _ := strconv.Atoi(items[0])
	finish, _ := strconv.Atoi(items[1])
	letters := []rune(items[2])
	return start, finish, letters
}

func ReadLinesCount(reader *bufio.Reader) (int, error) {
	line, _ := reader.ReadString('\n')
	return strconv.Atoi(strings.TrimRight(line, "\n\r"))
}

func ReadInitString(reader *bufio.Reader) (string, error) {
	line, _ := reader.ReadString('\n')
	return strings.TrimRight(line, "\n\r"), nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()

	usersCount, requestCount := ReadLine(in)
	fmt.Println(usersCount, requestCount)
}

func ReadLine(reader *bufio.Reader) (int, int) {
	line, _ := reader.ReadString('\n')
	slice := strings.Split(strings.TrimRight(line, "\n\r"), " ")
	n1, _ := strconv.Atoi(slice[0])
	n2, _ := strconv.Atoi(slice[1])

	return n1, n2
}

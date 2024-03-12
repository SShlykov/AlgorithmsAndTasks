package task_h

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const CantConvertError = "can't convert input to int"

type TaskParams struct {
	StadiumLength int
	StartPoint1   int
	Speed1        int
	StartPoint2   int
	Speed2        int
}

func main() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	taskParams, err := getTaskParams(reader)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	res := calc(taskParams)

	writeResult(writer, res)
}

func calc(params *TaskParams) string {
	return ""
}

func getTaskParams(reader *bufio.Reader) (*TaskParams, error) {
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n\r")
	items := strings.Split(line, " ")

	var result []int
	for i := 0; i < 5; i++ {
		num, err := strconv.Atoi(items[i])
		if err != nil {
			return nil, errors.New(CantConvertError)
		}
		result = append(result, num)
	}
	taskParams := &TaskParams{
		StadiumLength: result[0],
		StartPoint1:   result[1],
		Speed1:        result[2],
		StartPoint2:   result[3],
		Speed2:        result[4],
	}

	return taskParams, nil

}

func readLine(reader *bufio.Reader) (string, error) {
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n\r")

	return line, nil
}

func writeResult(writer *bufio.Writer, result string) {
	_, _ = writer.WriteString(result)
}

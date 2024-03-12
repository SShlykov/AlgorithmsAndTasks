package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

const (
	InvalidInputConversionError = "invalid input: cant be converted to int"
	InvalidInputParamLength     = "invalid input: slice length is not equal to 2"
)

const (
	AWAY = iota + 1
	HOME
)

func main() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	round1, _ := readRound(reader)
	round2, _ := readRound(reader)
	placement, _ := readPlacement(reader)

	goals := CalculateNeededGoals(round1, round2, placement)

	writeResult(writer, strconv.Itoa(goals))
}

func CalculateNeededGoals(round1, round2 *Round, placement int) int {
	if placement == AWAY {
		score1 := round1.Team1Score*100 + round2.Team1Score*101
		score2 := round1.Team2Score*101 + round2.Team2Score*100
		return max(0, (score2-score1+101)/101)
	} else {
		score1 := round1.Team1Score*101 + round2.Team1Score*100
		score2 := round1.Team2Score*100 + round2.Team2Score*101
		return max(0, (score2-score1+100)/100)
	}
}

type Round struct {
	Team1Score int
	Team2Score int
}

func NewRound(firstSubSlice, secondSubSlice string) (*Round, error) {
	var Team1Score, Team2Score int
	var err error

	Team1Score, err = strconv.Atoi(firstSubSlice)
	if err != nil {
		return nil, errors.New(InvalidInputConversionError)
	}
	Team2Score, err = strconv.Atoi(secondSubSlice)
	if err != nil {
		return nil, errors.New(InvalidInputConversionError)
	}

	return &Round{Team1Score, Team2Score}, nil
}

func readRound(reader *bufio.Reader) (*Round, error) {
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n\r")

	slice := strings.Split(line, ":")
	if len(slice) != 2 {
		return nil, errors.New(InvalidInputParamLength)
	}

	params, err := NewRound(slice[0], slice[1])
	if err != nil {
		return nil, err
	}

	return params, nil
}

func readPlacement(reader *bufio.Reader) (int, error) {
	res, err := reader.ReadByte()
	if err != nil {
		return 0, errors.New(InvalidInputConversionError)
	}
	placement := int(res - '0')
	return placement, nil
}

func writeResult(writer *bufio.Writer, result string) {
	_, _ = writer.WriteString(result)
}

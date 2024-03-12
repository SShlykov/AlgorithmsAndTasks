package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

const (
	InvalidInputConvertionError = "invalid input: cant be converted to int"
	InvalidInputParamLength     = "invalid input: slice length is not equal to 2"
)

const (
	AWAY = iota + 1
	HOME
)

func main() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	round1, err := readRound(reader)
	if err != nil {
		writeResult(writer, err.Error())
		os.Exit(2)
	}
	round2, err := readRound(reader)
	if err != nil {
		writeResult(writer, err.Error())
		os.Exit(2)
	}

	placement, err := readPlacement(reader)
	if err != nil {
		writeResult(writer, err.Error())
		os.Exit(2)
	}

	neededGoals := CalculateNeededGoals(round1, round2, placement)

	writeResult(writer, strconv.Itoa(neededGoals))
}

func CalculateNeededGoals(round1, round2 *Round, placement int) int {
	team1AwayScore, team2AwayScore := GetAwayScores(round1, round2, placement)
	team1TotalScore := round1.Team1Score + round2.Team1Score
	team2TotalScore := round1.Team2Score + round2.Team2Score

	var needToWin int
	switch {
	case team1TotalScore > team2TotalScore:
		needToWin = 0 // win by total goals = [case 1]
	case team1TotalScore == team2TotalScore:
		if team1AwayScore > team2AwayScore {
			needToWin = 0 // win by away goals = [case 2]
		} else {
			needToWin = 1 // -> case 1
		}
	case placement == HOME:
		if team1AwayScore > team2AwayScore {
			needToWin = team2TotalScore - team1TotalScore // we cant win by away goals -> case 2
		} else {
			needToWin = team2TotalScore - team1TotalScore + 1 // we cant win by total goals -> case 1
		}

	default:
		// if away we can add needToWin to team1AwayScore until it becomes more than team2AwayScore -> case 2
		// but not more than team2TotalScore-team1TotalScore+1 -> case 1
		needToWin = min(
			team2TotalScore-team1TotalScore+1,
			max(team2TotalScore-team1TotalScore, team2AwayScore-team1AwayScore+1),
		)
	}

	return needToWin
}

type Round struct {
	Team1Score int
	Team2Score int
}

func GetAwayScores(round1, round2 *Round, placement int) (team1 int, team2 int) {
	if placement == 1 {
		return round2.Team1Score, round1.Team2Score
	}
	return round1.Team1Score, round2.Team2Score
}

func NewRound(firstSubSlice, secondSubSlice string) (*Round, error) {
	var Team1Score, Team2Score int
	var err error

	Team1Score, err = strconv.Atoi(firstSubSlice)
	if err != nil {
		return nil, errors.New(InvalidInputConvertionError)
	}
	Team2Score, err = strconv.Atoi(secondSubSlice)
	if err != nil {
		return nil, errors.New(InvalidInputConvertionError)
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
		return 0, errors.New(InvalidInputConvertionError)
	}
	placement := int(res - '0')
	return placement, nil
}

func writeResult(writer *bufio.Writer, result string) {
	_, _ = writer.WriteString(result)
}

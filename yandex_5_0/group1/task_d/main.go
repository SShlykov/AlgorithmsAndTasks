package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	BoardSize = 8
)

const (
	None = iota
	OnDanger
	Rook
	Bishop
)

type Figure struct {
	Name int
}

type Board struct {
	Places [8][8]Figure
}

func newBoardRow() [8]Figure {
	res := [8]Figure{}
	for i := 0; i < BoardSize; i++ {
		res[i] = Figure{Name: None}
	}
	return res
}

func NewBoard() *Board {
	places := [8][8]Figure{}
	for i := 0; i < BoardSize; i++ {
		places[i] = newBoardRow()
	}

	return &Board{Places: places}
}

func main() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	board, err := getBoard(reader)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	markAttackedCells(board)

	writeResult(writer, strconv.Itoa(countSafeCells(board)))
}

func getBoard(reader *bufio.Reader) (*Board, error) {
	board := NewBoard()

	for i := 0; i < BoardSize; i++ {
		boardLine, err := readLine(reader)
		if err != nil {
			return nil, err
		}

		boardItems := strings.Split(boardLine, "")

		for j := 0; j < BoardSize; j++ {
			if boardItems[j] == "R" {
				board.Places[i][j] = Figure{Name: Rook}
			}
			if boardItems[j] == "B" {
				board.Places[i][j] = Figure{Name: Bishop}
			}
		}
	}

	return board, nil
}

func readLine(reader *bufio.Reader) (string, error) {
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n\r")

	return line, nil
}

func writeResult(writer *bufio.Writer, result string) {
	_, _ = writer.WriteString(result)
}

func countSafeCells(board *Board) int {
	safeCells := 0
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if board.Places[i][j].Name == None {
				safeCells++
			}
		}
	}
	return safeCells
}

func markAttackedCells(board *Board) {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			switch board.Places[i][j].Name {
			case Rook:
				markRookPaths(board, i, j)
			case Bishop:
				markBishopPaths(board, i, j)
			default:
				continue
			}
		}
	}
}

func markRookPaths(board *Board, row, col int) {
	for i := row - 1; i >= 0; i-- {
		if board.Places[i][col].Name > OnDanger {
			break
		}
		board.Places[i][col].Name = OnDanger
	}

	for i := row + 1; i < BoardSize; i++ {
		if board.Places[i][col].Name > OnDanger {
			break
		}
		board.Places[i][col].Name = OnDanger
	}

	for j := col - 1; j >= 0; j-- {
		if board.Places[row][j].Name > OnDanger {
			break
		}
		board.Places[row][j].Name = OnDanger
	}

	for j := col + 1; j < BoardSize; j++ {
		if board.Places[row][j].Name > OnDanger {
			break
		}
		board.Places[row][j].Name = OnDanger
	}
}

func markBishopPaths(board *Board, row, col int) {
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board.Places[i][j].Name > OnDanger {
			break
		}
		board.Places[i][j].Name = OnDanger
	}

	for i, j := row-1, col+1; i >= 0 && j < BoardSize; i, j = i-1, j+1 {
		if board.Places[i][j].Name > OnDanger {
			break
		}
		board.Places[i][j].Name = OnDanger
	}

	for i, j := row+1, col-1; i < BoardSize && j >= 0; i, j = i+1, j-1 {
		if board.Places[i][j].Name > OnDanger {
			break
		}
		board.Places[i][j].Name = OnDanger
	}

	for i, j := row+1, col+1; i < BoardSize && j < BoardSize; i, j = i+1, j+1 {
		if board.Places[i][j].Name > OnDanger {
			break
		}
		board.Places[i][j].Name = OnDanger
	}
}

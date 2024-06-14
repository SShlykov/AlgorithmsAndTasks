package chess

import (
	"fmt"
	"strconv"
)

func main() {
	board := [][]string{
		{"_", "_", "_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "s", "_"},
		{"x", "x", "x", "x", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "x", "x", "x", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_", "_"},
		{"x", "x", "_", "_", "_", "e", "_", "_"},
	}

	start := Point{1, 6}
	end := Point{8, 5}

	path := GetRoute(board, start, end)
	fmt.Println(path)
	for i, point := range path {
		board[point.r][point.c] = strconv.Itoa(i)
	}

	for _, row := range board {
		for _, field := range row {
			fmt.Print(field, "_")
		}
		fmt.Println()
	}
}

type Point struct {
	r, c int
}

func GetRoute(board [][]string, start, end Point) []Point {
	moves := []Point{
		{-2, -1},
		{-1, -2},
		{1, -2},
		{2, -1},
		{2, 1},
		{1, 2},
		{-1, 2},
		{-2, 1},
	}

	rows, cols := len(board), len(board[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	predecessor := make(map[Point]Point)
	queue := []Point{start}
	visited[start.r][start.c] = true
	found := false

	for len(queue) > 0 && !found {
		curr := queue[0]
		queue = queue[1:]

		if curr == end {
			found = true
			break
		}

		for _, move := range moves {
			next := Point{curr.r + move.r, curr.c + move.c}
			if isValid(board, visited, next) {
				queue = append(queue, next)
				visited[next.r][next.c] = true
				predecessor[next] = curr
			}
		}
	}

	if !found {
		return nil
	}

	path := []Point{}
	for p := end; p != start; p = predecessor[p] {
		path = append([]Point{p}, path...)
	}
	path = append([]Point{start}, path...)
	return path
}

func isValid(board [][]string, visited [][]bool, p Point) bool {
	return p.r >= 0 && p.r < len(board) && p.c >= 0 && p.c < len(board[0]) && board[p.r][p.c] != "x" && !visited[p.r][p.c]
}

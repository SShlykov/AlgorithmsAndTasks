package main

// time - O(n * n) n in row and n in col
// memory - O(n*n)
// can be solved O(n^2) + O(n) if only one set created, but we need to run 3 times in cycle
func isValidSudoku(board [][]byte) bool {
	rows := make(map[[2]int]struct{})
	cols := make(map[[2]int]struct{})
	blocks := make(map[[2]int]struct{})

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			val := int(board[i][j])
			if val == '.' {
				continue
			}

			row := [2]int{i, val}
			_, rowExist := rows[row]

			col := [2]int{j, val}
			_, colExist := cols[col]

			blockIDx := 3*(i/3) + j/3
			block := [2]int{blockIDx, val}
			_, blockExist := blocks[block]

			if rowExist || colExist || blockExist {
				return false
			}

			rows[row] = struct{}{}
			cols[col] = struct{}{}
			blocks[block] = struct{}{}
		}
	}

	return true
}

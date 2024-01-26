package zigzag_conversion

import "bytes"

// convert rearranges the characters of a string 's' into a zigzag pattern across 'numRows' rows,
// and then reads the characters line by line.
func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	res := bytes.Buffer{}
	step := numRows*2 - 2

	for i := 0; i < len(s); i += step {
		res.WriteByte(s[i])
	}

	for r := 1; r <= numRows-2; r++ {
		res.WriteByte(s[r])

		for k := step; k-r < len(s); k += step {
			res.WriteByte(s[k-r])
			if k+r < len(s) {
				res.WriteByte(s[k+r])
			}
		}
	}

	for i := numRows - 1; i < len(s); i += step {
		res.WriteByte(s[i])
	}

	return res.String()
}

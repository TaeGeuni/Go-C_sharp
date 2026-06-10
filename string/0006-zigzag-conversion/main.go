package main

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	zigzag := make([]strings.Builder, numRows)

	currRow := 0
	goingDown := true

	for i := 0; i < len(s); i++ {
		if goingDown {
			zigzag[currRow].WriteString(string(s[i]))
			currRow++
		} else {
			zigzag[currRow].WriteString(string(s[i]))
			currRow--
		}

		if currRow == numRows-1 || currRow == 0 {
			goingDown = !goingDown
		}
	}

	var res strings.Builder

	for i := 0; i < numRows; i++ {
		res.WriteString(zigzag[i].String())
	}

	return res.String()
}

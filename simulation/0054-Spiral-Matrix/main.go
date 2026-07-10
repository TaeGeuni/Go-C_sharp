package main

func spiralOrder(matrix [][]int) []int {
	res := make([]int, len(matrix)*len(matrix[0]))
	direction := 'E'
	x, y := 0, 0
	minX, minY, maxX, maxY := 0, 0, len(matrix[0])-1, len(matrix)-1
	for i := 0; i < len(res); i++ {
		if direction == 'E' {
			res[i] = matrix[y][x]
			if x == maxX {
				direction = 'S'
				minY++
				y++
			} else {
				x++
			}
		} else if direction == 'W' {
			res[i] = matrix[y][x]
			if x == minX {
				direction = 'N'
				maxY--
				y--
			} else {
				x--
			}
		} else if direction == 'N' {
			res[i] = matrix[y][x]
			if y == minY {
				direction = 'E'
				minX++
				x++
			} else {
				y--
			}
		} else if direction == 'S' {
			res[i] = matrix[y][x]
			if y == maxY {
				direction = 'W'
				maxX--
				x--
			} else {
				y++
			}
		}
	}

	return res
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func enqueue(q []byte, b byte) []byte {
	q = append(q, b)
	if len(q) > 4 {
		q = q[1:]
	}
	return q
}

func isXmas(q []byte) bool {
	return string(q) == "XMAS" || string(q) == "SAMX"
}

func countHorisontal(matrix [][]byte) int {
	xmasCount := 0
	for r := range matrix {
		q := []byte{}
		for c := range matrix[0] {
			q = enqueue(q, matrix[r][c])
			if isXmas(q) {
				xmasCount++
			}
		}
	}
	return xmasCount
}

func countVertical(matrix [][]byte) int {
	xmasCount := 0
	for c := range matrix[0] {
		q := []byte{}
		for r := range matrix {
			q = enqueue(q, matrix[r][c])
			if isXmas(q) {
				xmasCount++
			}
		}
	}
	return xmasCount
}

func countDiagonal(matrix [][]byte) int {
	xmasCount := 0
	rows, cols := len(matrix), len(matrix[0])

	// diagonal count left '\'
	for r1 := rows - 1; r1 >= 0; r1-- {
		q := []byte{}
		c := 0
		for r := r1; r < rows; r++ {
			q = enqueue(q, matrix[r][c])
			if isXmas(q) {
				xmasCount++
			}
			c++
		}
	}

	// diagonal count right '/'
	for r1 := rows - 1; r1 >= 1; r1-- {
		q := []byte{}
		c := cols - 1
		for r := r1; r < rows; r++ {
			q = enqueue(q, matrix[r][c])
			if isXmas(q) {
				xmasCount++
			}
			c--
		}
	}

	// diagonal count top '\'
	for c1 := 1; c1 < cols; c1++ {
		q := []byte{}
		r := 0
		for c := c1; c < cols; c++ {
			q = enqueue(q, matrix[r][c])
			if isXmas(q) {
				xmasCount++
			}
			r++
		}
	}

	// diagonal count top '/'
	for c1 := range cols {
		q := []byte{}
		r := 0
		for c := c1; c >= 0; c-- {
			q = enqueue(q, matrix[r][c])
			if isXmas(q) {
				xmasCount++
			}
			r++
		}
	}

	return xmasCount
}

func countXmas() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var xmasCount int

	matrix := make([][]byte, 0, 140)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []byte(line))
	}

	xmasCount += countHorisontal(matrix)
	xmasCount += countVertical(matrix)
	xmasCount += countDiagonal(matrix)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return xmasCount
}

func enqueueCross(q []byte, b byte) []byte {
	q = append(q, b)
	if len(q) > 3 {
		q = q[1:]
	}
	return q
}

func isMasLeft(q []byte) bool {
	return string(q) == "MAS" || string(q) == "SAM"
}

func isMasRight(matrix [][]byte, r, c int) bool {
	if matrix[r-1][c+1] == 'S' && matrix[r+1][c-1] == 'M' {
		return true
	}
	if matrix[r-1][c+1] == 'M' && matrix[r+1][c-1] == 'S' {
		return true
	}
	return false
}
func countDiagonalCross(matrix [][]byte) int {
	xmasCount := 0
	rows, cols := len(matrix), len(matrix[0])

	// diagonal count left '\'
	for r1 := rows - 1; r1 >= 0; r1-- {
		q := []byte{}
		c := 0
		for r := r1; r < rows; r++ {
			q = enqueueCross(q, matrix[r][c])
			if isMasLeft(q) && isMasRight(matrix, r-1, c-1) {
				xmasCount++
			}
			c++
		}
	}

	// diagonal count top '\'
	for c1 := 1; c1 < cols; c1++ {
		q := []byte{}
		r := 0
		for c := c1; c < cols; c++ {
			q = enqueueCross(q, matrix[r][c])
			if isMasLeft(q) && isMasRight(matrix, r-1, c-1) {
				xmasCount++
			}
			r++
		}
	}

	return xmasCount
}

func countXmasCross(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var xmasCount int

	matrix := make([][]byte, 0, 140)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []byte(line))
	}

	xmasCount += countDiagonalCross(matrix)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return xmasCount
}

func main() {
	// fmt.Println(countXmas())
	fmt.Println(countXmasCross("input.txt"))
	fmt.Println(countXmasCross("input1.txt"))
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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

	rows, cols := len(matrix), len(matrix[0])

	// horisontal search
	for r := range rows {
		for c := 0; c < cols-5; c++ {
			if string(matrix[r][c:c+4]) == "XMAS" || string(matrix[r][c:c+4]) == "SAMX" {
				xmasCount++
			}
		}
	}

	// vertical search
	q := []byte{}
	for c := range cols {
		for r := 0; r < rows; r++ {
			q = append(q, matrix[r][c])
			if len(q) > 4 {
				q = q[1:]
			}
			if string(q) == "XMAS" || string(q) == "SAMX" {
				xmasCount++
			}
		}
	}

	// diagonal count
	q = []byte{}
	for r1 := rows - 1; r1 >= 0; r1-- {
		c := 0
		for r := r1; r < rows; r++ {
			q = append(q, matrix[r][c])
			if len(q) > 4 {
				q = q[1:]
			}
			if string(q) == "XMAS" || string(q) == "SAMX" {
				xmasCount++
			}
			c++
		}
	}

	if string(q) == "XMAS" || string(q) == "SAMX" {
		xmasCount++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(xmasCount)
}

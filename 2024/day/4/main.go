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
	q := []byte{}

	countXmas := func(r, c int) {
		q = append(q, matrix[r][c])
		if len(q) > 4 {
			q = q[1:]
		}
		if string(q) == "XMAS" || string(q) == "SAMX" {
			xmasCount++
		}
	}

	// horisontal search
	for r := range rows {
		q = []byte{}
		for c := 0; c < cols; c++ {
			countXmas(r, c)
		}
	}

	// vertical search
	for c := range cols {
		q = []byte{}
		for r := 0; r < rows; r++ {
			countXmas(r, c)
		}
	}

	// diagonal count
	for r1 := rows - 1; r1 >= 0; r1-- {
		q = []byte{}
		c := 0
		for r := r1; r < rows; r++ {
			countXmas(r, c)
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

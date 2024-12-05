package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	rows, cols := len(matrix), len(matrix[0])
	q := []byte{}

	updateCount := func(r, c int) {
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
			updateCount(r, c)
		}
	}

	// vertical search
	for c := range cols {
		q = []byte{}
		for r := 0; r < rows; r++ {
			updateCount(r, c)
		}
	}

	// diagonal count left
	for r1 := rows - 1; r1 >= 0; r1-- {
		q = []byte{}
		c := 0
		for r := r1; r < rows; r++ {
			updateCount(r, c)
			c++
		}
	}

	// diagonal count top side
	for c1 := 1; c1 < cols; c1++ {
		q = []byte{}
		r := 0
		for c := c1; c < cols; c++ {
			updateCount(r, c)
			r++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return xmasCount
}

func main() {
	fmt.Println(countXmas())
}

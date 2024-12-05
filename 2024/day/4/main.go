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

	// horisontal search
	for r := range matrix {
		for c := 0; c < len(matrix)-5; c++ {
			if string(matrix[r][c:c+4]) == "XMAS" || string(matrix[r][c:c+4]) == "SAMX" {
				xmasCount++
			}
		}
	}

	// vertical search
	q := []byte{}
	for c := range matrix[0] {
		for r := 0; r < len(matrix); r++ {
			q = append(q, matrix[r][c])
			if len(q) > 4 {
				q = q[1:]
			}
			if string(q) == "XMAS" || string(q) == "SAMX" {
				xmasCount++
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(xmasCount)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func middlePageNumberSum(path1, path2 string) int {
	rules := compileRules(readFile(path1))

	return sum
}

func readFile(path string) [][]byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []byte(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return matrix
}

func main() {
	fmt.Println(middlePageNumberSum("input_part_1.txt", "input_part_2.txt"))
	fmt.Println(middlePageNumberSum("input_mini_part_1.txt", "input_mini_part_2.txt"))
}

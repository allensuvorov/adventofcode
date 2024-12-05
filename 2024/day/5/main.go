package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func compileRules(data []string) [][]int {
	rules := make([][]int, len(data))
	for _, v := range data {
		num1, err := strconv.Atoi(v[:2])
		if err != nil {
			fmt.Println(err)
		}
		num2, err := strconv.Atoi(v[:2])
		if err != nil {
			fmt.Println(err)
		}

		rule := []int{num1, num2}
		rules = append(rules, rule)
	}
	return rules
}

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return data
}

func middlePageNumberSum(path1, path2 string) int {
	rules := compileRules(readFile(path1))
	pages := compilePages(readFile(path2))

	return sum
}

func main() {
	fmt.Println(middlePageNumberSum("input_part_1.txt", "input_part_2.txt"))
	fmt.Println(middlePageNumberSum("input_mini_part_1.txt", "input_mini_part_2.txt"))
}

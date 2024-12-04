package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int

	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line)-9; i++ {
			if line[i:i+3] == "mul" {
				// Check Open Paren
				openParenPos := i + 3
				if openParenPos < len(line) && line[openParenPos] != '(' {
					continue
				}

				// Find number one
				num1Start := openParenPos + 1 // mul(123)
				num1End := num1Start
				if !isDigit(line[num1Start]) {
					continue
				}
				for num1End < len(line) && isDigit(line[num1End]) {
					num1End++
				}

				// Check Comma
				commaPos := num1End
				if commaPos < len(line) && line[commaPos] != ',' {
					continue
				}

				// Find number two
				num2Start := commaPos + 1
				if !isDigit(line[num2Start]) {
					continue
				}
				num2End := num2Start
				for num2End < len(line) && isDigit(line[num2End]) {
					num2End++
				}

				// Check Closed Paren
				closedParenPos := num2End
				if closedParenPos < len(line) && line[closedParenPos] != ')' {
					continue
				}

				// fmt.Print(line[i : closedParenPos+1])

				if num1End < len(line) && num2End < len(line) {
					num1Str := line[num1Start:num1End]
					num2Str := line[num2Start:num2End]

					num1, _ := strconv.Atoi(num1Str)
					num2, _ := strconv.Atoi(num2Str)

					sum += num1 * num2
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

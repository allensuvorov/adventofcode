package main

import (
	"bufio"
	"fmt"
	"os"
)

func getNextDirection(direction int) int {
	return (direction + 1) % 4
}

func getNextPosition(curPos []int) []int {
	r, c, direction := curPos[0], curPos[1], curPos[2]
	move := map[int]func(r, c int) []int{
		0: func(r, c int) []int {
			return []int{r - 1, c}
		},
		1: func(r, c int) []int {
			return []int{r, c + 1}
		},
		2: func(r, c int) []int {
			return []int{r + 1, c}
		},
		3: func(r, c int) []int {
			return []int{r, c - 1}
		},
	}

	return append(move[direction](r, c), direction)
}

func inMap(r, c, rows, cols int) bool {
	return r >= 0 && r <= rows-1 && c >= 0 && c <= cols-1
}

func findGuard(labMap [][]byte) []int {
	for r := range labMap {
		for c := range labMap[0] {
			direction := 0
			if labMap[r][c] != '.' && labMap[r][c] != '#' {
				switch labMap[r][c] {
				case '^':
					direction = 0
				case '>':
					direction = 1
				case 'v':
					direction = 2
				case '<':
					direction = 3
				}
				return []int{r, c, direction}
			}
		}
	}
	return []int{0, 0, 0}
}

func readFile(path string) [][]byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, []byte(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return data
}

func countPositions(path string) {
	labMap := readFile(path)
	rows, cols := len(labMap), len(labMap[0])
	curPos := findGuard(labMap)

	positionCount := 1
	for range 1_000_000 {
		labMap[curPos[0]][curPos[1]] = 'X'
		nextPos := getNextPosition(curPos)
		// fmt.Println(nextPos)

		if !inMap(nextPos[0], nextPos[1], rows, cols) {
			break
		}

		if labMap[nextPos[0]][nextPos[1]] == '#' {
			curPos[2] = getNextDirection(curPos[2])
		} else {
			copy(curPos, nextPos)
			if labMap[nextPos[0]][nextPos[1]] != 'X' {
				positionCount++
			}
		}
	}
	fmt.Println(positionCount)
}

func main() {
	countPositions("sample.txt")
	countPositions("input.txt")
}

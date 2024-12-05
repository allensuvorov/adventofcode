package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isValidOrder(rules [][]int, pages []int) bool {
	pageSet := make(map[int]bool, len(pages))
	pageInd := make(map[int]int, len(pages))

	for _, v := range pages {
		pageSet[v] = true
	}

	for _, v := range rules {
		if pageSet[v[0]] && pageSet[v[1]] {
			pageInd[v[1]]++
		}
	}

	for i := 0; i < len(pages)-1; i++ {
		if pageInd[pages[i]] >= pageInd[pages[i+1]] {
			return false
		}
	}

	return true
}

func fixOrder(rules [][]int, pages []int) {
	pageSet := make(map[int]bool, len(pages))
	pageInd := make(map[int]int, len(pages))

	for _, v := range pages {
		pageSet[v] = true
	}

	for _, v := range rules {
		if pageSet[v[0]] && pageSet[v[1]] {
			pageInd[v[1]]++
		}
	}

	// counting sort
	indPage := make([]int, len(pages))
	for page, ind := range pageInd {
		indPage[ind] = page
	}
	copy(pages, indPage)
}

func compileUpdates(data []string) [][]int {
	updates := make([][]int, 0, len(data))
	for _, line := range data {
		numsStr := strings.Split(line, ",")
		nums := make([]int, len(numsStr))
		for i, s := range numsStr {
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Print(err)
			}
			nums[i] = num
		}
		updates = append(updates, nums)
	}
	return updates
}

func compileRules(data []string) [][]int {
	rules := make([][]int, 0, len(data))
	for _, line := range data {
		num1, err := strconv.Atoi(line[:2])
		if err != nil {
			log.Print(err)
		}
		num2, err := strconv.Atoi(line[3:])
		if err != nil {
			log.Print(err)
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

func middlePageNumberSum(path1, path2 string) {
	rules := compileRules(readFile(path1))
	updates := compileUpdates(readFile(path2))

	initiallyCorrectUpdatesSum := 0
	correctedUpdatesSum := 0
	for _, update := range updates {
		if isValidOrder(rules, update) {
			pos := len(update) / 2
			initiallyCorrectUpdatesSum += update[pos]
		} else {
			fixOrder(rules, update)
			pos := len(update) / 2
			correctedUpdatesSum += update[pos]
		}
	}
	fmt.Printf("sum of initially correct updates: %v\n", initiallyCorrectUpdatesSum)
	fmt.Printf("sum of corrected updates: %v \n", correctedUpdatesSum)

}

func main() {
	middlePageNumberSum("input_part_1.txt", "input_part_2.txt")
	middlePageNumberSum("input_mini_part_1.txt", "input_mini_part_2.txt")
}

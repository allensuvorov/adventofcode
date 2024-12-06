package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func fixOrder(pageInd map[int]int, update []int) {

	// topological sort
	indPage := make([]int, len(update))
	for page, ind := range pageInd {
		indPage[ind] = page
	}
	copy(update, indPage)
}

func isValidOrder(pageInd map[int]int, update []int) bool {

	for i := 0; i < len(update)-1; i++ {
		if pageInd[update[i]] >= pageInd[update[i+1]] {
			return false
		}
	}

	return true
}

func countPageIndegree(rules [][]int, update []int) map[int]int {
	pageSet := make(map[int]bool, len(update))
	pageInd := make(map[int]int, len(update))

	for _, v := range update {
		pageSet[v] = true
	}

	for _, v := range rules {
		if pageSet[v[0]] && pageSet[v[1]] {
			pageInd[v[1]]++
		}
	}
	return pageInd
}

func compileUpdates(data []string) [][]int {
	updates := make([][]int, 0, len(data))
	afterLineBreak := false
	for _, line := range data {
		if line == "" {
			afterLineBreak = true
			continue
		}
		if afterLineBreak {
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
	}
	return updates
}

func compileRules(data []string) [][]int {
	rules := make([][]int, 0, len(data))
	for _, line := range data {
		if line == "" {
			break
		}
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

func middlePageNumberSum(path string) {
	rules := compileRules(readFile(path))
	updates := compileUpdates(readFile(path))

	correctOrderTotal := 0
	fixedOrderTotal := 0

	for _, update := range updates {
		pageInd := countPageIndegree(rules, update)
		if isValidOrder(pageInd, update) {
			pos := len(update) / 2
			correctOrderTotal += update[pos]
		} else {
			fixOrder(pageInd, update)
			pos := len(update) / 2
			fixedOrderTotal += update[pos]
		}
	}
	fmt.Printf("sum of initially correct updates: %v\n", correctOrderTotal)
	fmt.Printf("sum of corrected updates: %v \n", fixedOrderTotal)

}

func main() {
	middlePageNumberSum("sample.txt")
	middlePageNumberSum("input.txt")
}

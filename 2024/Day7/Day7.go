package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var totalEvaluated1 int
	var totalEvaluated2 int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		equation := strings.Split(line, ": ")
		result, _ := strconv.Atoi(equation[0])
		parts := equation[1]

		partsSlice := strings.Split(parts, " ")
		numbersSlice := []int{}
		for _, part := range partsSlice {
			number, _ := strconv.Atoi(part)
			numbersSlice = append(numbersSlice, number)
		}

		if equatesPart1(numbersSlice, result, 0) {
			totalEvaluated1 += result
		}
		if equatesPart2(numbersSlice, result, 0) {
			totalEvaluated2 += result
		}
	}

	fmt.Println("Part 1", totalEvaluated1)
	fmt.Println("Part 2", totalEvaluated2)
}

func equatesPart1(parts []int, target, subresult int) bool {
	if target == subresult {
		return true
	}
	if len(parts) == 0 {
		return false
	}

	return equatesPart1(parts[1:], target, subresult+parts[0]) || equatesPart1(parts[1:], target, subresult*parts[0])
}

func equatesPart2(parts []int, target, subresult int) bool {
	if len(parts) == 0 {
		return target == subresult
	}

	numberString := strconv.Itoa(subresult) + strconv.Itoa(parts[0])
	number, _ := strconv.Atoi(numberString)
	return equatesPart2(parts[1:], target, subresult+parts[0]) ||
		equatesPart2(parts[1:], target, subresult*parts[0]) ||
		equatesPart2(parts[1:], target, number) ||
		equatesPart2(parts[1:], target, number)
}

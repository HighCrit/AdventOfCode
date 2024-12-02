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

	var safeReports int
	var dampenedReports int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())

		if isSafe(values) {
			safeReports++
		} else if isSafeWithDampening(values) {
			dampenedReports++
		}
	}

	fmt.Println("Part 1", safeReports)
	fmt.Println("Part 2", safeReports+dampenedReports)
}

func isSafeWithDampening(values []string) bool {
	for i := 0; i < len(values); i++ {
		newValues := removeIndex(values, i)
		if isSafe(newValues) {
			return true
		}
	}

	return false
}

func removeIndex(a []string, i int) []string {
	newArray := make([]string, len(a))
	copy(newArray, a)
	return append(newArray[:i], newArray[i+1:]...)
}

func isSafe(values []string) bool {
	lastValue, _ := strconv.Atoi(values[0])
	secondValue, _ := strconv.Atoi(values[1])
	increasing := lastValue < secondValue

	for i := 1; i < len(values); i++ {
		value, _ := strconv.Atoi(values[i])

		if (value <= lastValue && increasing) || // The current value is less than or equal to the last value, but we are increasing
			(lastValue <= value && !increasing) || // The current value is more than or equal to the last value, but we are decreasing
			(value > lastValue+3 && increasing) || // The current value is an increase of more than 3
			(lastValue > value+3 && !increasing) { // The current value is a decrease of more than 3
			// The report is UNSAFE
			return false
		}

		lastValue = value
	}

	return true
}

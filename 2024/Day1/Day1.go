package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var leftList = make([]int, 0, 1000)
	var rightList = make([]int, 0, 1000)

	// Read lists from file
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())
		l, _ := strconv.Atoi(values[0])
		r, _ := strconv.Atoi(values[1])

		leftList = append(leftList, l)
		rightList = append(rightList, r)
	}

	// Sort the lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	partOne(leftList, rightList)
	partTwo(leftList, rightList)
}

func partOne(leftList, rightList []int) {
	// Calculate total distance
	var totalDistance int

	for i := 0; i < len(leftList); i++ {
		totalDistance += absDiff(leftList[i], rightList[i])
	}

	fmt.Println("Part 1:", totalDistance)
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func partTwo(leftList, rightList []int) {
	// Calculate similarity score
	var similarityScore int

	currVal := leftList[0]
	var lastI int
	var lastJ int
	for i := 0; i < len(leftList); i++ {
		// If we encounter a new value
		if currVal != leftList[i] {
			// Count the amount of occurence of currVal
			// We can start counting at lastJ since the arrays are sorted
			var occurrencesRightList int
			for lastJ < len(rightList) && currVal >= rightList[lastJ] {
				if currVal == rightList[lastJ] {
					occurrencesRightList++
				}
				lastJ++
			}
			occurrencesLeftList := i - lastI
			similarityScore += currVal * occurrencesLeftList * occurrencesRightList

			// Update values to compare
			currVal = leftList[i]
			lastI = i
		}
	}

	// Process last entry
	lastIndex := len(leftList) - 1
	currVal = leftList[lastIndex]

	var occurrencesRightList int
	for lastJ < len(rightList) && currVal >= rightList[lastJ] {
		if currVal == rightList[lastJ] {
			occurrencesRightList++
		}
		lastJ++
	}
	similarityScore += currVal * (lastIndex - lastI + 1) * occurrencesRightList

	fmt.Println("Part 2:", similarityScore)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	size := 140
	puzzle := make([][]rune, 0, size)
	for scanner.Scan() {
		line := scanner.Text()
		puzzle = append(puzzle, []rune(line))
	}

	var xMasCount int
	var masCount int
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if puzzle[row][col] == 'X' {
				xMasCount += findXmasses(puzzle, row, col)
			}
			if puzzle[row][col] == 'A' && spellsMas(puzzle, row, col) {
				masCount++
			}
		}
	}

	fmt.Println("Part 1", xMasCount)
	fmt.Println("Part 2", masCount)
}

func findXmasses(puzzle [][]rune, row, col int) int {
	var count int
	// Check horizontal LTR
	if spellsChristmas(puzzle, row, col+1, row, col+2, row, col+3) {
		count++
	}
	// Check horizontal RTL
	if spellsChristmas(puzzle, row, col-1, row, col-2, row, col-3) {
		count++
	}
	// Check vertical UTD
	if spellsChristmas(puzzle, row-1, col, row-2, col, row-3, col) {
		count++
	}
	// Check vertical DTU
	if spellsChristmas(puzzle, row+1, col, row+2, col, row+3, col) {
		count++
	}
	// Check diagonal LTR UTD
	if spellsChristmas(puzzle, row-1, col+1, row-2, col+2, row-3, col+3) {
		count++
	}
	// Check diagonal RTL UTD
	if spellsChristmas(puzzle, row-1, col-1, row-2, col-2, row-3, col-3) {
		count++
	}
	// Check diagonal LTR DTU
	if spellsChristmas(puzzle, row+1, col+1, row+2, col+2, row+3, col+3) {
		count++
	}
	// Check diagonal RTL DTU
	if spellsChristmas(puzzle, row+1, col-1, row+2, col-2, row+3, col-3) {
		count++
	}

	return count
}

func spellsChristmas(puzzle [][]rune, mr, mc, ar, ac, sr, sc int) bool {
	size := len(puzzle)
	if mr < 0 || mr >= size || mc < 0 || mc >= size ||
		ar < 0 || ar >= size || ac < 0 || ac >= size ||
		sr < 0 || sr >= size || sc < 0 || sc >= size {
		return false
	}
	return puzzle[mr][mc] == 'M' && puzzle[ar][ac] == 'A' && puzzle[sr][sc] == 'S'
}

func spellsMas(puzzle [][]rune, ar, ac int) bool {
	size := len(puzzle)
	if ar < 1 || ar+2 > size || ac < 1 || ac+2 > size {
		return false
	}

	// Diagonal LTR
	return ((puzzle[ar-1][ac-1] == 'M' || puzzle[ar-1][ac-1] == 'S') &&
		(puzzle[ar+1][ac+1] == 'S' || puzzle[ar+1][ac+1] == 'M') &&
		puzzle[ar+1][ac+1] != puzzle[ar-1][ac-1]) &&
		// Diagonal RTL
		((puzzle[ar-1][ac+1] == 'M' || puzzle[ar-1][ac+1] == 'S') &&
			(puzzle[ar+1][ac-1] == 'S' || puzzle[ar+1][ac-1] == 'M') &&
			puzzle[ar+1][ac-1] != puzzle[ar-1][ac+1])
}

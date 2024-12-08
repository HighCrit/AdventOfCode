package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int

const (
	Undefined Direction = iota
	North
	East
	South
	West
)

type Cell struct {
	obstacle         bool
	initialDirection Direction
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	field := [][]Cell{}
	var guardX, guardY int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := []Cell{}
		for i, c := range line {
			var direction Direction
			if c == '^' {
				direction = North
				guardX = i
				guardY = len(field)
			}
			row = append(row, Cell{obstacle: c == '#', initialDirection: direction})
		}

		field = append(field, row)
	}

	fieldCpy := deepCopyField(field)
	part1, _ := performGuardWalk(fieldCpy, guardX, guardY)

	var possibleLoops int
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			fieldCpy = deepCopyField(field)
			cell := &fieldCpy[y][x]
			if !cell.obstacle && !(x == guardX && y == guardY) {
				cell.obstacle = true
				_, hasLoop := performGuardWalk(fieldCpy, guardX, guardY)
				if hasLoop {
					possibleLoops++
				}
				cell.obstacle = false
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", possibleLoops)
}

func performGuardWalk(field [][]Cell, guardX, guardY int) (int, bool) {
	currentDirection := North
	totalWalked := 1
	hasLoop := false

	for {
		nextX, nextY := getNextIndices(currentDirection, guardX, guardY)
		if !areValidIndices(field, nextX, nextY) {
			break
		}
		cell := &field[nextY][nextX]
		if !cell.obstacle {

			if cell.initialDirection == Undefined {
				totalWalked++
				cell.initialDirection = currentDirection
			} else if cell.initialDirection == currentDirection {
				hasLoop = true
				break
			}
			guardX = nextX
			guardY = nextY
		} else {
			currentDirection = (currentDirection % 4) + 1
		}
	}

	return totalWalked, hasLoop
}

func deepCopyField(field [][]Cell) [][]Cell {
	newField := [][]Cell{}
	for row := 0; row < len(field); row++ {
		newCells := []Cell{}
		for col := 0; col < len(field[row]); col++ {
			oldCell := field[row][col]
			newCells = append(newCells, Cell{obstacle: oldCell.obstacle, initialDirection: oldCell.initialDirection})
		}
		newField = append(newField, newCells)
	}

	return newField
}

func getNextIndices(direction Direction, guardX, guardY int) (int, int) {
	switch direction {
	case North:
		return guardX, guardY - 1
	case East:
		return guardX + 1, guardY
	case South:
		return guardX, guardY + 1
	case West:
		return guardX - 1, guardY
	}

	// Shouldn't happen
	return -1, -1
}

func areValidIndices(field [][]Cell, x, y int) bool {
	return 0 <= y && y < len(field) && 0 <= x && x < len(field[y])
}

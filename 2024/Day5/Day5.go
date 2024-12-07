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

	updates := [][]int{}
	rules := map[int][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.ContainsRune(line, '|') {
			rule := strings.Split(line, "|")
			l, _ := strconv.Atoi(rule[0])
			r, _ := strconv.Atoi(rule[1])
			rules[l] = append(rules[l], r)
		}
		if strings.ContainsRune(line, ',') {
			updatePages := strings.Split(line, ",")
			var update []int
			for _, v := range updatePages {
				pageNumber, _ := strconv.Atoi(v)
				update = append(update, pageNumber)
			}

			updates = append(updates, update)
		}
	}

	var sumOfMiddlePages int
	var sumOfMiddlePagesFaulty int
	for _, update := range updates {
		if isValidUpdate(rules, update) {
			sumOfMiddlePages += update[(len(update)-1)/2]
		} else {
			newUpdate := makeUpdateValid(rules, update)
			sumOfMiddlePagesFaulty += newUpdate[(len(update)-1)/2]
		}
	}

	fmt.Println("Part 1", sumOfMiddlePages)
	fmt.Println("Part 2", sumOfMiddlePagesFaulty)
}

func isValidUpdate(rules map[int][]int, update []int) bool {
	for i := 0; i < len(update); i++ {
		otherPages, present := rules[update[i]]
		if present {
			for j := 0; j < i; j++ {
				for _, otherPage := range otherPages {
					if update[j] == otherPage {
						return false
					}
				}
			}
		}
	}
	return true
}

func makeUpdateValid(rules map[int][]int, update []int) []int {
	newUpdate := make([]int, len(update))
	copy(newUpdate, update)

	for i := 0; i < len(newUpdate); i++ {
		otherPages, present := rules[newUpdate[i]]
		if present {
			for j := 0; j < i; j++ {
				for _, otherPage := range otherPages {
					if newUpdate[j] == otherPage {
						temp := newUpdate[j]
						newUpdate[j] = newUpdate[i]
						newUpdate[i] = temp
						return makeUpdateValid(rules, newUpdate)
					}
				}
			}
		}
	}

	return newUpdate
}

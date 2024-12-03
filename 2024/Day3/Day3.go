package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var input string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input += scanner.Text()
	}

	r := regexp.MustCompile(`(mul\((?P<left>\d{1,3}),(?P<right>\d{1,3})\))|(do\(\))|(don't\(\))`)
	matches := r.FindAllStringSubmatch(input, -1)

	enabled := true
	var total1 int
	var total2 int

	for _, v := range matches {
		if strings.HasPrefix(v[0], "mul") {
			left, _ := strconv.Atoi(v[2])
			right, _ := strconv.Atoi(v[3])
			total1 += left * right

			if enabled {
				total2 += left * right
			}
		} else if v[0] == "do()" {
			enabled = true
		} else if v[0] == "don't()" {
			enabled = false
		}
	}

	fmt.Println("Part 1", total1)
	fmt.Println("Part 2", total2)
}

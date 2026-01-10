package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/migvas/aoc-2025/internal/utils"
)

func main() {
	lines, err := utils.ReadLines("./input/d6.txt")

	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	grandTotal := 0

	var matrix [][]string

	// Create a matrix with all the values and sanitize the lines
	for _, v := range lines {
		matrix = append(matrix, strings.Fields(v))
	}

	// Get line idx of operatores
	operatorIdx := len(matrix) - 1

	// Loop by column then by line
	for i := 0; i < len(matrix[0]); i++ {
		operator := matrix[operatorIdx][i]
		// Add the first element to a col counter for possible multiplication
		colCounter, err := strconv.Atoi(matrix[0][i])
		if err != nil {
			log.Fatalf("failed to parse string number: %v", err)
		}

		for j := 1; j < len(matrix)-1; j++ {
			element, err := strconv.Atoi(matrix[j][i])

			if err != nil {
				log.Fatalf("failed to parse string number: %v", err)
			}
			if operator == "+" {
				colCounter += element
			} else if operator == "*" {
				colCounter *= element
			}
		}
		// Sum the col counter to the grand total
		grandTotal += colCounter
	}
	fmt.Printf("Part 1 - Grand total: %d\n", grandTotal)
}

func part2(lines []string) {
	grandTotal := 0

	// Create a slice for operation elements
	var elements []int

	// Loop the cols from right to left
	for j := len(lines[0]) - 1; j >= 0; j-- {
		// Var to save the number that is saved in the col
		colNumber := ""

		// Get the number from the col
		for i := 0; i < len(lines)-1; i++ {
			if string(lines[i][j]) != " " {
				colNumber += string(lines[i][j])
			}
		}
		// If the col is only whitespaces then skip
		if len(colNumber) == 0 {
			continue
		}

		// Convert to int
		colInt, err := strconv.Atoi(colNumber)

		if err != nil {
			log.Fatalf("failed to parse string number: %v", err)
		}

		// Append to the operation elements slice
		elements = append(elements, colInt)

		// If there is an operator in the last line then we can compute the operation
		if string(lines[len(lines)-1][j]) != " " {
			operator := string(lines[len(lines)-1][j])

			colCounter := elements[0]
			for k := 1; k < len(elements); k++ {
				if operator == "+" {
					colCounter += elements[k]
				} else if operator == "*" {
					colCounter *= elements[k]
				}
			}
			// Empty the elements array for the next operation
			elements = elements[:0]

			// Sum the result to grand total
			grandTotal += colCounter
		}

	}
	fmt.Printf("Part 2 - Grand total: %d\n", grandTotal)
}

package main

import (
	"fmt"
	"log"

	"github.com/migvas/aoc-2025/internal/utils"
)

func main() {
	lines, err := utils.ReadLines("./input/d4.txt")

	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	part1(lines)
	part2(lines)
}

func part1(grid []string) {
	accessibleRolls := 0
	// Go through every position on the grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '@' && isAccessible(grid, i, j) {
				accessibleRolls++
			}
		}
	}
	fmt.Printf("Part 1 - Accessible Rols: %d\n", accessibleRolls)
}

func part2(grid []string) {
	totalRemoved := 0

	// We loop untill nothing is removed anymore
	for {
		// Set to save all the postions where the roll is removed
		removedPositionsSet := make(map[[2]int]struct{})

		// Check which rolls can be removed and save the positions
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == '@' && isAccessible(grid, i, j) {
					removedPositionsSet[[2]int{i, j}] = struct{}{}
				}
			}
		}

		// If nothing is removed we stop the loop
		if len(removedPositionsSet) == 0 {
			break
		}

		totalRemoved += len(removedPositionsSet)

		// Switch the char in the grid to represent a removed roll
		for position := range removedPositionsSet {
			aux := []rune(grid[position[0]])
			aux[position[1]] = '.'
			grid[position[0]] = string(aux)
		}
	}
	fmt.Printf("Part 2 - Removed Rols: %d\n", totalRemoved)
}

func isAccessible(grid []string, row int, col int) bool {
	adjacentRolls := 0
	// Access all 8 adjacent positions
	for i := row - 1; i <= row+1; i++ {
		if i < 0 || i >= len(grid) {
			// Outside of the grid, ignore
			continue
		}
		for j := col - 1; j <= col+1; j++ {
			if j < 0 || j >= len(grid[i]) {
				// Outside of the grid, ignore
				continue
			}

			if i == row && j == col {
				// The roll that is being evaluated, not an adjacent one
				continue
			}

			if grid[i][j] == '@' {
				adjacentRolls++
			}

			if adjacentRolls >= 4 {
				return false
			}
		}
	}

	return true
}

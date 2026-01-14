package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/migvas/aoc-2025/internal/utils"
)

func main() {
	lines, err := utils.ReadLines("./input/d7.txt")

	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	part1(lines)
	part2(lines)
}

// Function that goes through the grid and finds the starting position coordinates
func findStartingPos(grid [][]string) ([2]int, error) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "S" {
				return [2]int{i, j}, nil
			}
		}
	}

	return [2]int{}, errors.New("Starting position not found")
}

func part1(lines []string) {
	// Counter for the splits
	splits := 0

	var grid [][]string

	// Create a grid with all the values and sanitize the lines
	for _, v := range lines {
		grid = append(grid, strings.Split(v, ""))
	}

	// Get the starting position coordinates
	startingPos, err := findStartingPos(grid)

	if err != nil {
		log.Fatalf("Starting position not found in grid")
	}

	// Start BFS
	// Initiaze a queue of positions to process
	Q := [][2]int{startingPos}
	// Initialize a map to store visited positions
	visited := make(map[[2]int]bool)

	visited[startingPos] = true

	for {

		// Stopping criteria
		if len(Q) == 0 {
			break
		}

		ci, cj := Q[0][0], Q[0][1]

		Q = Q[1:]

		// The next position is always one line down
		// If the next position is outside the grid skip
		if ci+1 >= len(grid) {
			continue
		}

		nextPos := [2]int{ci + 1, cj}

		// If the next position is an open spot
		if grid[ci+1][cj] == "." {
			// If not yet visited then mark as visited and add it to the queue
			if _, exists := visited[nextPos]; !exists {
				Q = append(Q, nextPos)
				visited[nextPos] = true
			}
		} else {
			// If not an empty spot then the beam splits
			// Increment the counter
			splits += 1

			nextCols := []int{-1, 1}

			for _, v := range nextCols {
				// For each new position check if the col is inside the grid
				if cj+v < 0 || cj+v >= len(grid[ci+1]) {
					continue
				}

				nextPos = [2]int{ci + 1, cj + v}

				// If not visited add to queue and mark as visited
				if _, exists := visited[nextPos]; !exists {
					Q = append(Q, nextPos)
					visited[nextPos] = true
				}
			}
		}
	}

	fmt.Printf("Part 1 - Number of splits: %d\n", splits)
}

func part2(lines []string) {
	var grid [][]string

	// Create a grid with all the values and sanitize the lines
	for _, v := range lines {
		grid = append(grid, strings.Split(v, ""))
	}

	// Get the starting position coordinates
	startingPos, err := findStartingPos(grid)

	if err != nil {
		log.Fatalf("Starting position not found in grid")
	}

	// Create a map to use as cache for the DFS
	// If we already have a number of paths for a certain position then
	// we don't need to compute it again
	cache := make(map[[2]int]int)
	// Declare the DFS function
	var timelinesDFS func(grid [][]string, currentPosition [2]int) int

	timelinesDFS = func(grid [][]string, currentPosition [2]int) int {
		// Check if the number of paths for the current position is in cache
		if v, exists := cache[currentPosition]; exists {
			// If it is then just return the number of paths
			return v
		}
		ci, cj := currentPosition[0], currentPosition[1]

		// Check if next position (1 spot down) is inside the grid
		if ci+1 >= len(grid) {
			// If we are in the last line then just return 1 since the path is over
			cache[currentPosition] = 1
			return 1
		}

		// If the next position is an open spot then recursively call the function for the next spot
		if grid[ci+1][cj] == "." {
			return timelinesDFS(grid, [2]int{ci + 1, cj})
		}

		// If there is a split then compute the 2 next positions
		leftPos := [2]int{ci + 1, cj - 1}
		rightPos := [2]int{ci + 1, cj + 1}
		// Create a variable to hold the sum of paths for the 2 positions
		possiblePaths := 0

		// If the column of the position is inside the grid add the number of paths
		// produced from that position to possiblePaths
		if leftPos[1] >= 0 {
			possiblePaths += timelinesDFS(grid, leftPos)
		}

		if rightPos[1] < len(grid[ci+1]) {
			possiblePaths += timelinesDFS(grid, rightPos)
		}
		// Cache the result of the sum and return it
		cache[currentPosition] = possiblePaths
		return possiblePaths
	}
	numberOfTimelines := timelinesDFS(grid, startingPos)
	fmt.Printf("Part 2 - Number of timelines: %d\n", numberOfTimelines)
}

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/migvas/aoc-2025/internal/utils"
)

func main() {
	lines, err := utils.ReadLines("./input/d1.txt")

	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	part1(lines)
	part2(lines)
}

func part1(rotations []string) {
	zeroes := 0
	dialPos := 50

	for _, l := range rotations {
		direction := string(l[0])
		turns, err := strconv.Atoi(l[1:])
		if err != nil {
			log.Fatalf("failed to parse string input: %v", err)
		}

		if direction == "R" {
			dialPos += turns
		} else {
			dialPos -= turns
		}

		dialPos = dialPos % 100

		if dialPos == 0 {
			zeroes++
		}
	}

	fmt.Printf("Part 1 - Number of zero positions: %d\n", zeroes)
}

func part2(rotations []string) {
	zeroes := 0
	dialPos := 50

	for _, l := range rotations {
		direction := string(l[0])
		turns, err := strconv.Atoi(l[1:])
		startPos := dialPos

		if err != nil {
			log.Fatalf("failed to parse string input: %v", err)
		}

		if direction == "R" {
			dialPos += turns
		} else {
			dialPos -= turns
		}

		fullRotations := dialPos / 100

		if fullRotations < 0 {
			fullRotations = -fullRotations
		}

		// The product will check if both positions have different signs
		// If they do then we passed 0 in between
		if dialPos*startPos < 0 {
			fullRotations++
		}

		dialPos = dialPos % 100

		// Other zeros come from rotating the dial multiple times to the left or right
		if fullRotations > 0 {
			zeroes += fullRotations
		} else if dialPos == 0 {
			// Special case, if the dial stops on zero then we have to add one ore since the
			// rotation counter doesn't add this scenario
			zeroes++
		}
	}

	fmt.Printf("Part 2 - Number of zero positions: %d\n", zeroes)
}

package main

import (
	"fmt"
	"log"
	"math"

	"github.com/migvas/aoc-2025/internal/utils"
)

func main() {
	lines, err := utils.ReadLines("input/d3.txt")

	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	outputJoltage := 0
	for _, bank := range lines {
		largestJolt := 0
		secondLargestJolt := 0

		// We loop backwards from the second to last digit
		for i := len(bank) - 2; i >= 0; i-- {
			battery := int(bank[i] - '0')

			// This finds the largest and second largest jolt
			if battery >= largestJolt {
				secondLargestJolt = largestJolt
				largestJolt = battery
			}
		}

		// The last digit cannot be the largest jolt but we need to check if it is the second
		secondLargestJolt = max(int(bank[len(bank)-1]-'0'), secondLargestJolt)

		outputJoltage += largestJolt*10 + secondLargestJolt
	}

	fmt.Printf("Part 1 - Output Joltage: %d\n", outputJoltage)
}

func part2(lines []string) {
	outputJoltage := 0
	// This also solves part 1 if this num is changed to 2
	numBatteries := 12
	for _, bank := range lines {
		// Need to fix the idx where the prev digit was set
		limitIdx := -1
		for b := numBatteries - 1; b >= 0; b-- {
			// For each battery we want to find the largest jolt and the idx chosen
			largestJolt := 0
			largestJoltIdx := 0
			// We loop from the last digit where are still digits left to complete the target numBatteries
			// Untill the last chosen battery
			for i := len(bank) - 1 - b; i > limitIdx; i-- {
				battery := int(bank[i] - '0')

				if battery >= largestJolt {
					largestJolt = battery
					largestJoltIdx = i
				}
			}

			// We multiply the largest jolt by 10 to the power of the digit position
			outputJoltage += largestJolt * int(math.Pow10(b))
			limitIdx = largestJoltIdx
		}
	}

	fmt.Printf("Part 2 - Output Joltage: %d\n", outputJoltage)
}

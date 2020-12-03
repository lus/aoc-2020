package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Read the data of the input file
	fileData, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Panicln(err)
		return
	}

	// Split the grid into rows
	rows := strings.Split(string(fileData), "\r\n")

	// Define the slopes to check
	slopes := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{7, 2},
	}

	// Calculate and print the result
	result := 1
	for _, slope := range slopes {
		result *= getTreesInSlope(rows, slope[0], slope[1])
	}
	log.Printf("The result is %d.\n", result)
}

// getTreesInSlope returns the amount of trees with the given slope
func getTreesInSlope(rows []string, right, down int) int {
	next := 0
	step := 0
	trees := 0
	for current, row := range rows {
		// Check if this is the next step
		if current != next {
			continue
		}

		// Define the string index to use
		index := step * right
		for index >= len(row) {
			index -= len(row)
		}

		// Check if the character in this step is a tree
		if string(row[index]) == "#" {
			trees++
		}

		// Set the next step to execute and increase the step index
		step++
		next += down
	}
	return trees
}
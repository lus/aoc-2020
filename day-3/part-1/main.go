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

	// Count the trees the toboggan experiences on its way
	trees := 0
	for step, row := range rows {
		// Define the string index to use
		index := step * 3
		for index >= len(row) {
			index -= len(row)
		}

		// Check if the character in this step is a tree
		if string(row[index]) == "#" {
			trees++
		}
	}
	log.Printf("There are %d trees on the way.\n", trees)
}
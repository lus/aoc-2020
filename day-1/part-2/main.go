package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Create a map that maps integers to booleans
	// This is used to have an easy opportunity to check if a specific value is present in the input
	inputs := make(map[int]bool)

	// Read the data of the input file
	fileData, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Panicln(err)
		return
	}

	// Add all input numbers to the input map
	for _, raw := range strings.Split(string(fileData), "\r\n") {
		parsed, err := strconv.Atoi(raw)
		if err != nil {
			log.Println(fmt.Sprintf("Failed to parse '%s' into an integer. Skipping.", raw))
			continue
		}
		inputs[parsed] = true
	}

	// Loop through all input numbers and check if the two required summands also exist
	for first, _ := range inputs {
		// Continue if the first number already exceeds or reaches the wanted sum
		if first >= 2020 {
			continue
		}

		// Search the second number
		for second, _ := range inputs {
			// Continue if the first two numbers already exceed or reach the wanted sum
			if first + second >= 2020 {
				continue
			}

			// Search the third number
			for third, _ := range inputs {
				// Check if the sum of all three numbers is equal to 2020
				if first + second + third == 2020 {
					log.Println(fmt.Sprintf("The three summands are '%d', '%d' and '%d'. Their product is '%d'.", first, second, third, first * second * third))
					return
				}
			}
		}
	}
	log.Println("Could not find a suitable result.")
}

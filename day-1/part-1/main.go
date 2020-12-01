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

	// Loop through all input numbers and check if the corresponding summand also exists
	for number, _ := range inputs {
		if inputs[2020 - number] {
			log.Println(fmt.Sprintf("The two summands are '%d' and '%d'. Their product is '%d'.", number, 2020 - number, number * (2020 - number)))
			return
		}
	}
	log.Println("Could not find a suitable result.")
}

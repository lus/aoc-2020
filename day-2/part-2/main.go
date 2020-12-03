package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Compile the regex a password entry has to match
	regex := regexp.MustCompile("(\\d+)-(\\d+)\\s(.):\\s(.*)")

	// Read the data of the input file
	fileData, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Panicln(err)
		return
	}

	// Count the valid password entries
	valid := 0
	for _, raw := range strings.Split(string(fileData), "\r\n") {
		// Check if the raw string matches the required entry syntax
		if !regex.MatchString(raw) {
			log.Printf("%s does not match the required entry syntax!\n", raw)
			continue
		}

		// Parse the entry and check if it is valid
		groups := regex.FindStringSubmatch(raw)
		entry := &entry{
			Positions:     [2]int{
				mustInt(groups[1]),
				mustInt(groups[2]),
			},
			Character: rune(groups[3][0]),
			Input:     groups[4],
		}
		if entry.isValid() {
			valid++
		}
	}
	log.Printf("There are %d valid entries.\n", valid)
}

// entry represents a password entry
type entry struct {
	Positions [2]int
	Character rune
	Input string
}

// isValid checks if a password entry is valid
func (entry *entry) isValid() bool {
	substring := string(entry.Input[entry.Positions[0] - 1]) + string(entry.Input[entry.Positions[1] - 1])
	if strings.Count(substring, string(entry.Character)) != 1 {
		return false
	}
	return true
}

// mustInt forcefully parses a string into an integer
func mustInt(str string) int {
	parsed, _ := strconv.Atoi(str)
	return parsed
}

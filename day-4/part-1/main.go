package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	// Read the data of the input file
	fileData, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Panicln(err)
		return
	}

	// Define the possible entries and if they are mandatory or not
	possibleEntries := map[string]bool{
		"byr": true,
		"iyr": true,
		"eyr": true,
		"hgt": true,
		"hcl": true,
		"ecl": true,
		"pid": true,
		"cid": false,
	}

	// Split the input data into passport sections
	passports := strings.Split(string(fileData), "\r\n\r\n")

	// Define the RegEx which separates passport entries
	regex := regexp.MustCompile("[\\s|\\n]+")

	// Check the validity of all passports and count the valid ones
	validCount := 0
	for _, passport := range passports {
		// Parse the entries of this passport into a map
		rawEntries := regex.Split(passport, -1)
		entries := make(map[string]string)
		for _, rawEntry := range rawEntries {
			split := strings.Split(rawEntry, ":")
			entries[split[0]] = split[1]
		}

		// Check if all mandatory entries are set
		valid := true
		for possibleEntry, mandatory := range possibleEntries {
			if mandatory && entries[possibleEntry] == "" {
				valid = false
			}
		}

		// Increase the valid entry count if this one is valid
		if valid {
			validCount++
		}
	}
	log.Printf("There are %d valid passports.", validCount)
}
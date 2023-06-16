package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mdhemmings/anagramalyser/cmd/anagram"
)

func main() {
	var checkedWords []string
	filepath := "./testfile.txt"
	// First open the file on disk
	file, err := os.Open(filepath)
	// pass over any errors to our checkError function to save on lines
	checkError(err)
	// remember to close the file eventually
	defer file.Close()
	// we're going to need a byte array for our buffer
	buf := []byte{}
	// we'll use bufio to read the file because buffers
	scanner := bufio.NewScanner(file)
	// set some kind of reasonable default buffer size
	scanner.Buffer(buf, 2048*1024)
	lineNumber := 1
	// iterate over our file using our scanner with it's lovely buffer
	for scanner.Scan() {
		// we'll need a variable to see what word we're checking currently
		var checkword string
		// read in the line from the file
		line := scanner.Text()
		for _, word := range strings.Fields(line) {
			// for every identified word (using strings.Fields) we will want to ...
			// first make a note of what word we're checking
			checkword = word
			// check we haven't already checked this word using our stringExists function, if we have then break this iteration
			if stringExists(checkedWords, checkword) {
				break
			}
			// assuming we haven't already checked this word then pass over to the CheckAnagrams function to check each line for anagrams
			anagrams := anagram.CheckAnagrams(checkword, strings.Fields(line))
			// if we find some anagrams then let the user know
			if len(anagrams) > 0 {
				fmt.Printf("Found anagrams for %v are %v\n", checkword, anagrams)
				// or if we dont then also let them know
			} else {
				fmt.Printf("No anagrams found for %v", checkword)
			}
			// increment our line number each time
			lineNumber++
			// update which words we've actually checked
			checkedWords = append(checkedWords, checkword)
		}
	}
}

func checkError(err error) {
	// straightforward way of not having to repeatedly type the following every time we need to handle an error
	// which in go tends to be a lot because not using a variable will make the compiler sad
	if err != nil {
		log.Fatalf("Error - %v", err)
	}
}

func stringExists(arr []string, target string) bool {
	// loop over each element in the array and check if it matches the checkword we've been passed, if it does then we'll return true
	for _, element := range arr {
		if element == target {
			return true
		}
	}
	// and if not then we'll return false
	return false
}

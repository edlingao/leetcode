package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/edlingao/leet-code/interfaces"
	"github.com/edlingao/leet-code/problems"
)

func main() {
	args := os.Args[1]
	expected := os.Args[2]
	inputFile := os.Args[3]
	output := ""
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	inputString := string(input[:])

	switch args {
	case "numofsubstrings":
		output = strconv.Itoa(problems.NumSub(inputString))
	case "bitChars":
		var inputJSON interfaces.BitProblem
		err := json.Unmarshal(input, &inputJSON)
		if err != nil {
			log.Fatal("ERROR", err)
		}
		output = strconv.FormatBool(problems.IsOneBitCharacter(inputJSON.Input))
	default:
		log.Println("Invalid argument")
	}

	if expected == output {
		// Use emojis
		log.Println("✅ Output matches expected:", output)
	} else {
		log.Println("❌ Output does not match expected. Got:", output, "Expected:", expected)
	}
}

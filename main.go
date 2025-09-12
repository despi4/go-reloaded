package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	goreloaded "go-reloaded/auto-correction"
)

// (hex), (bin) numeric cmd's
// (up), (low), (cap) alpha cmd's
// a, i, o, e, u (a), other (an)
// , . ! ? : ; ... !? ''

func main() {
	// Take arguments from console
	args := os.Args[1:]

	// Checkintg Input
	CheckInput(args)

	inputFileName, outputFileName := args[0], args[1]

	// Read existing input file. Creating new [.txt] file and send editing text for new [.txt] file
	Result(inputFileName, outputFileName)
}

func CheckInput(input []string) {
	// Checking input for having two arguments
	if len(input) == 2 {
		firstFileName, secondFileName := input[0], input[1]

		// Checking names of [.txt] files
		if firstFileName[len(firstFileName)-4:] == ".txt" && secondFileName[len(secondFileName)-4:] == ".txt" {
			_, err := os.Stat(firstFileName)
			if err != nil {
				log.Fatal("Input file not exist\n", err)
			}
		}
	} else {
		log.Fatal("Write two [.txt] files, for example | [inputfile.txt] [outputfile.txt]")
	}
}

func Result(inputFileName, outputFileName string) {
	// Open input file for take text
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal("Reading file is unachievable\n", err)
	}

	// Use built-in package for scanning text line by line
	scanner := bufio.NewScanner(inputFile)

	// Creating file for correctin text
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal("Error of creating file\n", err)
	}
	defer inputFile.Close()
	defer outputFile.Close()

	// Editing text
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			outputFile.Write([]byte(line + "\n"))
			continue
		}

		line = goreloaded.ClearHexBin(line)
		line = goreloaded.BinToDec(line)
		line = goreloaded.HexToDec(line)
		line = goreloaded.AlphaUp(line)

		outputFile.Write([]byte(line + "\n"))
	}
}

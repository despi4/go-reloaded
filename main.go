package main

import (
	"log"
	"os"

	goreloaded "go-reloaded/auto-correction"
)

// (hex), (bin) numeric cmd's
// (up), (low), (cap) alpha cmd's

func main() {
	args := os.Args[1:]

	text := ReadText(CheckInput(args))
	goreloaded.ClearHexBin(text)
}

func CheckInput(input []string) string {
	fileName := ""

	if len(input) == 2 {
		firstFileName, secondFileName := input[0], input[1]

		if firstFileName[len(firstFileName)-4:] == ".txt" && secondFileName[len(secondFileName)-4:] == ".txt" {
			_, err := os.Stat(firstFileName)
			if err != nil {
				log.Fatal("Input file not exist\n", err)
			}

			fileName = firstFileName
		}
	} else {
		log.Fatal("Write two [.txt] files, for example | [inputfile.txt] [outputfile.txt]")
	}

	return fileName
}

func ReadText(fileName string) string {
	text, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Reading file is unachievable\n", err)
	}

	return string(text)
}

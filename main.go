package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
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

// func CheckInput(input []string) {
// 	// Checking input for having two arguments
// 	if len(input) == 2 {
// 		firstFileName, secondFileName := input[0], input[1]

// 		// Checking names of [.txt] files
// 		if filepath.Ext(firstFileName) == ".txt" && filepath.Ext(secondFileName) == ".txt" {
// 			_, err := os.Stat(firstFileName)
// 			if err != nil {
// 				log.Fatal("Input file not exist\n", err)
// 			}
// 		}
// 	} else {
// 		log.Fatal("Write two [.txt] files, for example | [inputfile.txt] [outputfile.txt]")
// 	}
// }

func CheckInput(input []string) {
	if len(input) != 2 {
		log.Fatal("Write two [.txt] files, for example: [inputfile.txt] [outputfile.txt]")
	}

	inputPath, outputPath := input[0], input[1]

	// Проверка расширений
	if filepath.Ext(inputPath) != ".txt" || filepath.Ext(outputPath) != ".txt" {
		log.Fatal("Both files must have .txt extension")
	}

	// Проверка существования входного файла
	info, err := os.Stat(inputPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Input file does not exist:", inputPath)
		}
		log.Fatal("Error accessing input file:", err)
	}

	// Убедимся, что это файл, а не директория
	if info.IsDir() {
		log.Fatal("Input path is a directory, not a file:", inputPath)
	}

	// (Опционально) Проверка, что входной и выходной файлы — не одно и то же
	if inputPath == outputPath {
		log.Fatal("Input and output files must be different")
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

		line = goreloaded.Navigator(line)

		outputFile.Write([]byte(line + "\n"))
	}
}

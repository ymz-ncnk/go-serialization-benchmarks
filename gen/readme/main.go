package main

import (
	"log"
	"os"
)

// This package contains AI generated code.
func main() {
	file, err := os.Open("results/benchstat.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 1. Parse the file content
	rawData, err := Parse(file)
	if err != nil {
		log.Fatalf("Error parsing file: %v", err)
	}

	// 2. Process the raw data
	proc := NewProcessor()
	safeTable, unsafeTable, err := proc.Process(rawData)
	if err != nil {
		log.Fatalf("Error processing data: %v", err)
	}

	generator := NewReadmeGenerator()
	if err := generator.Generate(safeTable, unsafeTable); err != nil {
		log.Fatalf("Failed to generate README: %v", err)
	}
}

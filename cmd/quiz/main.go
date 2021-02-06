package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

func main() {
	csvPtr := flag.String("csv", "problems", "file used for quiz must be in csvs directory")
	flag.Parse()

	var score, maxScore int
	absPath, err := filepath.Abs("../../problems.csv")
	if *csvPtr != "problems" {
		absPath, err = filepath.Abs("../../" + *csvPtr + ".csv")
	}
	if err != nil {
		fmt.Errorf("file not found")
		return
	} else {
		score, maxScore = quiz(absPath)
	}

	fmt.Printf("Congratulations! You scored %d out of %d.\n", score, maxScore)
}

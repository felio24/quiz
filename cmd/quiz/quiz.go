package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func quiz(absPath string) (int, int) {
	fmt.Println("Quiz starting...")

	csvFile, err := os.Open(absPath)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	score := 0
	reader := bufio.NewReader(os.Stdin)
	for i, line := range csvLines {
		fmt.Printf("Problem #%d: %s = ", i + 1, line[0])
		userAnsStr, _ := reader.ReadString('\n')
		userAnsStr = strings.Replace(userAnsStr, "\n", "", -1)
		userAnsStr = strings.TrimSpace(userAnsStr)
		userAnsInt, err := strconv.Atoi(userAnsStr)
		correctAns, err2 := strconv.Atoi(line[1])
		if err != nil && err2 != nil {
			log.Fatal(err)
		}
		if userAnsInt != correctAns {
			break
		}
		score++
	}
	return score, len(csvLines)
}
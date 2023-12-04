package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

// 92801 too high

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	var allLines []string
	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
	}

	sum := 0
	// p2
	// totalCardsWon := 0
	numbersOfWinMap := make(map[int]int)
	for _, line := range allLines {
		colonIdx := strings.Index(line, ":")
		gameNumberString := line[colonIdx-3 : colonIdx]
		gameNumber := strings.Trim(gameNumberString, " ")
		gameDigit, err := strconv.Atoi(gameNumber)
		if err != nil {
			continue
		}

		winningNumberCount := calculateWinningNumberCount(line)
		sum += doubleXTimes(winningNumberCount)
		numbersOfWinMap[gameDigit] = winningNumberCount
		// // PART 2
		// copiesWon := calculateCopiesWon(winningNumberCount, i, allLines)
		// totalCardsWon += copiesWon
	}
	fmt.Print(numbersOfWinMap)
}

func calculateWinningNumberCount(line string) int {
	colonIdx := strings.Index(line, ":")
	separatorIdx := strings.Index(line, "|")

	winningNumsString := line[colonIdx+1 : separatorIdx-1]
	revealedNumsString := line[separatorIdx+1:]

	winningNumsSlice := strings.Split(winningNumsString, " ")
	revealedNumsSlice := strings.Split(revealedNumsString, " ")

	winningNumCount := 0
	for _, num := range revealedNumsSlice {
		if num == " " || num == "" {
			continue
		}
		if slices.Contains(winningNumsSlice, num) {
			winningNumCount++
		}
	}

	return winningNumCount
}

func calculateCopiesWon(winningNumCount int, currentLineIdx int, allLines []string) int {
	totalWon := 0
	for i := currentLineIdx; i <= currentLineIdx+winningNumCount; i++ {
		if i == len(allLines)-1 {
			break
		}
		count := calculateWinningNumberCount(allLines[i])
		copies := calculateCopiesWon(count, i+1, allLines)
		totalWon += copies
	}
	return totalWon
}

func doubleXTimes(x int) int {
	if x == 0 {
		return 0
	}
	value := 1
	for i := 1; i < x; i++ {
		value *= 2
	}
	return value
}

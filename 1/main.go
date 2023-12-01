package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

type number struct {
	value        string
	writtenValue string
	lowestIdx    int
	highestIdx   int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var writtenNumbers []number
		for k, v := range numberMap {
			if strings.Contains(line, k) {
				firstIndex := strings.Index(line, k)
				lastIndex := strings.LastIndex(line, k)
				writtenNumbers = append(writtenNumbers, number{
					value:        v,
					writtenValue: k,
					lowestIdx:    firstIndex,
					highestIdx:   lastIndex,
				})
			}
			if strings.Contains(line, v) {
				firstIndex := strings.Index(line, v)
				lastIndex := strings.LastIndex(line, v)
				writtenNumbers = append(writtenNumbers, number{
					value:        v,
					writtenValue: k,
					lowestIdx:    firstIndex,
					highestIdx:   lastIndex,
				})
			}
		}

		var firstDigit, lastDigit string
		currentLowestIdx := len(line)
		currentHighestIdx := 0
		for _, num := range writtenNumbers {
			if len(writtenNumbers) == 1 {
				firstDigit = writtenNumbers[0].value
				lastDigit = writtenNumbers[0].value
				break
			}
			if num.highestIdx > currentHighestIdx {
				currentHighestIdx = num.highestIdx
				lastDigit = num.value
			}
			if num.lowestIdx < currentLowestIdx {
				currentLowestIdx = num.lowestIdx
				firstDigit = num.value
			}
		}

		finalDigit, err := strconv.Atoi(fmt.Sprintf("%s%s", firstDigit, lastDigit))
		if err == nil {
			sum += finalDigit
		}
	}
	// total
	fmt.Print(sum)
}

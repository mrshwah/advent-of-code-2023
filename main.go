package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var stringToDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func convertLineToDigits(line string) int {
	var first int
	var last int
	for i, char := range line {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			num = findDigitInString(line, i)
		}
		if num == 0 {
			continue
		}
		if first == 0 {
			first = num
		}
		last = num
	}
	return first*10 + last
}

func sumArray(digits []int) int {
	sum := 0

	for _, num := range digits {
		sum += num
	}
	return sum
}

func findDigitInString(line string, startIndex int) int {
	if startIndex+3 > len(line) {
		return 0
	}

	substring := line[startIndex : startIndex+3]
	digit, ok := stringToDigits[substring]

	if ok {
		return digit
	}

	if startIndex+4 > len(line) {
		return 0
	}

	substring = line[startIndex : startIndex+4]
	digit, ok = stringToDigits[substring]

	if ok {
		return digit
	}

	if startIndex+5 > len(line) {
		return 0
	}

	substring = line[startIndex : startIndex+5]
	digit, ok = stringToDigits[substring]

	if ok {
		return digit
	}

	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var digits []int
	for scanner.Scan() {
		line := scanner.Text()
		digits = append(digits, convertLineToDigits(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum of digits: %d\n", sumArray(digits))
}

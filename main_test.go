package main

import "testing"

func TestConvertLineToDigits(t *testing.T) {
	line := "1asonek19"
	expected := 19
	result := convertLineToDigits(line)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = "three2"
	expected = 32
	result = convertLineToDigits(line)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = "1212896"
	expected = 16
	result = convertLineToDigits(line)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = "fourefh5199ajkfj9eight"
	expected = 48
	result = convertLineToDigits(line)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = "11"
	expected = 11
	result = convertLineToDigits(line)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSumArray(t *testing.T) {
	digits := []int{1, 2, 3, 4, 5}
	expected := 15
	result := sumArray(digits)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	digits = []int{10, 20, 30, 40, 50}
	expected = 150
	result = sumArray(digits)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	digits = []int{0, 0, 0, 0, 0}
	expected = 0
	result = sumArray(digits)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	digits = []int{-1, -2, -3, -4, -5}
	expected = -15
	result = sumArray(digits)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	digits = []int{1}
	expected = 1
	result = sumArray(digits)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestFindDigitInString(t *testing.T) {
	line := "one"
	startIndex := 0
	expected := 1
	result := findDigitInString(line, startIndex)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = "two"
	startIndex = 0
	expected = 2
	result = findDigitInString(line, startIndex)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = "three"
	startIndex = 0
	expected = 3
	result = findDigitInString(line, startIndex)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = "four"
	startIndex = 0
	expected = 4
	result = findDigitInString(line, startIndex)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = "five"
	startIndex = 0
	expected = 5
	result = findDigitInString(line, startIndex)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = "six"
	startIndex = 0
	expected = 6
	result = findDigitInString(line, startIndex)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = "seven"
	startIndex = 0
	expected = 7
	result = findDigitInString(line, startIndex)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = "eight"
	startIndex = 0
	expected = 8
	result = findDigitInString(line, startIndex)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = "nine"
	startIndex = 0
	expected = 9
	result = findDigitInString(line, startIndex)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = "10twooo"
	startIndex = 2
	expected = 2
	result = findDigitInString(line, startIndex)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = "ajwefjnwk"
	startIndex = 3
	expected = 0
	result = findDigitInString(line, startIndex)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = "ab"
	startIndex = 1
	expected = 0
	result = findDigitInString(line, startIndex)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	line = ""
	startIndex = 0
	expected = 0
	result = findDigitInString(line, startIndex)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

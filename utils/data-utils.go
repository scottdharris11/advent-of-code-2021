package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ReadLines read lines from data file as strings
func ReadLines(dir string, filename string) []string {
	file, err := os.Open(buildPath(dir, filename))
	if err != nil {
		log.Fatalln(err)
	}
	defer closeFile(file)

	var values []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}
	return values
}

// ReadIntegers read lines from data file as integers
func ReadIntegers(dir string, filename string) []int {
	file, err := os.Open(buildPath(dir, filename))
	if err != nil {
		log.Fatalln(err)
	}
	defer closeFile(file)

	var values []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		values = append(values, val)
	}
	return values
}

// ReadIntegersFromLine reads single line from file and parses integers from csv format
func ReadIntegersFromLine(dir string, filename string) []int {
	lines := ReadLines(dir, filename)
	sValues := strings.Split(lines[0], ",")
	var values []int
	for _, sValue := range sValues {
		value, err := strconv.Atoi(sValue)
		if err != nil {
			log.Fatalln("non-numeric value detected", sValue, err)
		}
		values = append(values, value)
	}
	return values
}

// ReadIntegerGrid reads grid of integers from set of lines
func ReadIntegerGrid(lines []string) [][]int {
	var grid [][]int
	for _, line := range lines {
		var gridRow []int
		for _, r := range line {
			gridRow = append(gridRow, int(r-'0'))
		}
		grid = append(grid, gridRow)
	}
	return grid
}

// build platform independent file path
func buildPath(dir string, filename string) string {
	return filepath.Join(".", dir, filename)
}

// close file and log any errors
func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Println("Error closing file: ", err)
	}
}

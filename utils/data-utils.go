package utils

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ReadLines read lines from data file as strings
func ReadLines(dir string, filename string) []string {
	file := openFile(dir, filename)
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
	file := openFile(dir, filename)
	defer closeFile(file)

	var values []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values = append(values, Number(scanner.Text()))
	}
	return values
}

// ReadIntegersFromLine reads single line from file and parses integers from csv format
func ReadIntegersFromLine(dir string, filename string) []int {
	lines := ReadLines(dir, filename)
	sValues := strings.Split(lines[0], ",")
	var values []int
	for _, sValue := range sValues {
		values = append(values, Number(sValue))
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

// Number parses the supplied string into number
func Number(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln("bad coordinate received: ", s)
	}
	return n
}

// FilePath builds path to file with supplied directory if exists, otherwise current directory
func FilePath(dir string, filename string) string {
	_, err := os.Stat(dir)
	if errors.Is(err, os.ErrNotExist) {
		return buildPath(".", filename)
	}
	return buildPath(dir, filename)
}

// build platform independent file path
func buildPath(dir string, filename string) string {
	return filepath.Join(".", dir, filename)
}

// open filename supplied.  look in current directly if not found in supplied directory
func openFile(dir string, filename string) *os.File {
	file, err := os.Open(buildPath(dir, filename))
	if err != nil {
		file, err = os.Open(buildPath(".", filename))
		if err != nil {
			log.Fatalln(err)
		}
	}
	return file
}

// close file and log any errors
func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Println("Error closing file: ", err)
	}
}

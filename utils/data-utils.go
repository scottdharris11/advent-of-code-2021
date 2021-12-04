package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// ReadLines read lines from data file as strings
func ReadLines(dir string, filename string) []string {
	file, err := os.Open(buildPath(dir, filename))
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

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
	defer file.Close()

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

// build platform independent file path
func buildPath(dir string, filename string) string {
	return filepath.Join(".", dir, filename)
}

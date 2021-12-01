package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// read lines from data file as integers
func readIntegers(filename string) []int {
	file, err := os.Open(buildPath(filename))
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
func buildPath(filename string) string {
	return filepath.Join(".", "data", filename)
}

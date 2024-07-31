package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]uint)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, fileMap := range counts {
		count := getMapSum(fileMap)
		if count > 1 {
			fmt.Printf("%d\t%s\nFound in files: ", count, line)
			for filename := range fileMap {
				fmt.Printf("%s ", filename)
			}
			fmt.Print("\n")
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]uint) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = make(map[string]uint)
		}
		counts[line][f.Name()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func getMapSum(fileMap map[string]uint) uint {
	var sum uint
	for _, v := range fileMap {
		sum += v
	}
	return sum
}

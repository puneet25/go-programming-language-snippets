package main

/*
 Modify dup2 to print names of all files in which each duplicated line occurs
*/
import (
	"bufio"
	"fmt"
	"os"
)

func ReadingFromNamedFiles() {
	counts := make(map[string]int)
	fileNames := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				// %v - prints any value
				fmt.Printf("error opening named files: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("count is %d\tfor line:%s\tin fileNames:%v\n", n, line, fileNames[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string][]string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
		fileNames[input.Text()] = appendAfterCheckingDuplicate(fileNames[input.Text()], f.Name())
	}
}

func appendAfterCheckingDuplicate(fileNames []string, fileName string) []string {
	for _, name := range fileNames {
		if name == fileName {
			return fileNames
		}
	}
	fileNames = append(fileNames, fileName)
	return fileNames
}

func main() {
	ReadingFromNamedFiles()
}

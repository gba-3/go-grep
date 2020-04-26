package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func hasError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func scanFile(searchStr string, filePath string) {
	file, err := os.Open(filePath)
	hasError(err)
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		lineStr := s.Text()
		if strings.Contains(lineStr, searchStr) {
			fmt.Printf("%s:%s\n", filePath, lineStr)
		}
	}
}

func main() {
	flag.Parse()
	searchStr := flag.Arg(0)
	filePath := flag.Arg(1)

	scanFile(searchStr, filePath)
}

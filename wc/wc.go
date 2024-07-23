package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func count(fileName string) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fileSize := fileInfo.Size()
	fmt.Printf("The file is %d bytes long\n", fileSize)
}

func lines(fileName string) {
	lineCount := 0
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The file is %d lines long\n", lineCount)
}

func words(fileName string) {
	wordCount := 0
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		for _, w := range words {
			if w != "" {
				wordCount++
			}
		}
		//fmt.Println(words)
		// if len(words) == 1 && words[0] == "" {
		// 	continue
		// } else if words[0] == "\n" {
		// 	continue
		// }
		// wordCount += len(words)
		fmt.Println(wordCount)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The file is %d words long\n", wordCount)
}

func main() {

	//need to add check to make sure 2 args set
	action := os.Args[1]
	fileName := os.Args[2]

	switch action {
	case "-c":
		fileName := os.Args[2]
		count(fileName)
	case "-l":
		lines(fileName)
	case "-w":
		words(fileName)
	default:
		fmt.Println("No function matching", action)
	}
}

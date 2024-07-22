package main

import (
	"fmt"
	"os"
)

func count(fileName string) {
	//file, err := os.Open(fileName)
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
	}
	//defer file.Close()
	fileSize := fileInfo.Size()
	fmt.Printf("The file is %d bytes long\n", fileSize)
}

func main() {

	action := os.Args[1]

	switch action {
	case "-c":
		fileName := os.Args[2]
		count(fileName)
	default:
		fmt.Println("No function matching", action)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	readFile, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var result int
	for fileScanner.Scan() {
		switch fileScanner.Text() {
		case "A X":
			result += 3
		case "A Y":
			result += 4
		case "A Z":
			result += 8
		case "B X":
			result += 1
		case "B Y":
			result += 5
		case "B Z":
			result += 9
		case "C X":
			result += 2
		case "C Y":
			result += 6
		case "C Z":
			result += 7
		}
	}

	fmt.Println(result)

}

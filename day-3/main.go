package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// split de cada linha na metade
// encontrar qual letra se repete
// a-z vai de 1 a 26
// A-Z vai de 27  a 52

var priorities = map[string]int{}

func ASCIIIntToChar(code int) string {
	return string(rune(code))
}

func main() {

	for i := 0; i < 26; i++ {
		priorities[ASCIIIntToChar('a'+i)] = i + 1
		priorities[ASCIIIntToChar('A'+i)] = i + 27
	}

	readFile, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := 0
	for fileScanner.Scan() {
		result += splitHands(fileScanner.Text())
	}

	fmt.Println(result)

}

func splitHands(text string) int {
	prioritiesSum := 0
	for _, l := range strings.Split(text, "\n") {
		rightHalf := l[len(l)/2:]
		for i := 0; i < len(l)/2; i++ {
			leftChar := l[i : i+1]
			if strings.Contains(rightHalf, leftChar) {
				prioritiesSum += priorities[leftChar]
				break
			}
		}
	}
	return prioritiesSum
}

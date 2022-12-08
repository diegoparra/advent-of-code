package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/// part 1
// split de cada linha na metade
// encontrar qual letra se repete
// a-z vai de 1 a 26
// A-Z vai de 27  a 52

/// part 2
// identificar cada grupo a cada tres linhas
// pegar o char que se repete nas 3 linhas de cada grupo

var priorities = map[string]int{}

func ASCIIIntToChar(code int) string {
	return string(rune(code))
}

func main() {

	for i := 0; i < 26; i++ {
		priorities[ASCIIIntToChar('a'+i)] = i + 1
		priorities[ASCIIIntToChar('A'+i)] = i + 27
	}

	readFile, err := os.Open("test-file.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	// result := 0
	var group []string
	// value := 0
	for fileScanner.Scan() {
		if len(group) == 3 {
			fmt.Println("exiting from the loop")
			// var word string
			var temp []string
			for i := 0; i < 3; i++ {
				temp = append(temp, group...)
				// fmt.Println(temp)
				for _, char := range temp {
					fmt.Println(char)
					// if word == char {
					// 	fmt.Println("found the char equal: ", char)
					// }
				}
			}
			group = nil

		} else {
			group = append(group, fileScanner.Text())
			// fmt.Println(group)
		}

	}
	// fmt.Println(result)

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

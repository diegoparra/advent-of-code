package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	readFile, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var groups [][]int
	var valueSlice []int

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			groups = append(groups, valueSlice)
			valueSlice = nil
		} else {
			value, _ := strconv.Atoi(fileScanner.Text())
			valueSlice = append(valueSlice, value)
		}
	}

	var result []int
	for i := 0; i < len(groups); i++ {
		count := 0
		for _, v := range groups[i] {
			count = count + v
		}
		result = append(result, count)
	}

	var largerNumber, temp int

	for _, element := range result {
		if element > temp {
			temp = element
			largerNumber = temp
		}
	}
	fmt.Println("Largest number of Array/slice is  ", largerNumber)
}

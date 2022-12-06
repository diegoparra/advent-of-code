package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func sumElfs(filename string) []int {

	readFile, err := os.Open(filename)
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
	return result
}

func getBiggest(values []int) int {

	var largerNumber, temp int

	for _, element := range values {
		if element > temp {
			temp = element
			largerNumber = temp
		}
	}
	return largerNumber
}

func getThreeBiggest(values []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	var result int

	for i := 0; i < 3; i++ {
		result += values[i]
	}

	return result

}

func main() {

	elfsConsumed := sumElfs("file.txt")

	result := getBiggest(elfsConsumed)

	resultThreeElfs := getThreeBiggest(elfsConsumed)

	fmt.Println("Elf carring the biggest value: ", result)

	fmt.Println("Total 3 Elf's: ", resultThreeElfs)
}

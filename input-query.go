package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countOccurrences(input []string, query []string) []int {

	wordCount := make(map[string]int)
	for _, word := range input {
		wordCount[word]++
	}

	result := make([]int, len(query))
	for i, q := range query {
		result[i] = wordCount[q]
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Masukkan elemen input:")
	strInput, _ := reader.ReadString('\n')
	strInput = strings.TrimSpace(strInput)
	strArr := strings.Split(strInput, " ")

	fmt.Println("Masukkan elemen query:")
	strQuery, _ := reader.ReadString('\n')
	strQuery = strings.TrimSpace(strQuery)
	queryArr := strings.Split(strQuery, " ")

	output := countOccurrences(strArr, queryArr)
	fmt.Println(output)
}

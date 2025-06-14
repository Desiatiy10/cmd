package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile, err := os.Open("in.txt")
	if err != nil {
		fmt.Println("Ошибка открытия in.txt - ", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create("out.txt")
	if err != nil {
		fmt.Println("Ошибка создания out.txt - ", err)
		return
	}
	defer outputFile.Close()

	scaner := bufio.NewScanner(inputFile)
	lineNumber := 1

	for scaner.Scan() {
		line := scaner.Text()

		numberedLine := fmt.Sprintf("%d. %s\n", lineNumber, line)

		_, err = outputFile.WriteString(numberedLine)
		if err != nil {
			fmt.Println("Ошбика записи - ", err)
			return
		}
		lineNumber++
	}
}

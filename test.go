package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	v, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	outputFile, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("there is an error creating an output file :", err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(v)
	writer := bufio.NewWriter(outputFile)
	for scanner.Scan() {
		line := scanner.Text()
		modifiedLine := modifiedLine(line)

		_, err := writer.WriteString(modifiedLine + "\n")
		if err != nil {
			fmt.Println("error writing the output:", err)
			os.Exit(1)
		}

	}
	writer.Flush()
}

func modifiedLine(line string) string {
	return line
}

func hex(hexString string) int64 {
	decimalInt, err := strconv.ParseInt(hexString, 16, 32)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return decimalInt
}

func binnary(binString string) int64 {
	binnarylInt, err := strconv.ParseInt(binString, 2, 32)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return binnarylInt
}

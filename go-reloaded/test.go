package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() { // better os readFile and writeFile \\
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . sample.txt output.txt")
		return
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	text, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	modifiedText := textChanges(string(text))

	err = os.WriteFile(outputFileName, []byte(modifiedText), 0420)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("check the result file:", outputFileName)
}

// func to applay changes
func textChanges(line string) string {
	line = hexToDec(line)
	line = binToDec(line)
	line = up(line)
	line = low(line)
	line = cap(line)
	line = rePunctuation(line)
	line = An(line)
	line = rePunctuation2(line)
	line = transformWords(line)
	line = spaceafterQ(line)
	return line
}

// a3ml kol func l7al odefho 3ala (func textChanges)
func hexToDec(line string) string {
	//////////////////////////////////hex
	reHex := regexp.MustCompile(`([0-9A-Fa-f]+) \(hex\)`)
	return reHex.ReplaceAllStringFunc(line, func(match string) string {
		hexValue := reHex.FindStringSubmatch(match)[1]
		decimalValue, _ := strconv.ParseInt(hexValue, 16, 64)

		return strconv.FormatInt(decimalValue, 10)
	})
}

// //////////////////////////////bin
func binToDec(line string) string {
	reBin := regexp.MustCompile(`([01]+) \(bin\)`)
	return reBin.ReplaceAllStringFunc(line, func(match string) string {
		binValue := reBin.FindStringSubmatch(match)[1]
		decimalValue, _ := strconv.ParseInt(binValue, 2, 64)

		return strconv.FormatInt(decimalValue, 10)
	})
}

func up(line string) string {
	reUp := regexp.MustCompile(`(\w+)\s*\(up\)`) //we used s* insted of +
	return reUp.ReplaceAllStringFunc(line, func(match string) string {
		up := reUp.FindStringSubmatch(match)[1]
		match = strings.ToUpper(up)
		return match
	})
}

// /////////////////////////low
func low(line string) string {
	reLow := regexp.MustCompile(`(\w+)\s+\(low\)`) //`([^\s]+) \(low\)`
	return reLow.ReplaceAllStringFunc(line, func(match string) string {
		low := reLow.FindStringSubmatch(match)[1]
		match = strings.ToLower(low)
		return match
	})

}

// //////////////////////////cap
func cap(line string) string {
	reCap := regexp.MustCompile(`(\w+)\s+\(cap\)`) //([^\s]+) \(cap\)
	return reCap.ReplaceAllStringFunc(line, func(match string) string {
		cap := reCap.FindStringSubmatch(match)[1]
		match = strings.Title(cap)
		return match
	})
}

// /////////////////////////////remove space befor !!
func rePunctuation(line string) string {
	rePunctuation := regexp.MustCompile(`(\w+)\s*([.,!?;:]+)`) // better `(\w) ?([.,!?:;]+) ?(\w)
	result := rePunctuation.ReplaceAllStringFunc(line, func(match string) string {
		return strings.Join(strings.Fields(match), "")
		//return formatPunctuation(match)
	})
	result = regexp.MustCompile(`([.]{3}|[!?]+)`).ReplaceAllString(result, "${1}")

	return result
}
func rePunctuation2(input string) string {
	pattern := regexp.MustCompile(`,(\w)`)
	result := pattern.ReplaceAllString(input, ", $1")
	return result
}

func An(line string) string {
	reLowerA := regexp.MustCompile(`\b([a])\b`)
	line = reLowerA.ReplaceAllString(line, "an")

	reUpperA := regexp.MustCompile(`\b([A])\b`)
	line = reUpperA.ReplaceAllString(line, "An")

	return line
}
func spaceafterQ(line string) string {
	line = strings.ReplaceAll(line, " '", "'")
	line = strings.ReplaceAll(line, "' ", "'")
	return line
}

// //                       up 2 down 2 .....
func transformWords(line string) string {
	lineFilds := strings.Fields(line)
	for index, stringsFromline := range lineFilds {
		if stringsFromline == "(up," {
			submatch2, err := strconv.Atoi(string(lineFilds[index+1][0]))
			if err != nil {
				fmt.Println(err)
			}
			for modIndex := submatch2; modIndex >= 1; modIndex-- {
				WodrsIwantToModify := strings.ToUpper(lineFilds[index-modIndex])
				lineFilds[index-modIndex] = WodrsIwantToModify
			}
			lineFilds[index] = ""
			lineFilds[index+1] = ""
		}
		if stringsFromline == "(low," {
			submatch2, err := strconv.Atoi(string(lineFilds[index+1][0]))
			if err != nil {
				fmt.Println(err)
			}
			for modIndex := submatch2; modIndex >= 1; modIndex-- {
				WodrsIwantToModify := strings.ToLower(lineFilds[index-modIndex])
				lineFilds[index-modIndex] = WodrsIwantToModify
			}
			lineFilds[index] = ""
			lineFilds[index+1] = ""
		}
		if stringsFromline == "(cap," {
			submatch2, err := strconv.Atoi(string(lineFilds[index+1][0]))
			if err != nil {
				fmt.Println(err)
			}
			for modIndex := submatch2; modIndex >= 1; modIndex-- {
				WodrsIwantToModify := strings.Title(lineFilds[index-modIndex])
				lineFilds[index-modIndex] = WodrsIwantToModify
			}
			lineFilds[index] = ""
			lineFilds[index+1] = ""
		}
	}
	joinString := strings.Join(lineFilds, " ")
	result := strings.ReplaceAll(joinString, "  ", " ")
	return result
}

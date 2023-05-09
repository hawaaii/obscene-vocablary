package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filename, userInput, result string
	fmt.Scan(&filename)
	file := openFile(filename)
	lines := readFileLines(file)
	file.Close()

	for {
		fmt.Scanf("%s", &userInput)
    inputSlice := strings.Split(userInput," ")
		if "exit" == userInput {
			fmt.Println("Bye!")
			break
		} else {
      var resultStr string
      for _, wrd := range inputSlice {
        result = checkWords(lines,wrd)
        resultStr = resultStr + result
      }
      fmt.Println(resultStr)
		}
	}
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	return file
}

func readFileLines(file *os.File) []string {
	var lines []string
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func checkWords(lines []string,userInput string) string {
  var result string
	if contains(lines, strings.ToLower(userInput)) {
		for i := 0; i < len(userInput); i++ {
			result += "*"
		}
	} else {
		result = userInput
	}
  return result
}

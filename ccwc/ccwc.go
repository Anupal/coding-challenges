package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Command line args")
	cmdlineArgs := os.Args
	for index, arg := range cmdlineArgs {
		fmt.Println(index, arg)
	}

	flagBytes, flagChars, flagLines, flatWords, fileList := captureFlags(cmdlineArgs)

	// If no flags are passed
	if !(flagBytes && flagChars && flagLines && flatWords) {
		flagBytes = true
		flagLines = true
		flatWords = true
	}

	// loop over fileName list
	for _, fileName := range fileList {
		filePtr, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("File '%s' not found!", fileName)
		}
		fileData := make([]byte, 100)
		byteCount, err := filePtr.Read(fileData)
		if err != nil {
			log.Fatalf("File '%s' not readable", fileName)
		}
		fileData = bytes.TrimRight(fileData, "\x00")

		fileDataString := string(fileData)
		charCount := len(fileDataString)

		lines := strings.Split(fileDataString, "\n")
		lineCount := len(lines) - 1

		wordCount := 0
		for _, line := range lines {
			wordCount += len(strings.Fields(line))
		}

		byteCountLen := len(strconv.Itoa(byteCount))
		lineCountLen := len(strconv.Itoa(lineCount))
		wordCountLen := len(strconv.Itoa(wordCount))
		charCountLen := len(strconv.Itoa(charCount))
		maximumLen := max(byteCountLen, lineCountLen, wordCountLen, charCountLen)

		fmt.Printf("%*d ", maximumLen, lineCount)
		fmt.Printf("%*d ", maximumLen, wordCount)
		fmt.Printf("%*d ", maximumLen, byteCount)

		//fmt.Printf("%*d ", maximumLen, charCount)
		fmt.Println()

		//fmt.Printf("Byte count for file '%s': %d\n", fileName, byteCount)
		//fmt.Printf("Line count for file '%s': %d\n", fileName, lineCount)
		//fmt.Printf("Word count for file '%s': %d\n", fileName, wordCount)
		//fmt.Printf("Char count for file '%s': %d\n", fileName, charCount)
	}

}

func captureFlags(args []string) (bytes bool, chars bool, lines bool, words bool, fileList []string) {
	fileList = make([]string, 0, 10)
	for _, arg := range args[1:] {
		switch arg {
		case "-c", "--bytes":
			bytes = true
		case "-m", "--chars":
			chars = true
		case "-l", "--lines":
			lines = true
		case "-w", "--words":
			words = true
		default:
			fileList = append(fileList, arg)
		}
	}
	return
}

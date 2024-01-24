package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type fileData struct {
	name  string
	bytes int
	chars int
	lines int
	words int
}

func maxWidth(flags map[string]bool, fileDataList []fileData) int {
	maximum := 0

	for _, file := range fileDataList {
		includedValues := make([]int, 0, 4)
		if flags["lines"] {
			includedValues = append(includedValues, len(strconv.Itoa(file.lines)))
		}
		if flags["words"] {
			includedValues = append(includedValues, len(strconv.Itoa(file.words)))
		}
		if flags["bytes"] {
			includedValues = append(includedValues, len(strconv.Itoa(file.bytes)))
		}
		if flags["chars"] {
			includedValues = append(includedValues, len(strconv.Itoa(file.chars)))
		}
		for _, value := range includedValues {
			if value > maximum {
				maximum = value
			}
		}
	}

	return maximum
}

func (file *fileData) parseFile(filePath string) {
	_, file.name = path.Split(filePath)

	filePtr, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("File '%s' not found!", file.name)
	}

	fileData := make([]byte, 100)
	file.bytes, err = filePtr.Read(fileData)
	if err != nil {
		log.Fatalf("File '%s' not readable", file.name)
	}
	fileData = bytes.TrimRight(fileData, "\x00")

	fileDataString := string(fileData)
	file.chars = len(fileDataString)

	lines := strings.Split(fileDataString, "\n")
	file.lines = len(lines) - 1

	file.words = 0
	for _, line := range lines {
		file.words += len(strings.Fields(line))
	}
}

func (file *fileData) displayRow(flags map[string]bool, maximumWidth int) {
	if flags["lines"] {
		fmt.Printf("%*d ", maximumWidth, file.lines)
	}
	if flags["words"] {
		fmt.Printf("%*d ", maximumWidth, file.words)
	}
	if flags["bytes"] {
		fmt.Printf("%*d ", maximumWidth, file.bytes)
	}
	if flags["chars"] {
		fmt.Printf("%*d ", maximumWidth, file.chars)
	}
	fmt.Printf("%s\n", file.name)
}

func main() {
	cmdlineArgs := os.Args

	flags, fileList := captureFlags(cmdlineArgs)

	// If no flags are passed
	ifFlagsPassed := false
	for _, flag := range flags {
		ifFlagsPassed = ifFlagsPassed || flag
	}

	if !ifFlagsPassed {
		flags["bytes"] = true
		flags["lines"] = true
		flags["words"] = true
	}

	fileDataList := make([]fileData, len(fileList))
	// loop over file list
	for i, filePath := range fileList {
		fileDataList[i].parseFile(filePath)
	}

	maximumWidth := maxWidth(flags, fileDataList)

	for _, file := range fileDataList {
		file.displayRow(flags, maximumWidth)
	}
}

// bytes bool, chars bool, lines bool, words bool
func captureFlags(args []string) (map[string]bool, []string) {
	flags := make(map[string]bool)
	fileList := make([]string, 0, 10)

	flags["bytes"] = false
	flags["chars"] = false
	flags["lines"] = false
	flags["words"] = false

	for _, arg := range args[1:] {
		switch arg {
		case "-c", "--bytes":
			flags["bytes"] = true
		case "-m", "--chars":
			flags["chars"] = true
		case "-l", "--lines":
			flags["lines"] = true
		case "-w", "--words":
			flags["words"] = true
		default:
			fileList = append(fileList, arg)
		}
	}
	return flags, fileList
}

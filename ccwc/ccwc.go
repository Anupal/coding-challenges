package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type data struct {
	bytes int
	chars int
	lines int
	words int
	name  string
}

func maxWidth(flags map[string]bool, fileDataList []data) int {
	maximum := 0

	for _, file := range fileDataList {
		fieldLengths := map[string]int{
			"lines": len(strconv.Itoa(file.lines)),
			"words": len(strconv.Itoa(file.words)),
			"bytes": len(strconv.Itoa(file.bytes)),
			"chars": len(strconv.Itoa(file.chars)),
		}

		for flag, length := range fieldLengths {
			if flags[flag] && length > maximum {
				maximum = length
			}
		}
	}

	return maximum
}

func (d *data) parseStdio() {
	d.name = ""

	stdioData := make([]byte, 0, 100)

	stdinScanner := bufio.NewScanner(os.Stdin)
	for stdinScanner.Scan() {
		stdioData = append(stdioData, stdinScanner.Bytes()...)
	}

	d.bytes = len(stdioData)

	fileDataString := string(stdioData)
	d.chars = len(fileDataString)

	lines := strings.Split(fileDataString, "\n")
	d.lines = len(lines) - 1

	d.words = 0
	for _, line := range lines {
		d.words += len(strings.Fields(line))
	}
}

func (d *data) parseFile(filePath string) {
	_, d.name = path.Split(filePath)

	filePtr, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("File '%s' not found!", d.name)
	}

	fileData := make([]byte, 100)
	d.bytes, err = filePtr.Read(fileData)
	if err != nil {
		log.Fatalf("File '%s' not readable", d.name)
	}
	fileData = bytes.TrimRight(fileData, "\x00")

	fileDataString := string(fileData)
	d.chars = len(fileDataString)

	lines := strings.Split(fileDataString, "\n")
	d.lines = len(lines) - 1

	d.words = 0
	for _, line := range lines {
		d.words += len(strings.Fields(line))
	}
}

func (d *data) displayRow(flags map[string]bool, maximumWidth int) {
	if flags["lines"] {
		fmt.Printf("%*d ", maximumWidth, d.lines)
	}
	if flags["words"] {
		fmt.Printf("%*d ", maximumWidth, d.words)
	}
	if flags["bytes"] {
		fmt.Printf("%*d ", maximumWidth, d.bytes)
	}
	if flags["chars"] {
		fmt.Printf("%*d ", maximumWidth, d.chars)
	}
	fmt.Printf("%s\n", d.name)
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

	if len(fileList) == 0 {
		stdio := data{}
		stdio.parseStdio()
		maximumWidth := maxWidth(flags, []data{stdio})

		stdio.displayRow(flags, maximumWidth)
	} else {

		fileDataList := make([]data, len(fileList))
		// loop over file list
		for i, filePath := range fileList {
			fileDataList[i].parseFile(filePath)
		}

		maximumWidth := maxWidth(flags, fileDataList)

		for _, file := range fileDataList {
			file.displayRow(flags, maximumWidth)
		}
	}
}

// bytes bool, chars bool, lines bool, words bool
func captureFlags(args []string) (map[string]bool, []string) {
	flags := map[string]bool{
		"bytes": false,
		"chars": false,
		"lines": false,
		"words": false,
	}
	fileList := make([]string, 0, 10)

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

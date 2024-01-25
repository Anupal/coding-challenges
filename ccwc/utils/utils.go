package utils

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

type Data struct {
	bytes int
	chars int
	lines int
	words int
	name  string
}

// MaxWidth gets the maximum width among all data values across all rows of output.
func MaxWidth(flags map[string]bool, dataList []Data) int {
	maximum := 0

	for _, data := range dataList {
		fieldLengths := map[string]int{
			"lines": len(strconv.Itoa(data.lines)),
			"words": len(strconv.Itoa(data.words)),
			"bytes": len(strconv.Itoa(data.bytes)),
			"chars": len(strconv.Itoa(data.chars)),
		}

		for flag, length := range fieldLengths {
			if flags[flag] && length > maximum {
				maximum = length
			}
		}
	}

	return maximum
}

// ParseStdio parses trough Stdio input and assigns values to the calling Data structure
func (d *Data) ParseStdio() {
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

// ParseFile parses file by filePath and assigns values to the calling Data structure
func (d *Data) ParseFile(filePath string) {
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

// DisplayRow prints data values based on passed cmdline flags
func (d *Data) DisplayRow(flags map[string]bool, maximumWidth int) {
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

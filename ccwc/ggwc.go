package main

import (
	"fmt"
	"log"
	"os"
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
		flagLines = true
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
		fmt.Printf("Byte count for file '%s': %d\n", fileName, byteCount)
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

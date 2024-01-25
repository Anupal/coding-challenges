package main

import (
	"github.com/anupal/coding-challenges/ccwc/utils"
	"os"
)

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
		stdio := utils.Data{}
		stdio.ParseStdio()
		maximumWidth := utils.MaxWidth(flags, []utils.Data{stdio})

		stdio.DisplayRow(flags, maximumWidth)
	} else {

		fileDataList := make([]utils.Data, len(fileList))
		// loop over file list
		for i, filePath := range fileList {
			fileDataList[i].ParseFile(filePath)
		}

		maximumWidth := utils.MaxWidth(flags, fileDataList)

		for _, file := range fileDataList {
			file.DisplayRow(flags, maximumWidth)
		}
	}
}

// captureFlags parses cmdline flags and file list
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

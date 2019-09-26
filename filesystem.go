package gart

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

// InitFS is given an output path and does any fs specific initializations
// (currently just creating the output folder).
func InitFS(outputPath string) {
	// TODO: Better permissioning?
	os.MkdirAll(outputPath, 0777)
}

// GetOutputFileName will look at a given path with a given name string and
// looks for the next available file name. Files are assumed to be named like
// this: "<outName>_<id>.png". The ID is determined by the files in the folder
// and will automatically be set to the max ID found + 1.
func GetOutputFileName(outputPath string, outName string) (string, error) {
	files, err := ioutil.ReadDir(outputPath)
	if err != nil {
		return "", err
	}

	// Get the max number value from the list of files
	maxidx := -1
	re, err := regexp.Compile(outName + "_(\\d+)\\.\\w+$")
	if err != nil {
		return "", err
	}

	for _, file := range files {
		filename := file.Name()
		if match := re.FindStringSubmatch(filename); match != nil {
			idx, _ := strconv.Atoi(match[1])
			if idx > maxidx {
				maxidx = idx
			}
		}
	}

	outname := fmt.Sprintf("%s_%04d.png", outName, maxidx+1)

	return outname, nil
}

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	process "spaceconvention/internal"
)

func main() {
	DirectoryName := os.Args[1]
	ExceptionFiles := os.Args[2:]
	var arrayOfPath []string

	err := filepath.Walk(DirectoryName,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			arrayOfPath = append(arrayOfPath, path)
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	for _, filePath := range arrayOfPath {
		ok := func() bool {
			for _, exceptionFile := range ExceptionFiles {
				if exceptionFile == filePath {
					return false
				}
			}
			return true
		}
		if ok() {
			isDir := isDirectory(filePath)
			if !isDir {
				err := process.FormateFile(filePath)
				if err != nil {
					errors.New(fmt.Sprintln(err))
				}
			}
		}

	}

}

func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

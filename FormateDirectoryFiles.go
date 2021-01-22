package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/02amanag/spaceconvention/spaceconvention"
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

	fmt.Println(arrayOfPath)

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
			err := spaceconvention.FormateFile(filePath)
			if err != nil {
				errors.New(fmt.Sprintln(err) + filePath)
			}
		}

	}

}

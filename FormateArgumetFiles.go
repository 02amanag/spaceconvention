package main

import (
	"errors"
	"fmt"
	"os"

	process "spaceconvention/internal"
)

func main() {

	arrayOfPath := os.Args[1:]

	for _, filePath := range arrayOfPath {
		err := process.FormateFile(filePath)
		if err != nil {
			errors.New(fmt.Sprintln(err) + filePath)
		}
	}

}

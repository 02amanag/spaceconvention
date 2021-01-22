package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/02amanag/spaceconvention/spaceconvention"
)

func main() {

	arrayOfPath := os.Args[1:]

	for _, filePath := range arrayOfPath {
		err := spaceconvention.FormateFile(filePath)
		if isError(err) {
			errors.New(fmt.Sprintln(err) + filePath)
		}
	}

}

package spaceconvention

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func FormateFile(path string) error {
	text, _ := readFile(path)
	formatedText := spaceconvention(text)

	ok := deleteFile(path)
	if isError(ok) {
		errors.New("error in deletion ")
	}
	err := createFile(path)
	if isError(err) {
		errors.New("error in create file ")
	}
	writeFile(path, formatedText)
	if isError(err) {
		errors.New("error in writing file ")
	}

	return nil
}

func spaceconvention(text string) string {
	var formatedText string
	for indexNo := range text {
		if text[indexNo] == 32 { //on SPACE
			if indexNo == 0 { //find SPACE at start
			} else if indexNo == len(text)-1 { //space at the end of the file
			} else if text[indexNo+1] == 10 { //find SPACE before ENTER
			} else if indexNo > 0 && text[indexNo-1] == 10 { // find Space after ENTER
			} else {
				formatedText += string(text[indexNo])
			}
		} else if text[indexNo] == 10 { //on ENTER
			if indexNo == len(text)-1 { //ENTER at the end of the file
			} else if indexNo+2 <= len(text)-1 {
				if text[indexNo+1] == 10 && text[indexNo+2] == 10 { //two ENTER continious
				} else {
					formatedText += string(text[indexNo])
				}
			} else {
				formatedText += string(text[indexNo])
			}
		} else {
			formatedText += string(text[indexNo])
		}
	}
	newLenght := len(formatedText)
	if formatedText[newLenght-1] == 32 || formatedText[newLenght-1] == 10 {
		formatedText = spaceconvention(formatedText)
	}
	return formatedText

	// TODO : continue space and then Enter -> remove space
}

func createFile(path string) error {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return err
		}
		defer file.Close()
	}

	return nil
}

func writeFile(path string, text string) error {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return err
	}
	defer file.Close()

	// Write text file.
	_, err = file.WriteString(text)
	if isError(err) {
		return err
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return err
	}
	return nil
}

func readFile(path string) (string, error) {
	// Open file for reading.
	byetArray, err := ioutil.ReadFile(path)
	if isError(err) {
		log.Panicf("failed reading data from file: %s ", err)
	}

	return string(byetArray), nil
}

// delete file
func deleteFile(path string) error {
	var err = os.Remove(path)
	if isError(err) {
		return err
	}
	return nil
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

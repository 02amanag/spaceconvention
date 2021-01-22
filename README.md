# spaceconvention

This repository is build to format files.
there are useless spaces in file at start , in between or in end.
With this repository we can remove the extra spaces and formate the file into smaller size.

how to use -

#method 1
passing file name we want to format after the run comand...
 go run FormateArgumentFile.go <fileName1> <fileName2> ....
 
 #method 2
 passing the Directory name which formate all file in directory and Sub-directory, with name of file we dont want to format.
 where the first agrument id Directory name and after that all are consider as ignoring file names.
 go run FormateDirectoryFiles.go <directoryName> <toIgnorFile1> <toIgnorFile2> ...
 
 #method 3
 use in go file
 
 import (
  "log"
  "github.com/02amanag/spaceconvention/spaceconvention"
  )
  
  func main(){
    err := spaceconvention.FormateFile("filePath") //filePath is path of file we want to formate
      if err != nil {
       log.Fatal(err)
      }
   }

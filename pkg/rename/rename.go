package rename

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/victorbischoff/fileutilities/pkg/entities"
)

// GetPath gets the directory containing the files which the user whishes to rename
func GetPath() string {
	fmt.Println("Please enter path (use ./ for a folder in current directory): ")
	var path string
	fmt.Scan(&path)

	return path
}

// GetPrefix allows the user to enter a prefix which all the files will contain
func GetPrefix() string {
	fmt.Println("Please enter prefix: ")
	var prefix string
	fmt.Scan(&prefix)
	return prefix
}

// RenameFilesInDirectory loops through all the files in the directory and renames them starting with the index number + 1 followed by the useres prefix
func RenameFilesInDirectory(cleanPath, prefix string) {
	_, err := os.Stat(cleanPath)
	if os.IsNotExist(err) {
		log.Fatalf("directory %s does not exist", cleanPath)
	}

	if prefix == "" {
		log.Fatal("prefix cannot be empty")
	}
	userIn := &entities.InputData{
		Path:   cleanPath,
		Prefix: prefix,
	}
	fullPath := fmt.Sprintf("%v/*", cleanPath)

	files, err1 := filepath.Glob(fullPath)
	if err1 != nil {
		log.Fatal(err1)
	}

	for index, file := range files {
		fname := strings.Split(file, ".")
		s := fname[len(fname)-1]
		os.Rename(file, fmt.Sprintf("%v/%v-%v.%v", userIn.Path, index+1, userIn.Prefix, s))
	}
}

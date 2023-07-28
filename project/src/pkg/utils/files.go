package utils

import (
	"os"
	"path/filepath"
	"strings"
)

type Utils struct{}

func New() *Utils {
	return &Utils{}
}

func (u *Utils) GetRelativeRootDir() (*string, error) {
	fullRoot, err := u.GetRootDir()
	if err != nil {
		return nil, err
	}
	arr := strings.Split(*fullRoot, "/")

	src := arr[len(arr)-2]

	return &src, nil
}

// * Return the project root as a string
func (u *Utils) GetRootDir() (*string, error) {
	var result string = ""
	// * this path is from the caller file
	thisDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}

	fullDir := strings.Split(thisDir, "/")

	// * here is tha folder name above src dir
	// * The name "project" must match the go mod package name, and the most outer directory
	for _, x := range fullDir {
		if x != "project" {
			result += x + "/"
		} else {
			break
		}
	}

	result += "project/"

	return &result, nil
}

func (u *Utils) GetFilePath(pathFromRoot *[]string) (*string, error) {
	// * 2 is the level from current file
	root, err := u.GetRootDir()
	if err != nil {
		return nil, err
	}

	for i, x := range *pathFromRoot {
		if i < len(*pathFromRoot)-1 {
			*root += x + "/"
		} else {
			*root += x
		}
	}
	return root, nil
}

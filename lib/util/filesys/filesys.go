package filesys

import (
	"fmt"
	"os"
)

const (
	statFailedErr    = "stat path %s failed: %v"
	pathIsEmptyErr   = "given path is empty"
	pathNotExistsErr = "given path %s is not exist"
	notDirectoryErr  = "given path %s is not a directory"
	notFileErr       = "given path %s is not a file"
)

// IsFileExists checks a file and return true if it is exists, and it is a file
func IsFileExists(file string) (bool, error) {
	if len(file) == 0 {
		return false, fmt.Errorf(pathIsEmptyErr)
	}

	fileInfo, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf(statFailedErr, file, err)
	}

	if fileInfo.IsDir() {
		return false, fmt.Errorf(notFileErr, file)
	}

	return true, nil
}

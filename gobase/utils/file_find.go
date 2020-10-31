package utils

import (
	"path/filepath"
)

func IsFileExist(directory string, fileName string) bool {
	matches, _ := filepath.Glob(directory + fileName)
	if len(matches) != 0 {
		return true
	}
	return false
}

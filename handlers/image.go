package handlers

import (
	"errors"
	"math/rand"
	"os"
)

// PickRandomImageName selects a random file name from the specified directory path.
// It returns the name of a randomly chosen file (not a directory) within the given path.
// If the directory cannot be read or contains no files, an error is returned.
//
// Parameters:
//   path - the directory path to search for files.
//
// Returns:
//   string - the randomly selected file name.
//   error  - an error if the directory cannot be read or contains no files.
func PickRandomImageName(path string) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	var list []string
	for _, f := range files {
		if !f.IsDir() {
			list = append(list, f.Name())
		}
	}

	if len(list) == 0 {
		return "", errors.New("no image found")
	}

	return list[rand.Intn(len(list))], nil
}

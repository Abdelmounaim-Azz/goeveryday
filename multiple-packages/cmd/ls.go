package cmd

import (
	"fmt"
	"os"
)

func ExecuteLs(path string) (string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}
	output := fmt.Sprintf("Files in %s\n", path)
	output += "Name\tDirectory\t\n"
	for _, e := range entries {
		output += fmt.Sprintf("%s\t%v\n", e.Name(), e.IsDir())
	}
	return output, nil
}

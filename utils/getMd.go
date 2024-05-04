package utils

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

func GetMd(path string) (string, error) {
	_, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if filepath.Ext(path) != ".md" {
		return "", errors.New("file is not a .md file")
	}

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

package asellus

import (
	"os"
	"path/filepath"
)

func FindFiles(dirPath string, ext string) ([]string, error) {
	var textFiles []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ext {
			textFiles = append(textFiles, path)
		}
		return nil
	})
	return textFiles, err
}

func ForEachFile(dirPath string, ext string, callback func(string) error) ([]string, error) {
	var textFiles []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ext {
			textFiles = append(textFiles, path)
			err = callback(path)
			// generate an error to stop the walk
			if err != nil {
				return err
			}
		}
		return nil
	})
	return textFiles, err
}

func ForEveryFile(dirPath string, callback func(string) error) ([]string, error) {
	var textFiles []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			textFiles = append(textFiles, path)
			err = callback(path)
			// generate an error to stop the walk
			if err != nil {
				return err
			}
		}
		return nil
	})
	return textFiles, err
}

func ReadTextFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteTextFile(path, content string) error {
	// Create a new file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

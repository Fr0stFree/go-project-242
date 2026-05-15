package helpers

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	var totalSize int
	onFile := func(info fs.FileInfo) {
		totalSize += int(info.Size())
	}

	err := iterDir(path, onFile)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s\t%s", formatBytes(totalSize), path), nil
}

func iterDir(path string, onFileCallback func(fs.FileInfo)) error {
	info, err := os.Lstat(path)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		onFileCallback(info)
		return nil
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			err = iterDir(filepath.Join(path, entry.Name()), onFileCallback)
			if err != nil {
				return err
			}
		}
		info, err = entry.Info()
		if err != nil {
			return err
		}
		onFileCallback(info)
	}
	return nil
}

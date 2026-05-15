package helpers

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	var totalSize int
	onFile := func(info fs.FileInfo) {
		fmt.Println(info.Name())
		if all || !strings.HasPrefix(info.Name(), ".") {
			totalSize += int(info.Size())
		}
		
	}

	err := iterDirRec(path, onFile)
	if err != nil {
		return "", err
	}

	var bytes string
	if human {
		bytes = BytesToStringPretty(totalSize)
	} else {
		bytes = bytesToString(totalSize)
	}
	return fmt.Sprintf("%s\t%s", bytes, path), nil
}

func iterDirRec(path string, onFileCallback func(fs.FileInfo)) error {
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
			err = iterDirRec(filepath.Join(path, entry.Name()), onFileCallback)
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

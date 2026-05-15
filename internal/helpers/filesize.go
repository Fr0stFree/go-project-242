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
	var maxDepth = 1
	if recursive {
		maxDepth = -1
	}
	err := iterPathRec(path, 0, maxDepth, func(info fs.FileInfo) {
		if all || !strings.HasPrefix(info.Name(), ".") {
			totalSize += int(info.Size())
		}
	})
	if err != nil {
		return "", err
	}
	return formatResult(totalSize, path, human), nil
}

func formatResult(size int, path string, isHumanReadable bool) string {
	var bytes string
	if isHumanReadable {
		bytes = BytesToStringPretty(size)
	} else {
		bytes = bytesToString(size)
	}
	return fmt.Sprintf("%s\t%s", bytes, path)
}

func iterPathRec(path string, currDepth, maxDepth int, onFileCallback func(fs.FileInfo)) error {
	info, err := os.Lstat(path)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		onFileCallback(info)
		return nil
	}
	if maxDepth >= 0 && currDepth >= maxDepth {
		return nil
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			err = iterPathRec(filepath.Join(path, entry.Name()), currDepth+1, maxDepth, onFileCallback)
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

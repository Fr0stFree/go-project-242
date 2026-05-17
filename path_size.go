package code

import (
	"fmt"
	"io/fs"
	"iter"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	iterPath := newPathIterator(recursive)
	fmtBytes := newBytesFormatter(human)

	var totalSize int
	for info, err := range iterPath(path) {
		if err != nil {
			return "", err
		}
		if shouldSkip(all, info) {
			continue
		}
		totalSize += int(info.Size())
	}
	return fmt.Sprintf("%s\t%s", fmtBytes(totalSize), path), nil
}

func shouldSkip(all bool, info fs.FileInfo) bool {
	if all {
		return false
	}
	isHidden := strings.HasPrefix(info.Name(), ".")
	return isHidden
}

func newPathIterator(recursive bool) func(string) iter.Seq2[fs.FileInfo, error] {
	depth := 0
	maxDepth := 1
	if recursive {
		maxDepth = -1
	}
	return func(path string) iter.Seq2[fs.FileInfo, error] {
		return iterPathRec(path, depth, maxDepth)
	}
}

func iterPathRec(path string, depth, maxDepth int) iter.Seq2[fs.FileInfo, error] {
	return func(yield func(fs.FileInfo, error) bool) {
		info, err := os.Lstat(path)
		if err != nil {
			yield(nil, err)
			return
		}
		if !yield(info, nil) {
			return
		}
		if maxDepth >= 0 && depth >= maxDepth {
			return
		}
		if !info.IsDir() {
			return
		}
		entries, err := os.ReadDir(path)
		if err != nil {
			yield(nil, err)
			return
		}
		for _, entry := range entries {
			childPath := filepath.Join(path, entry.Name())
			for childInfo, err := range iterPathRec(childPath, depth+1, maxDepth) {
				if !yield(childInfo, err) {
					return
				}
			}
		}
	}
}

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
	PB = TB * 1024
	EB = PB * 1024
)

func newBytesFormatter(human bool) func(int) string {
	if human {
		return bytesToStringPretty
	}
	return bytesToString
}

func bytesToString(bytes int) string {
	return fmt.Sprintf("%dB", bytes)
}

func bytesToStringPretty(bytes int) string {
	switch {
	case bytes < KB:
		return fmt.Sprintf("%dB", bytes)
	case bytes < MB:
		return fmt.Sprintf("%.1fKB", float64(bytes)/KB)
	case bytes < GB:
		return fmt.Sprintf("%.1fMB", float64(bytes)/MB)
	case bytes < TB:
		return fmt.Sprintf("%.1fGB", float64(bytes)/GB)
	case bytes < PB:
		return fmt.Sprintf("%.1fPB", float64(bytes)/PB)
	case bytes < EB:
		return fmt.Sprintf("%.1fEB", float64(bytes)/EB)
	default:
		return fmt.Sprintf("%.1fEB", float64(bytes)/EB)
	}
}

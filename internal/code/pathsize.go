package code

import (
	"fmt"
	"io/fs"
	"iter"
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

func newBytesFormatter(human bool) func(int) string {
	if human {
		return bytesToStringPretty
	}
	return bytesToString
}

func shouldSkip(all bool, info fs.FileInfo) bool {
	if all {
		return false
	}
	isHidden := strings.HasPrefix(info.Name(), ".")
	return isHidden
}

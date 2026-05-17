package helpers

import (
	"io/fs"
	"iter"
	"os"
	"path/filepath"
)

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

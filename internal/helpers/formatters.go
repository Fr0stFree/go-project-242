package helpers

import "fmt"

const (
	KB = 1024
	MB = 1024 * 1024
	GB = 1024 * 1024 * 1024
	TB = 1024 * 1024 * 1024 * 1024
)

func formatBytes(bytes int) string {
	switch {
	case bytes < KB:
		return fmt.Sprintf("%dB", bytes)
	case bytes < MB:
		return fmt.Sprintf("%.1fKB", float64(bytes)/KB)
	case bytes < GB:
		return fmt.Sprintf("%.1fMB", float64(bytes)/MB)
	case bytes < TB:
		return fmt.Sprintf("%.1fGB", float64(bytes)/GB)
	default:
		return fmt.Sprintf("%.1fTB", float64(bytes)/TB)
	}
}

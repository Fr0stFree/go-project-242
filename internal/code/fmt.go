package code

import "fmt"

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
	PB = TB * 1024
	EB = PB * 1024
)

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

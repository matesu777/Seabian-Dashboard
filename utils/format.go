package utils

import (
	"fmt"
)

func FormatUsageCPU(usage float64) string {
	return fmt.Sprintf("%.0f%%", usage)
}

func FormatBytes(bytes uint64) string {
	const (
		MB = 1024 * 1024
		GB = 1024 * MB
		TB = 1024 * GB
	)

	switch {
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/TB)
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/GB)
	default:
		return fmt.Sprintf("%.2f MB", float64(bytes)/MB)
	}
}

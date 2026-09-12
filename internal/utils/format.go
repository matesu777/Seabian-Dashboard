package utils

import (
	"fmt"
)

func FormatUsageCPU(usage float64) string {
	return fmt.Sprintf("%.0f%%", usage)
}

func FormatTemp(temp int64) string {
	return fmt.Sprintf("%d°C", temp/1000)
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

func FormatUptime(seconds uint64) string {
	days := seconds / 86400
	seconds %= 86400

	hours := seconds / 3600
	seconds %= 3600

	minutes := seconds / 60
	seconds %= 60

	if days > 0 {
		return fmt.Sprintf("%dd:%dh:%dm:%ds", days, hours, minutes, seconds)
	}

	if hours > 0 {
		return fmt.Sprintf("%dh:%dm", hours, minutes)
	}

	if minutes > 0 {
		return fmt.Sprintf("%dm:%ds", minutes, seconds)
	}

	return fmt.Sprintf("%ds", seconds)
}

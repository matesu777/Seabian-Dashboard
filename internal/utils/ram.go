package utils

import (
	"fmt"

	"github.com/matesu777/Seabian-dashboard/internal/models"
)

func PercentRAM(Ram models.Memory) string {
	ramUsage := float64(Ram.Used) / float64(Ram.Total) * 100
	return fmt.Sprintf("%.1f%%", ramUsage)
}

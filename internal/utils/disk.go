package utils

import "github.com/matesu777/Seabian-dashboard/internal/models"

func TotalDisk(data models.Response) (uint64, uint64) {
	var totalUsed uint64
	var totalSize uint64

	for _, disk := range data.Hardware.Disk {
		if disk.MountPoint == "/" || disk.MountPoint == "/home" {
			totalUsed += disk.Used
			totalSize += disk.Total
		}
	}
	return totalUsed, totalSize
}

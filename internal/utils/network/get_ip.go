package network

import "github.com/matesu777/Seabian-dashboard/internal/models"

func GetNetworkIp(data models.Response) (string, string) {
	var localIP string
	var tailscaleIP string

	for _, network := range data.Hardware.Network {
		switch network.Name {
		case "wlp0s20f3":
			localIP = network.IPv4

		case "tailscale0":
			tailscaleIP = network.IPv4

		case "enp3s0":
			localIP = network.IPv4
		}
	}

	return localIP, tailscaleIP
}

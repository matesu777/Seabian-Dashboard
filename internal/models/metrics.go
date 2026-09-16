package models

import "time"

type Response struct {
	System   System   `json:"system"`
	Hardware Hardware `json:"hardware"`
}

type System struct {
	Hostname  string    `json:"hostname"`
	Uptime    uint64    `json:"uptime"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Hardware struct {
	CPU         CPU         `json:"cpu"`
	Disk        []Disk      `json:"disk"`
	Memory      Memory      `json:"memory"`
	Network     []Network   `json:"network"`
	Temperature Temperature `json:"temperature"`
}

type CPU struct {
	Usage float64 `json:"usage"`
	Cores []Core  `json:"cores"`
}

type Core struct {
	ID    int     `json:"id"`
	Usage float64 `json:"usage"`
}

type Disk struct {
	Device      string  `json:"device"`
	MountPoint  string  `json:"mouint_point"`
	FSType      string  `json:"fsType"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

type Memory struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
}

type Network struct {
	Name    string `json:"name"`
	MAC     string `json:"mac"`
	IPv4    string `json:"ipv4"`
	RxBytes uint64 `json:"rx_bytes"`
	TxBytes uint64 `json:"tx_bytes"`
	RxSpeed uint64 `json:"rx_speed"`
	TxSpeed uint64 `json:"tx_speed"`
}

type NetworkSample struct {
	RxSpeed uint64 `json:"rx_speed"`
	TxSpeed uint64 `json:"tx_speed"`
}

type Temperature struct {
	CPU int64 `json:"cpu"`
}

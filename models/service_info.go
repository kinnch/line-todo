package models

import "time"

type ServiceInfo struct {
	IPAddress   string
	Port        string
	ServiceName string
}

type ServicesInfo []ServiceInfo

type Todo struct {
	Task string
	Time time.Time
}
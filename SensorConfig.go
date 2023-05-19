package main

type SensorDetails struct {
	name   string
	ipAddr string
	port   string
}

func Init(ip, port, name string) *SensorDetails {
	return &SensorDetails{
		ipAddr: ip,
		port:   port,
		name:   name,
	}
}

package config

type SensorDetails struct {
	name   string `json:"Name"`
	ipAddr string `json:"IP"`
	port   string `json:"Port"`
}

/*func Init(ip, port, name string) *SensorDetails {
	return &SensorDetails{
		ipAddr: ip,
		port:   port,
		name:   name,
	}
}*/

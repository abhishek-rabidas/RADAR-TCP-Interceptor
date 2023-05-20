package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

type RadarInterceptorConfig struct {
	Sensor        SensorDetails
	Interceptor   Interceptor
	MaxBufferSize int
}

func LoadConfig() (*RadarInterceptorConfig, error) {

	configFile, err := os.OpenFile("./radarconfig.cfg", os.O_RDONLY, 0666)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	configBytes, _ := io.ReadAll(configFile)

	var config RadarInterceptorConfig = RadarInterceptorConfig{}

	json.Unmarshal(configBytes, &config)

	log.Printf("\n%+v", config)

	return &config, nil

}

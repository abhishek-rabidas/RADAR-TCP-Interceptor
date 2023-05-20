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

	err = json.Unmarshal(configBytes, &config)
	if err != nil {

		return nil, err
	}

	log.Printf("\n%+v", config)

	return &config, nil

}

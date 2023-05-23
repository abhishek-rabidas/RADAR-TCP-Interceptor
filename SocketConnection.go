package main

import (
	log "github.com/sirupsen/logrus"
	"net"
	"radar/StreamUtils"
	"radar/config"
)

type Connection struct {
	connection  net.Conn
	config      config.RadarInterceptorConfig
	interceptor *StreamUtils.InterceptorDetails
}

func NewSocketConnection() (*Connection, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &Connection{
		config: *cfg,
	}, nil
}

func (c *Connection) DialConnection() error {
	var err error

	address := c.config.Sensor.IP + ":" + c.config.Sensor.Port

	raddr, err := net.ResolveTCPAddr("tcp", address)

	c.connection, err = net.DialTCP("tcp", nil, raddr)

	if err != nil {
		log.Error(err)
		return err
	} else {
		log.Infof("-----Connection established [%s]-----", address)
	}
	return nil
}

func (c *Connection) Stream(sync bool) {
	id := 3
	for {
		buff := make([]byte, 512)

		_, err := c.connection.Read(buff)
		if err != nil {
			log.Error(err)
			return
		}

		c.interceptor = StreamUtils.InitializeInterceptor(buff, &c.config.Interceptor, sync)

		c.interceptor.GetPayload(id)
		id++

	}
}

func (c *Connection) Close() {
	defer c.connection.Close()
	log.Info("-----Connection Closed-----")
}

package main

import (
	log "github.com/sirupsen/logrus"
	"net"
	"radar/StreamUtils"
	"radar/config"
)

type Connection struct {
	connection net.Conn
}

func NewSocketConnection(ip, port, name string) *Connection {
	return &Connection{
		sensorDetails: *config.Init(ip, port, name),
	}
}

func (c *Connection) DialConnection() error {
	var err error

	address := c.sensorDetails.ipAddr + ":" + c.sensorDetails.port
	c.connection, err = net.Dial("tcp", address)

	if err != nil {
		log.Error(err)
		return err
	} else {
		log.Infof("-----Connection established [%s]-----", address)
	}
	return nil
}

func (c *Connection) Stream() {
	for {
		buff := make([]byte, 512)

		_, err := c.connection.Read(buff)
		if err != nil {
			log.Error(err)
			return
		}

		c.interceptor = StreamUtils.InitializeInterceptor(buff)
		//log.Infof("[%s]: %x", c.sensorDetails.name, buff) //print buffer
	}
}

func (c *Connection) Close() {
	defer c.connection.Close()
	log.Info("-----Connection Closed-----")
}

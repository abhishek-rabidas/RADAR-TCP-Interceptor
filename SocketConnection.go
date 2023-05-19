package main

import (
	log "github.com/sirupsen/logrus"
	"net"
)

type Connection struct {
	sensorDetails SensorDetails
	connection    net.Conn
}

func NewSocketConnection() *Connection {
	return &Connection{
		sensorDetails: *Init("192.168.1.17", "55555", "umrr0c"),
	}
}

func (c *Connection) DialConnection() {
	var err error

	address := c.sensorDetails.ipAddr + ":" + c.sensorDetails.port
	c.connection, err = net.Dial("tcp", address)

	if err != nil {
		log.Error(err)
	} else {
		log.Infof("-----Connection established [%s]-----", address)
	}

	c.Stream()

	defer c.connection.Close()
	log.Infof("-----Connection Exited [%s]-----", address)

}

func (c *Connection) Stream() {
	for {
		buff := make([]byte, 512)

		_, err := c.connection.Read(buff)
		if err != nil {
			return
		}

		log.Infof("[%s]: %x", c.sensorDetails.name, buff)
	}
}

package main

import log "github.com/sirupsen/logrus"

func main() {

	socket := NewSocketConnection("192.168.1.17", "55555", "umrr0c")
	defer socket.Close()
	err := socket.DialConnection()

	if err != nil {
		log.Error(err)
	}

	socket.Stream()

}

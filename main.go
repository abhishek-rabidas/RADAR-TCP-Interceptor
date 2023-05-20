package main

import log "github.com/sirupsen/logrus"

func main() {

	socket, err := NewSocketConnection()

	if err != nil {
		log.Error(err)
		panic(err)
	}

	defer socket.Close()
	err = socket.DialConnection()

	if err != nil {
		log.Error(err)
		panic(err)
	}

	socket.Stream()

}

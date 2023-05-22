package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	socket, err := NewSocketConnection()

	var sync bool

	if os.Args[1] == "--sync" {
		sync = true
		log.Info("SYNC Mode ON : Sync and Co processor message will appear")
	} else {
		sync = false
	}

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

	socket.Stream(sync)

}

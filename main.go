package main

import (
	"github.com/alex-korobko/http-server-fault-emulator/config"
	"github.com/alex-korobko/http-server-fault-emulator/behaviours"
	"net"
	"log"
)

func main() {
	config.Init()
	portNumberStr := config.GetEmulationPort()
	if len(portNumberStr)==0 {
		log.Fatal("Fatal error: no emulation port found")
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":" + portNumberStr)
	exitIfError(err, "Err resolving TCP addr")
	listener, err := net.ListenTCP("tcp", tcpAddr)
	exitIfError(err, "Err creating listener")
	for {
		conn, err := listener.Accept()
		if err != nil {
			exitIfError(err, "Err accepting connection")
		}
		invokeBehaviour(conn)
		conn.Close()                // we're finished with this client
	}
}
func exitIfError(err error, msg string) {
	if err != nil {
		log.Fatal(msg+ " : " +err.Error())
	}
}

func invokeBehaviour(conn net.Conn) {
	currentBehName := config.GetCurrentBehaviourName()
	currentBehConfig := config.GetCurrentBehaviourParams()
	currentBehFunc := behaviours.GetBehaviourFunc(currentBehName)
	if (currentBehFunc == nil) {
		log.Fatal("There is no func for behaviour : " + currentBehName)
	}
	currentBehFunc(conn, currentBehConfig)
}

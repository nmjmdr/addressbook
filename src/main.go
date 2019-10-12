package main

import (
	"addressbook/router"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func handleExit() {

}

const listenAddress = ":8080"

func main() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM)
	signal.Notify(signalChan, syscall.SIGINT)

	go func() {
		select {
		case sig := <-signalChan:
			// handle graceful close here
			handleExit()
			logrus.Printf("Signal received: %v, Exiting...\n", sig)
			os.Exit(0)
		}
	}()

	logrus.Println("API listening on: %s ...", listenAddress)
	router.Start(listenAddress)
}

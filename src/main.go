package main

import (
	"addressbook/router"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"addressbook/configuration"
	"addressbook/cache"
	"addressbook/autopilot"
	"addressbook/store"
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

	config, err := configuration.ReadConfig()
	fmt.Println(config, err)
	cache := cache.NewRedisCache(config.RedisConfig.Addr, config.RedisConfig.Addr)
	apiProxy := autopilot.NewAutoPilotProxy(config.APIConfig.BaseUrl, config.APIConfig.ApiKey)

	err = cache.Ping()
	if err != nil {
		logrus.Fatalf("Unable to connect to redis, Error: %v", err)
	}

	st := store.NewStore(cache, apiProxy)

	logrus.Printf("API listening on: %s ...\n", listenAddress)
	router.Start(listenAddress, st)
}

package main

import (
	"OrderService/cmd/pub/generator"
	appconfig "OrderService/configs"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	// init config
	config, err := appconfig.Initialize()
	if err != nil {
		logrus.Fatalf("configuration file initialization error: %s", err.Error())
	}

	logrus.Println("Loaded config")

	// connect nats streaming server
	sc, err := stan.Connect(config.Nats.ClusterID, "service")
	if err != nil {
		logrus.Fatalf("connection error with nats streaming: %s", err.Error())
	}
	defer func(sc stan.Conn) {
		err := sc.Close()
		if err != nil {
			logrus.Fatalf("connection closing error with nats streaming: %s", err.Error())
		}
	}(sc)

	logrus.Println("Successful connection to streaming nats server")

	for {
		jsonObject, err := json.Marshal(generator.NewOrder())
		if err != nil {
			logrus.Fatalf("new order processing error: %s", err.Error())
		}
		err = sc.Publish(config.Nats.Subject, jsonObject)
		if err != nil {
			logrus.Fatalf("publishing error from nat streaming server: %s", err.Error())
		}
		time.Sleep(10 * time.Second)
	}
}

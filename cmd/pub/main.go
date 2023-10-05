package main

import (
	"OrderService/cmd/pub/generator"
	"OrderService/configs"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func main() {
	// init config
	if err := configs.InitConf(); err != nil {
		logrus.Fatalf("configuration file initialization error: %s", err.Error())
	}

	logrus.Println("Loaded config")

	// connect nats streaming server
	sc, err := stan.Connect(
		viper.GetString("nats.clusterID"),
		"servicePub")
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
		err = sc.Publish(viper.GetString("nats.subject"), jsonObject)
		if err != nil {
			logrus.Fatalf("publishing error from nat streaming server: %s", err.Error())
		}
		time.Sleep(viper.GetDuration("time") * time.Second)
	}
}

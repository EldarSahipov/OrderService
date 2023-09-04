package main

import (
	appconfig "OrderService/configs"
	_ "OrderService/docs"
	"OrderService/internal/http"
	"OrderService/internal/pkg/handler"
	"OrderService/internal/pkg/nats"
	repository "OrderService/internal/pkg/repository"
	"OrderService/internal/pkg/service"
	"context"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

// @title Order Service
// @version         1.0
// @description "Order service" - a service for managing order data. The service provides a convenient and reliable way to interact with order data through the API. You can use this API to get order information, create, and cache data in memory for faster access. A simple and intuitive API makes it easy to integrate Order Service into your applications and systems.

// @contact.name   Eldar Sahipov
// @contact.url    https://t.me/arassvet
// @contact.email  eldar.shpv@gmail.com

// @host localhost:8080
// @BasePath /

func main() {
	// config
	appConfig, err := appconfig.Initialize()
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Println("Loaded config")

	// database
	postgresDB, err := repository.NewPostgresDB(repository.Config{
		Host:     appConfig.Database.Host,
		Port:     appConfig.Database.Port,
		Username: appConfig.Database.Username,
		Password: appConfig.Database.Password,
		NameDB:   appConfig.Database.Name,
	})
	if err != nil {
		logrus.Fatalf("database connection error: %s", err.Error())
	}

	logrus.Println("Successful database connection")

	repo := repository.NewRepository(postgresDB)
	services := service.NewService(repo)

	// connect nats
	sc, _ := stan.Connect(appConfig.Nats.ClusterID, "sub-1")
	defer func(sc stan.Conn) {
		err := sc.Close()
		if err != nil {
			logrus.Fatalf("Subscription connection closed error: %s", err.Error())
		}
	}(sc)

	logrus.Println("Successful connection to streaming nats server")

	// sub
	subscribe, err := sc.Subscribe(appConfig.Nats.Subject, func(msg *stan.Msg) {
		order, err := nats.UnmarshalTheMessage(string(msg.Data))
		if err != nil {
			logrus.Fatalf("unmarshal message error: %s", err.Error())
		}
		_, err = services.Create(order)
		if err != nil {
			logrus.Fatalf("error adding order to database: %s", err.Error())
		}

	}, stan.DurableName("durableClient"))
	if err != nil {
		logrus.Fatalf("connection close error for subscribing: %s", err.Error())
	}

	logrus.Println("Successful nats streaming subscription")

	// http Server
	server := new(http.Server)

	go func() {
		handlers := handler.NewHandler(services)

		if err := server.Run(appConfig.Server.Port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("http server start error: %s", err.Error())
		}

		logrus.Println("Successful server start")

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT)
	<-quit

	logrus.Println("Server stopped")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("an error occurred while stopping the server: %s")
	}

	err = subscribe.Unsubscribe()
	if err != nil {
		logrus.Fatalf("connection close error for unsubscribing: %s", err.Error())
	}

	postgresDB.Close()

}

package main

import (
	"time"

	"github.com/micro/go-micro"
	goConf "github.com/micro/go-config"
	"github.com/sirupsen/logrus"

	"github.com/lukasjarosch/service-boilerplate/proto/example"
	"github.com/lukasjarosch/service-boilerplate/handler"
	"github.com/lukasjarosch/service-boilerplate/config"
)

// ServiceName is the global service-name
const ServiceName = "go.example.srv"

var (
	cfg config.ServiceConfiguration
	baseLogger *logrus.Logger
)

// init is called on package initialization and can therefore be used to initialize global stuff like logging, config, ..
func init() {
	goConf.Scan(&cfg)
	baseLogger = initLogging(cfg.Log)
}

func main() {
	service := micro.NewService(
		micro.Name(ServiceName),
		micro.RegisterTTL(time.Second*60),
		micro.RegisterInterval(time.Second*15),
		micro.WrapHandler(LogWrapper),
	)
	service.Init()

	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// attach instance to all consecutive logs by default
	logger := baseLogger.WithField("instance", ServiceName + "." + service.Options().Server.Options().Id)
	logger.Info("Starting service " + ServiceName)

	// init database

	// fire
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

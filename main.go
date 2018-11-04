package main

import (
	"log"
	"time"

	"github.com/micro/go-micro"
	goConf "github.com/micro/go-config"

	"github.com/lukasjarosch/service-boilerplate/proto/example"
	"github.com/lukasjarosch/service-boilerplate/handler"
	"github.com/lukasjarosch/service-boilerplate/config"
)

// ServiceName is the global service-name
const ServiceName = "go.example.srv"

var (
	cfg config.ServiceConfiguration
)

// init is called on package initialization and can therefore be used to initialize global stuff like logging, config, ..
func init() {
	goConf.Scan(&cfg)
}

func main() {
	service := micro.NewService(
		micro.Name(ServiceName),
		micro.RegisterTTL(time.Second*60),
		micro.RegisterInterval(time.Second*15),
	)
	service.Init()

	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// init database

	// fire
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

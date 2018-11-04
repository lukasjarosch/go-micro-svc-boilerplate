package main

import (
	"github.com/lukasjarosch/service-boilerplate/handler"
	example "github.com/lukasjarosch/service-boilerplate/proto/example"
	"github.com/micro/go-micro"
	"log"
	"time"
)

// ServiceName is the global service-name
const ServiceName = "go.example.srv"

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

package main

import (
	"time"

	goConf "github.com/micro/go-config"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
	k8s "github.com/micro/kubernetes/go/micro"

	"github.com/lukasjarosch/go-micro-svc-boilerplate/config"
	"github.com/lukasjarosch/go-micro-svc-boilerplate/handler"
	"github.com/lukasjarosch/go-micro-svc-boilerplate/proto/example"
	"github.com/lukasjarosch/go-micro-svc-boilerplate/datastore"
	_ "github.com/lukasjarosch/go-micro-svc-boilerplate/datastore"
)

// ServiceName is the global service-name
const ServiceName = "go.micro.srv.example"

var (
	cfg        config.ServiceConfiguration
	baseLogger *logrus.Logger
)

// init is called on package initialization and can therefore be used to initialize global stuff like logging, config, ..
func init() {
	goConf.Scan(&cfg)
	baseLogger = initLogging(cfg.Log)
}

func main() {

	var service micro.Service

	// init service based on environment
	// If LocalEnv is set, use micro, else use kubernetes-native services
	if cfg.LocalEnv {
		service = micro.NewService()
	} else {
		service = k8s.NewService()
	}

	// setup service
	service.Init(
		micro.Name(ServiceName),
		micro.RegisterTTL(time.Second*60),
		micro.RegisterInterval(time.Second*15),
		micro.WrapHandler(LogWrapper),
	)

	// init database
	if len(cfg.Database.Uri) == 0 {
		baseLogger.Fatal("no database connection string provided")
	}
	if err := datastore.Init(cfg.Database); err != nil {
		baseLogger.Fatal(err)
	}
	baseLogger.Info("database initialized")
	defer datastore.Close()


	// create handlers
	userHandler := handler.InitHelloHandler(baseLogger)

	// register service handlers
	example.RegisterExampleHandler(service.Server(), userHandler)

	// fire
	if err := service.Run(); err != nil {
		baseLogger.Fatal(err)
	}
}

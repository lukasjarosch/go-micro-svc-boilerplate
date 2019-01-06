package main

import (
	"time"

	goConf "github.com/micro/go-config"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
	k8s "github.com/micro/kubernetes/go/micro"

	"github.com/lukasjarosch/go-micro-svc-boilerplate/handler"
	"github.com/lukasjarosch/go-micro-svc-boilerplate/proto/example"
	"github.com/lukasjarosch/go-micro-svc-boilerplate/datastore"
	_ "github.com/lukasjarosch/go-micro-svc-boilerplate/datastore"
	"github.com/lukasjarosch/go-micro-svc-boilerplate/config"
	"github.com/lukasjarosch/go-micro-svc-boilerplate/subscribe"
	"context"
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
	// If environment is set to "k8s", use kubernetes, else use the local environment
	if cfg.Environment == "k8s" {
		service = k8s.NewService()
	} else {
		service = micro.NewService()
	}

	// setup service
	service.Init(
		micro.Name(ServiceName),
		micro.RegisterTTL(time.Second*60),
		micro.RegisterInterval(time.Second*15),
		micro.WrapHandler(HandleWrapper),
		micro.WrapSubscriber(SubscribeWrapper),
	)

	// setup database connection
	initDatabase()
	defer datastore.Close()

	// register example subscriber on topic "topic.example"
	err := micro.RegisterSubscriber(subscribe.ExampleTopic, service.Server(), new(subscribe.ExampleSubscriber))
	if err != nil {
	    baseLogger.WithError(err).Warn("failed to register subscriber")
	}
	baseLogger.WithField("topic", subscribe.ExampleTopic).Info("subscribed to topic")

	// register example service handler
	userHandler := handler.NewExampleHandler(baseLogger)
	example.RegisterExampleHandler(service.Server(), userHandler)

	// uncomment to generate some events on the topic 'topic.example'
	go pubExample(service, 1000)

	// fire
	if err := service.Run(); err != nil {
		baseLogger.Fatal(err)
	}


}

// pubExample is a simple publisher example to periodically publish the ExampleEvent
// Usually the publisher can be attached to the service handler struct to easily pub the events from within the
// business logic
func pubExample(service micro.Service, intervalMs int)  {
	timerChan := time.Tick(time.Duration(intervalMs) * time.Millisecond)
	for range timerChan {

		// example publisher
		publisher := micro.NewPublisher(subscribe.ExampleTopic, service.Client())
		err := publisher.Publish(context.Background(), &example.ExampleEvent{Status:"pubbed"})
		if err != nil {
			baseLogger.WithError(err).Warn("failed to publish event")
		}
		baseLogger.Infof("pubbed to: %s", subscribe.ExampleTopic)
	}
}

// initDatabase ensures that the URI is set, initializes the datastore and defers the close of the connection
func initDatabase() {
	if len(cfg.Database.Uri) == 0 {
		baseLogger.Fatal("no database connection string provided")
	}
	if err := datastore.Init(cfg.Database); err != nil {
		baseLogger.Fatal(err)
	}
	baseLogger.Info("database connection initialized")
}
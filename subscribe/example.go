package subscribe

import (
	"context"
	"github.com/lukasjarosch/go-micro-svc-boilerplate/proto/example"
	"github.com/sirupsen/logrus"
)

type ExampleSubscriber struct {}

const ExampleTopic = "topic.example"

func (s *ExampleSubscriber) Handle(ctx context.Context, event *example.ExampleEvent) error {
	logrus.WithField("status", event.Status).Infof("received event from topic %s", ExampleTopic)
	return nil
}

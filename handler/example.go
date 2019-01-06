package handler

import (
	"context"
	"github.com/sirupsen/logrus"

	"github.com/lukasjarosch/go-micro-svc-boilerplate/datastore"
	ex "github.com/lukasjarosch/go-micro-svc-boilerplate/proto/example"
	"fmt"
)

type Example struct {
	log *logrus.Logger
}

func NewExampleHandler(log *logrus.Logger) *Example {
	return &Example{log: log}
}

func (e *Example) Hello(ctx context.Context, req *ex.HelloRequest, rsp *ex.HelloResponse) error {

	user := &datastore.User{
		Name:  "Hans Peter",
		Email: "hans3@peter.com",
	}

	err := datastore.CreateUser(user)
	if err != nil {
		e.log.WithField("email", user.Email).WithError(err).Info("unable to create user")
		return err
	}

	e.log.WithField("user_id", user.ID).Info("created new user")

	rsp.Status = fmt.Sprintf("created user %s", user.ID)

	return nil
}

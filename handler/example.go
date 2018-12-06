package handler

import (
	"context"
	"github.com/sirupsen/logrus"

	ex "github.com/lukasjarosch/go-micro-svc-boilerplate/proto/example"
	"github.com/lukasjarosch/go-micro-svc-boilerplate/datastore"
)

type Example struct{
	log *logrus.Logger
}

func InitHelloHandler(log *logrus.Logger) *Example {
	return &Example{log: log}
}

func (e *Example) Hello(ctx context.Context, req *ex.HelloRequest, rsp *ex.HelloResponse) error {

	user := &datastore.User{
		Name: "Hans Peter",
		Email: "hans3@peter.com",
	}

	 err := datastore.CreateUser(user)
	 if err != nil {
		e.log.WithField("email", user.Email).WithError(err).Info("unable to create user")
	 	return err
	 }

	 e.log.WithField("user_id", user.ID).Info("created new user")

	return nil
}

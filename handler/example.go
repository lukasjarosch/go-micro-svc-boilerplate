package handler

import (
	"context"
	ex "github.com/lukasjarosch/service-boilerplate/proto/example"
)

type Example struct{}

func (e *Example) Hello(ctx context.Context, req *ex.HelloRequest, rsp *ex.HelloResponse) error {
	return nil
}

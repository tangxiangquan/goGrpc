package handler

import (
	"context"
	"fmt"
	h "goGrpc/proto/hello"
)

type HelloServerImpl struct {

}

func (helloServer *HelloServerImpl) SayHello(ctx context.Context,say *h.Say) (*h.Say, error)  {
	fmt.Println("进入服务........")
	return &h.Say{
		Name: fmt.Sprintf("this is %v",say.Name),
	},nil
}
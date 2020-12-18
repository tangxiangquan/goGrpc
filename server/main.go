package main

import (
	h "goGrpc/proto/hello"
	"goGrpc/server/handler"
	"google.golang.org/grpc"
	"log"
	"net"
)

func newServer()  *handler.HelloServerImpl {
	s := &handler.HelloServerImpl{

	}
	return s
}

func main() {
	lis,err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("failed to listen : %v" ,err)
	}

	grpcServer := grpc.NewServer()

	h.RegisterHelloServer(grpcServer , newServer())

	grpcServer.Serve(lis)
}

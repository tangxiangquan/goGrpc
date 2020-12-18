package main

import (
	"context"
	h "goGrpc/proto/hello"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	addr = "localhost:8081"
	defMsg = "hello txq"
)

func main() {
	var conn *grpc.ClientConn
	//var ctx context.Context
	//var cancel context.CancelFunc
	var err error
	var r *h.Say

	if conn,err = grpc.Dial(addr ,grpc.WithInsecure() , grpc.WithBlock()); err != nil{
		log.Fatalf("conn server err : %v" , err)
	}

	defer conn.Close()

	c := h.NewHelloClient(conn)
	//ctx,cancel = context.WithTimeout(context.Background() , time.Second)
	//defer cancel()

	for  {

		if r,err = c.SayHello(context.Background() , &h.Say{Name: defMsg}); err != nil {
			log.Fatalf("get err : %v" , err)
		}
		log.Printf("get msg :%s" , r.Name)

		time.Sleep(time.Second)
	}


	select {

	}
}



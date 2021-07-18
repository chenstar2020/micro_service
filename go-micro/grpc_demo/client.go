package grpc_demo

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "go-micro/myproto"
	"time"
)

const(
	address = "localhost:1234"
	defaultName = "world"
)

func Init(){
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloserverClient(conn)
	name := defaultName
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Sayhello(ctx, &pb.HelloReq{
		Name: name,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMsg())
}

package main

import (
	"fmt"
	pd "go-micro/myproto"
	"google.golang.org/grpc"
	"log"
	"net"
	"go-micro/grpc_demo"
	"time"
)

func main(){
/*	test := &myproto.Request{
		Group: "aaa",
		Key: "key",
	}
	fmt.Println(test)

	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("编码失败")
	}
	fmt.Println(data)

	new := &myproto.Request{}
	err = proto.Unmarshal(data, new)
	if err != nil {
		fmt.Println("解码失败")
	}
	fmt.Println(new)*/

/*	arith := new(rpc_demo.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e  != nil{
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	time.Sleep(time.Second)

	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil{
		log.Fatal("dialing:", err)
	}
	args := &rpc_demo.Args{
		A: 7, B: 8,
	}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	fmt.Printf("Arith: %d * %d = %d", args.A, args.B, reply)*/
	go func() {
		time.Sleep(time.Second * 3)
		grpc_demo.Init()
	}()

	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("网络错误", err)
	}
	srv := grpc.NewServer()
	pd.RegisterHelloserverServer(srv, &grpc_demo.Server{})
	if err := srv.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

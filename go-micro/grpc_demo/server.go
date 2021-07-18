package grpc_demo

import (
	"context"
	pd "go-micro/myproto"
)

type Server struct{}

func (s *Server) Sayhello(ctx context.Context, req *pd.HelloReq) (*pd.HelloRes, error) {
	return &pd.HelloRes{
		Msg: "messageaa:" + req.Name,
	}, nil
}

func (s *Server) Sayname(ctx context.Context, req *pd.NameReq) (*pd.NameRes, error) {
	return &pd.NameRes{
		Msg: "name:" + req.Name,
	}, nil
}

/*func main(){
	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("网络错误", err)
	}
	srv := grpc_demo.NewServer()
	pd.RegisterHelloserverServer(srv, &server{})
	if err := srv.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}*/

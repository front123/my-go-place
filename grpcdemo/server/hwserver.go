package main

import (
	"context"
	"fmt"
	"net"

	"go-example/grpcdemo/hw"

	"google.golang.org/grpc"
)

func main()  {
	StartHwServer()
}

type HwServer struct {
	hw.UnimplementedGreeterServer
}

func StartHwServer() {
	conn, err := net.Listen("tcp", ":4001")
	if err != nil {
		panic(err)
	}

	fmt.Println("Started grpc hw server!")
	s := grpc.NewServer()
	hw.RegisterGreeterServer(s, &HwServer{})
	s.Serve(conn)
}

func (s *HwServer) SayHello(ctx context.Context, req *hw.HelloReq) (*hw.HelloResp, error) {
	fmt.Println("Get client SayHello request, name:", req.Name)
	return &hw.HelloResp{Message: "Hello, " + req.Name}, nil
}

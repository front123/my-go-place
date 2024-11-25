package main

import (
	"context"
	"fmt"

	"go-example/grpcdemo/hw"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	CallHw()
}

func CallHw() {
	conn, err := grpc.NewClient("localhost:4001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	cli := hw.NewGreeterClient(conn)
	resp, err := cli.SayHello(context.TODO(), &hw.HelloReq{Name: "front"})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.Message)
	}
}

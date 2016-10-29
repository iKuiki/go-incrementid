package main

import (
	"github.com/yinhui87/go-component/config"
	"github.com/yinhui87/go-incrementid/service/increment"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", config.Env("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	increment.RegisterIncrementIdServer(s, &increment.IncrServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}

package main

import (
	"net"

	"github.com/uli-rahmani/grpc/proto"
	"github.com/uli-rahmani/grpc/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	ucm := new(usecase.Math)
	proto.RegisterMathServer(s, ucm)
	reflection.Register(s)

	if e := s.Serve(listener); e != nil {
		panic(e)
	}
}

package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func GetGrpc(port int64, server func(s *grpc.Server)) error {
	//创建一个tcp链接 并打印端口
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	//生成一个新的grpc链接
	s := grpc.NewServer()
	server(s)
	log.Printf("server listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		return err
	}
	return nil
}

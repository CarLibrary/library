package main

import (
	"CarLibrary/carlibrary/carlibrary"
	"CarLibrary/carlibrary/config"
	"CarLibrary/carlibrary/model"
	"fmt"
	pb "github.com/CarLibrary/proto/carlibrary"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {

	//加载配置
	model.InitMYSQL()
	config.Config()
	fmt.Println("ok")

	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCarLibraryServiceServer(s, &carlibrary.CarLibraryServiceSever{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

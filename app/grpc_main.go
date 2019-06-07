package main

import (
	"net"

	"github.com/Huhaokun/loki/api"
	pb "github.com/Huhaokun/loki/api/loki"
	"github.com/Huhaokun/loki/core/trick"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	endpoint = ":33033"
)

func main() {
	log.Printf("Start loki GRPC Service")
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	nodeController := api.NewNodeControllerServer()

	pb.RegisterNodeControllerServer(s, nodeController)
	pb.RegisterStateTrickerServer(s, api.NewStatetrickerServer(nodeController, &trick.StateTricker{}))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

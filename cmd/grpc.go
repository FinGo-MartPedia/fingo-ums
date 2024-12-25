package cmd

import (
	"log"
	"net"

	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))
	if err != nil {
		log.Fatal("Failed to listen grpc port: ", err)
	}

	s := grpc.NewServer()

	logrus.Info("Starting GRPC Server on port: ", helpers.GetEnv("GRPC_PORT", "7000"))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve grpc server: ", err)
	}
}

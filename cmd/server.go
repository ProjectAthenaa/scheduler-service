package main

import (
	"github.com/ProjectAthenaa/scheduler-service/connector"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	"net"
)

func handleServer() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln("error starting listener: ", err)
	}

	server := grpc.NewServer()

	connector.RegisterSchedulerServer(server, &connector.Server{})

	if err = server.Serve(listener); err != nil {
		log.Error(err)
	}
}

package main

import (
	"golangProjects/Microservice/Currency/protos"
	"golangProjects/Microservice/Currency/server"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	// 1. Create a new client with the default settings:
	log1 := log.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(log1)

	protos.RegisterCurrencyServer(gs, cs)
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Error("unable to listen", "error", err)
		os.Exit()
	}
	gs.Serve()

}

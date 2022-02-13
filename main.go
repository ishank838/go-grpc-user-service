package main

import (
	"log"

	"github.com/ishank838/go-users-grpc/config"
	"github.com/ishank838/go-users-grpc/server"
	"google.golang.org/grpc"
)

func main() {
	app := config.MustLoad()
	srv := grpc.NewServer()
	listener := server.SetupServer(&app, srv)

	log.Println("server started at port", app.ServerConfig.Port)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("error at starting server: %v", err)
	}
}

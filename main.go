package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/ishank838/go-users-grpc/config"
	"github.com/ishank838/go-users-grpc/server"
	"google.golang.org/grpc"
)

func main() {
	app := config.MustLoad()
	srv := grpc.NewServer()
	listener := server.SetupServer(&app, srv)

	shutdown(srv)

	log.Println("server started at port", app.ServerConfig.Port)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("error at starting server: %v", err)
	}
}

func shutdown(srv *grpc.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Println("shutting down GRPC server....")
			srv.GracefulStop()
		}
	}()
}

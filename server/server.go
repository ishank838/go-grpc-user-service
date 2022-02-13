package server

import (
	"log"
	"net"

	"github.com/ishank838/go-users-grpc/config"
	userService "github.com/ishank838/go-users-grpc/proto/user"
	users "github.com/ishank838/go-users-grpc/services/users"
	"google.golang.org/grpc"
)

func SetupServer(app *config.Application, srv *grpc.Server) net.Listener {
	setupAllServices(srv)

	l, err := net.Listen(app.ServerConfig.Protocol, "localhost"+app.ServerConfig.Port)
	if err != nil {
		log.Fatalf("cannot listen at port %s using protocol %s : %v",
			app.ServerConfig.Port,
			app.ServerConfig.Protocol,
			err)
	}

	return l
}

func setupAllServices(srv *grpc.Server) {
	user := users.UserService{}
	userService.RegisterUsersServer(srv, &user)
}

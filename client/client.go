package main

import (
	"context"
	"log"

	user "github.com/ishank838/go-users-grpc/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}

	defer conn.Close()

	c := user.NewUsersClient(conn)
	response, err := c.GetAllUsers(context.Background(), &user.Void{})
	if err != nil {
		log.Println(err)
	}
	log.Println(response)
}

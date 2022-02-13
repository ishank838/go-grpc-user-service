package main

import (
	"context"
	"log"

	user "github.com/ishank838/go-users-grpc/proto/user"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
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
